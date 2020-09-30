package main

import (
	"crypto/rsa"
	"flag"
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"github.com/golang/glog"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"

	"garymenezes.com/xposer/common"
)

func timeReceivedMiddleware(c *gin.Context) {
	utc, err := time.LoadLocation("UTC")
	if err != nil {
		glog.Error(err.Error())
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"message": "Couldn't create UTC timestamp."})
		return
	}
	c.Set("TimeReceived", time.Now().In(utc))
}

func jwtMiddleware(key *rsa.PublicKey) gin.HandlerFunc {
	return func(c *gin.Context) {
		reqToken := c.GetHeader("Authorization")
		splitToken := strings.Split(reqToken, "Bearer ")
		if len(splitToken) != 2 {
			c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"message": "Bearer token not in proper format (Bearer: token)"})
			return
		}
		reqToken = strings.TrimSpace(splitToken[1])
		token, err := jwt.ParseWithClaims(reqToken, &common.XposerClaims{}, func(t *jwt.Token) (interface{}, error) {
			return key, nil
		})
		claims, ok := token.Claims.(*common.XposerClaims)
		if !ok || !token.Valid {
			glog.Error(err)
			c.AbortWithStatus(http.StatusUnauthorized)
			return
		}
		c.Set("DeviceName", claims.DeviceName)
	}
}

func main() {
	var pubKeyPath string
	flag.StringVar(&pubKeyPath, "k", "", "public key path")
	var config string
	flag.StringVar(&config, "c", "pong/conf.json", "config file path")
	var test bool
	flag.BoolVar(&test, "test", false, "whether to run in test mode")
	flag.Parse()
	if pubKeyPath == "" {
		glog.Fatal("you must enter a public key path")
	}

	c, err := common.ReadConfig(config)
	if err != nil {
		glog.Fatal("couldn't read config: ", err)
	}

	// Connect to MySQL database
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8", c.MySQLUser, c.MySQLPass, c.MySQLHost, c.MySQLPort, c.MySQLDatabase)
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		glog.Fatal("Failed to connect to database: ", err)
	}

	// Set up database
	db.AutoMigrate(&Ping{})

	// Set GIN release mode
	if !test {
		gin.SetMode(gin.ReleaseMode)
	}

	// Read in public key
	verifyBytes, err := ioutil.ReadFile(pubKeyPath)
	if err != nil {
		glog.Fatal(err)
	}
	verifyKey, err := jwt.ParseRSAPublicKeyFromPEM(verifyBytes)
	if err != nil {
		glog.Fatal(err)
	}

	// Set up router
	r := gin.Default()
	r.Use(common.RequestIdMiddleware)
	r.Use(timeReceivedMiddleware)
	r.Use(jwtMiddleware(verifyKey))

	r.GET("/ping", handlePing(db))
	glog.Fatal(r.Run(":8080"))
}

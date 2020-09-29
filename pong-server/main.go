package main

import (
	"crypto/rsa"
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"strings"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
)

func jwtMiddleware(key *rsa.PublicKey) gin.HandlerFunc {
	return func(c *gin.Context) {
		r := c.Request
		reqToken := r.Header.Get("Authorization")
		splitToken := strings.Split(reqToken, "Bearer ")
		if len(splitToken) != 2 {
			c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": "Bearer token not in proper format (Bearer: token)"})
			return
		}
		reqToken = strings.TrimSpace(splitToken[1])
		token, err := jwt.Parse(reqToken, func(t *jwt.Token) (interface{}, error) {
			return key, nil
		})
		if err != nil || !token.Valid {
			log.Print(err)
			c.AbortWithStatus(http.StatusUnauthorized)
			return
		}
		fmt.Println(token)
	}
}

func main() {
	var pubKeyPath string
	flag.StringVar(&pubKeyPath, "k", "", "public key path")
	flag.Parse()
	if pubKeyPath == "" {
		log.Fatal("you must enter a public key path")
	}

	verifyBytes, err := ioutil.ReadFile(pubKeyPath)
	if err != nil {
		log.Fatal(err)
	}
	verifyKey, err := jwt.ParseRSAPublicKeyFromPEM(verifyBytes)
	if err != nil {
		log.Fatal(err)
	}

	r := gin.Default()
	r.Use(jwtMiddleware(verifyKey))

	r.GET("/ping", func(c *gin.Context) {
		c.String(http.StatusOK, "ok")
	})
	log.Fatal(r.Run(":8080"))
}

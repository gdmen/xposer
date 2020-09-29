package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"os"

	"github.com/dgrijalva/jwt-go"

	"garymenezes.com/xfinity-xposer/common"
)

func fatal(msg interface{}) {
	fmt.Println(msg)
	os.Exit(1)
}

func main() {
	var device string
	flag.StringVar(&device, "d", "", "name for the connecting device")
	var privKeyPath string
	flag.StringVar(&privKeyPath, "k", "", "private key path")
	var jwtOut string
	flag.StringVar(&jwtOut, "o", "bin/xposer.jwt", "path to write out signed jwt")
	flag.Parse()
	if privKeyPath == "" {
		fatal("you must enter a private key path")
	}
	signBytes, err := ioutil.ReadFile(privKeyPath)
	if err != nil {
		fatal(err)
	}
	signKey, err := jwt.ParseRSAPrivateKeyFromPEM(signBytes)
	if err != nil {
		fatal(err)
	}

	claims := common.XposerClaims{
		Device: device,
		StandardClaims: jwt.StandardClaims{
			Issuer: "garymenezes.com",
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodRS256, claims)
	signed, err := token.SignedString(signKey)
	if err != nil {
		fatal(err)
	}
	err = ioutil.WriteFile(jwtOut, []byte(signed), 0700)
	if err != nil {
		fatal(err)
	}
}

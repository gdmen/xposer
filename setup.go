package main

import (
	"flag"
	"io/ioutil"
	"log"

	"github.com/dgrijalva/jwt-go"

	"garymenezes.com/xfinity-xposer/common"
)

func main() {
	var device string
	flag.StringVar(&device, "d", "", "name for the connecting device")
	var privKeyPath string
	flag.StringVar(&privKeyPath, "k", "", "private key path")
	var jwtOut string
	flag.StringVar(&jwtOut, "o", "xposer.jwt", "path to write out signed jwt")
	flag.Parse()
	signBytes, err := ioutil.ReadFile(privKeyPath)
	if err != nil {
		log.Fatal(err)
	}
	signKey, err := jwt.ParseRSAPrivateKeyFromPEM(signBytes)
	if err != nil {
		log.Fatal(err)
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
		log.Fatal(err)
	}
	err = ioutil.WriteFile(jwtOut, []byte(signed), 0700)
	if err != nil {
		log.Fatal(err)
	}
}

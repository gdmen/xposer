package common

import "github.com/dgrijalva/jwt-go"

type XposerClaims struct {
	Device string `json:device`
	jwt.StandardClaims
}

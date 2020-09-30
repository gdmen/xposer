package common

import "github.com/dgrijalva/jwt-go"

type XposerClaims struct {
	DeviceName string `json:device_name`
	jwt.StandardClaims
}

package utils

import (
	"io/ioutil"
	"time"

	"github.com/dgrijalva/jwt-go"
)

// GenerateTokenString read private key and return token string
func GenerateTokenString(id string) (string, error) {
	// read privateKey
	privateKey, err := ioutil.ReadFile("keys/jwtRS512.key")
	if err != nil {
		return "", err
	}

	// Sign and get the complete encoded token as a string
	key, err := jwt.ParseRSAPrivateKeyFromPEM(privateKey)
	if err != nil {
		return "", err
	}

	// Create the token
	token := jwt.New(jwt.SigningMethodRS512)
	token.Claims = jwt.MapClaims{
		"id":  id,
		"exp": time.Now().Add(time.Hour * 72).Unix(),
		"iat": time.Now().Unix(),
	}
	return token.SignedString(key)
}

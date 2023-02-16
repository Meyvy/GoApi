package auth

import (
	"monitor/config"
	"time"

	"github.com/golang-jwt/jwt"
)

// secret key for signing it should be loaded from a file or env variable and
// be generated from a cryptographically secure hash function
var SecretKey = []byte("Anything really!")

// creates a token sets the expiration date and the user id in the jwt token
func CreateToken(id int64) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"exp": time.Now().Add(config.TokenDuration).Unix(),
		"id":  id,
	})
	tokenString, err := token.SignedString(SecretKey)
	if err != nil {
		return "", err
	}
	return tokenString, nil
}

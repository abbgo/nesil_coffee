package helpers

import (
	"os"
	"strconv"
	"time"

	"github.com/dgrijalva/jwt-go"
)

var JwtKey = []byte(os.Getenv("JWT_SECRET_KEY"))

type JWTClaimForAdmin struct {
	Login string `json:"login"`
	ID    string `json:"id"`
	jwt.StandardClaims
}

func GenerateAccessToken(login, id string) (string, error) {
	accessTokenTimeOut, err := strconv.Atoi(os.Getenv("ACCESS_TOKEN_TIMEOUT"))
	if err != nil {
		return "", err
	}
	expirationTimeAccessToken := time.Now().Add(time.Duration(accessTokenTimeOut) * time.Second)

	claimsAccessToken := &JWTClaimForAdmin{
		Login: login,
		ID:    id,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: expirationTimeAccessToken.Unix(),
		},
	}
	accessToken := jwt.NewWithClaims(jwt.SigningMethodHS256, claimsAccessToken)
	accessTokenString, err := accessToken.SignedString(JwtKey)
	if err != nil {
		return "", err
	}

	return accessTokenString, nil
}

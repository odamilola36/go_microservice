package security

import (
	"errors"
	"fmt"
	"github.com/dgrijalva/jwt-go"
	"net/http"
	"os"
	"strings"
	"time"
)

var (
	ErrInvalidToken = errors.New("invalid token")
	jwtSecret       = []byte(os.Getenv("JWT_SECRET"))
)

type TokenPayload struct {
	UserId    string
	CreatedAt time.Time
	ExpiresAt time.Time
}

func NewToken(userId string) (string, error) {
	claims := &jwt.StandardClaims{
		ExpiresAt: time.Now().Add(time.Hour * 72).Unix(),
		Issuer:    userId,
		IssuedAt:  time.Now().Unix(),
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString(jwtSecret)
}

func parseJwtCallback(token *jwt.Token) (interface{}, error) {
	if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
		return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
	}
	return jwtSecret, nil
}

func ParseToke(token string) (*jwt.Token, error) {
	return jwt.Parse(token, parseJwtCallback)
}

func NewTokenPayload(tokenString string) (*TokenPayload, error) {
	token, erro := ParseToke(tokenString)
	if erro != nil {
		return nil, erro
	}
	claims, ok := token.Claims.(jwt.MapClaims)
	if !token.Valid || !ok {
		return nil, ErrInvalidToken
	}
	createdAt := claims["iat"].(float64)
	expiresAt := claims["exp"].(float64)
	sub := claims["iss"].(string)
	return &TokenPayload{
		UserId:    sub,
		CreatedAt: time.Unix(int64(createdAt), 0),
		ExpiresAt: time.Unix(int64(expiresAt), 0),
	}, nil
}

func ExtractToken(r *http.Request) (string, error) {
	authToken := strings.TrimSpace(r.Header.Get("Authorization"))
	splitToken := strings.Split(authToken, " ")
	if len(splitToken) != 2 {
		return "", ErrInvalidToken
	}
	return splitToken[1], nil
}

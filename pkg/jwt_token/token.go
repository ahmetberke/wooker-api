package jwt_token

import (
	"errors"
	"github.com/dgrijalva/jwt-go"
	"time"
)

var ExpiredDuration time.Duration = time.Minute * 100
var TokenInvalidErr error = errors.New("token invalid")
var TokenExpiredErr error = errors.New("token expired")

type claims struct {
	UserID uint
	jwt.StandardClaims
}

func GenerateToken(userID uint, secretKey string) (string, error)  {
	cl := jwt.NewWithClaims(jwt.SigningMethodHS256, &claims{
		UserID: userID,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Add(ExpiredDuration).Unix(),
		},
	})
	return cl.SignedString([]byte(secretKey))
}

func ParseToken(token string, secretKey string) (uint, error) {
	claim := &claims{}
	tkn, err := jwt.ParseWithClaims(token, claim, func(token *jwt.Token) (interface{}, error) {
		return []byte(secretKey), nil
	})
	if err != nil {
		return 0, err
	}
	if !tkn.Valid {
		return 0, TokenInvalidErr
	}
	if time.Unix(claim.ExpiresAt, 0).Sub(time.Now()) > ExpiredDuration {
		return 0, TokenExpiredErr
	}
	return claim.UserID, nil
}
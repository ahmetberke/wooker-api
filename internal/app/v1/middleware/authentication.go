package middleware

import (
	"github.com/ahmetberke/wooker-api/configs"
	"github.com/ahmetberke/wooker-api/internal/models"
	"github.com/ahmetberke/wooker-api/internal/service"
	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"net/http"
	"strings"
	"time"
)

type Auth struct {
	UserService *service.UserService
}

type AuthResponse struct {
	Success bool `json:"success"`
	Error string `json:"error"`
	Data interface{} `json:"data"`
}

func NewAuthMiddleware(userService *service.UserService) *Auth {
	return &Auth{
		UserService: userService,
	}
}

type Claims struct {
	UserID uint
	jwt.StandardClaims
}

var ExpiredDuration time.Duration = time.Minute * 100

func Authenticate(user *models.User) (string, error) {
	claims := jwt.NewWithClaims(jwt.SigningMethodHS256, &Claims{
		UserID: user.ID,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Add(ExpiredDuration).Unix(),
		},
	})
	token, err := claims.SignedString([]byte(configs.Manager.JWTSecretKey))
	if err != nil {
		return "", err
	}
	return token, nil
}

func (a *Auth) Authorization(c *gin.Context) {
	// example [autherization] field : Bearer XXXXXXXXXXXXXXXXX
	tokenAr := strings.Split(c.GetHeader("authorization"), " ")
	if len(tokenAr) <= 1 {
		c.Next()
		return
	}
	token := tokenAr[1]
	if token == "" {
		c.Next()
		return
	}
	claim := &Claims{}
	tkn, err := jwt.ParseWithClaims(token, claim, func(token *jwt.Token) (interface{}, error) {
		return []byte(configs.Manager.JWTSecretKey), nil
	})
	if err != nil {
		c.Next()
		return
	}
	if !tkn.Valid {
		c.Next()
		return
	}
	if time.Unix(claim.ExpiresAt, 0).Sub(time.Now()) > ExpiredDuration {
		c.Next()
		return
	}
	user, err := a.UserService.FindByID(claim.UserID)
	if err != nil {
		c.Next()
		return
	}
	c.Set("x-user", user)
}

func (a *Auth) AdminAuthorization(c *gin.Context) {
	val, ok := c.Get("x-user")
	if !ok {
		c.JSON(http.StatusUnauthorized, AuthResponse{
			Success: false,
			Error:   "unauthorized",
			Data:    nil,
		})
		return
	}
	user := val.(*models.User)
	if !user.IsAdmin {
		c.JSON(http.StatusUnauthorized, AuthResponse{
			Success: false,
			Error:   "unauthorized",
			Data:    nil,
		})
		return
	}
	c.Next()
}
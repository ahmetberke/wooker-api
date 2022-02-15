package middleware

import (
	"github.com/ahmetberke/wooker-api/configs"
	"github.com/ahmetberke/wooker-api/pkg/jwt_token"
	"github.com/gin-gonic/gin"
	"strings"
)

func (m *Manager) Authorization(c *gin.Context)  {

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

	userID, err := jwt_token.ParseToken(token, configs.Manager.JWTSecretKey)
	if err != nil {
		c.Next()
		return
	}

	user, err := m.UserService.FindByID(userID)
	if err != nil {
		c.Next()
		return
	}

	c.Set("x-user", user)

}
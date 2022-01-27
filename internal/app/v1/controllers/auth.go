package controllers

import (
	"github.com/ahmetberke/wooker-api/internal/auth"
	"github.com/gin-gonic/gin"
	"net/http"
)

type Auth struct {
	GoogleAuth *auth.Google
}

func (a *Auth) URL() gin.HandlerFunc {
	return func(context *gin.Context) {
		url := a.GoogleAuth.GenerateURL()
		context.JSON(http.StatusOK, gin.H{
			"data" : gin.H{
				"url" : url,
			},
		})
	}
}

func (a *Auth) Login() gin.HandlerFunc {
	return func(c *gin.Context) {
		code := c.Query("code")
		state := c.Query("state")
		data, err := a.GoogleAuth.GetUserData(state, code)
		if err != nil {
			c.AbortWithStatus(400)
		}
		c.JSON(200, data)
	}
}
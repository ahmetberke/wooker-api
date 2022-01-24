package controllers

import (
	"github.com/ahmetberke/wooker-api/internal/auth"
	"github.com/gin-gonic/gin"
	"net/http"
)

type User struct {
	GoogleAuth *auth.Google
}

func (u *User) URL() gin.HandlerFunc {
	return func(context *gin.Context) {
		url := u.GoogleAuth.GenerateURL()
		context.JSON(http.StatusOK, gin.H{
			"data" : gin.H{
				"url" : url,
			},
		})
	}
}
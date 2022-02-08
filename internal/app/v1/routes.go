package api

import (
	"github.com/ahmetberke/wooker-api/internal/app/v1/controllers"
	"github.com/gin-gonic/gin"
)

func (a *api) HelloRouteInitialise()  {
	hello := a.Router.Group("/hello")
	{
		hello.GET("/", func(context *gin.Context) {
			context.JSON(200, gin.H{
				"message": "hello world",
			})
		})
	}
}

func (a *api) UserRoutesInitialize(uc *controllers.UserController) {
	user := a.Router.Group("/user")
	{
		user.GET("/", uc.GoogleAuth.IsAdmin ,uc.All)
		user.GET("/auth", uc.Auth)
		user.GET("/url", uc.URL)
		user.GET("/:username", uc.Get)
		user.PUT("/:username", uc.Update)
	}
}

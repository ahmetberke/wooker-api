package api

import "github.com/gin-gonic/gin"

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

func (a *api) UserRoutesInitialize() {
	uc := a.controllers.User
	user := a.Router.Group("/user")
	{
		user.GET("/url", uc.URL())
	}
}
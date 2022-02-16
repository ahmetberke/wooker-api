package api

import (
	"github.com/ahmetberke/wooker-api/internal/app/v1/controllers"
)

func (a *api) AuthRoutesInitialize(ac *controllers.AuthController)  {
	auth := a.Router.Group("/auth")
	{
		auth.GET("/google", ac.AuthenticationWithGoogle)
		auth.GET("/url", ac.URL)
	}
}

func (a *api) UserRoutesInitialize(uc *controllers.UserController) {
	user := a.Router.Group("/users")
	{
		user.GET("/",uc.All)
		user.GET("/:username", uc.Get)
		user.PUT("/:username",a.Middleware.IsLoggedUserOrAdminChecking, uc.Update)
	}
}

func (a *api) WordRoutesInitialize(wc *controllers.WordController)  {
	word := a.Router.Group("/words")
	{
		word.GET("/", wc.All)
		word.GET("/:id", wc.Get)
		word.POST("/", wc.New)
		word.DELETE("/:id", wc.Delete)
		word.PUT("/:id", wc.Update)
	}
}
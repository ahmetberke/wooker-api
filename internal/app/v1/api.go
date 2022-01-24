package api

import (
	"fmt"
	"github.com/ahmetberke/wooker-api/configs"
	"github.com/ahmetberke/wooker-api/internal/app/v1/controllers"
	"github.com/ahmetberke/wooker-api/internal/auth"
	"github.com/gin-gonic/gin"
)

type api struct {
	PORT string
	Engine *gin.Engine
	Router *gin.RouterGroup
	controllers *controllers.Controller
}

func NewAPI(config *configs.Manager) (*api, error)  {

	a := api{
		PORT: config.HostCredentials.PORT,
		Engine: gin.Default(),
	}
	a.Router = a.Engine.Group("/v1")

	newAuth := auth.NewOauth2(config.Oauth2Credentials.ClientID, config.Oauth2Credentials.ClientSecret)

	a.controllers = &controllers.Controller{
		User: &controllers.User{
			GoogleAuth: newAuth,
		},
	}

	a.HelloRouteInitialise()
	a.UserRoutesInitialize()

	return &a, nil
}

func (a *api) Run()  {
	fmt.Println(a.Engine.Run(":" + a.PORT))
}

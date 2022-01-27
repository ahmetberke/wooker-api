package api

import (
	"fmt"
	"github.com/ahmetberke/wooker-api/configs"
	"github.com/ahmetberke/wooker-api/internal/app/v1/controllers"
	"github.com/ahmetberke/wooker-api/internal/auth"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"time"
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

	a.Engine.Use(cors.New(cors.Config{
			AllowOrigins:     []string{"http://localhost:3000"},
			AllowMethods:     []string{"GET", "POST", "PUT", "PATCH", "DELETE", "HEAD"},
			AllowHeaders:     []string{"Origin", "Content-Length", "Content-Type", "Authorization"},
			ExposeHeaders:    []string{"Content-Length"},
			AllowCredentials: true,
			MaxAge: 12 * time.Hour,
	}))
	a.Router = a.Engine.Group("/v1")

	newAuth := auth.NewOauth2(config.Oauth2Credentials.ClientID, config.Oauth2Credentials.ClientSecret)

	a.controllers = &controllers.Controller{
		User: &controllers.User{},
		Auth: &controllers.Auth{
			GoogleAuth: newAuth,
		},
	}

	a.HelloRouteInitialise()
	a.UserRoutesInitialize()
	a.AuthRoutesInitialize()

	return &a, nil
}

func (a *api) Run() {
	fmt.Println(a.Engine.Run(":" + a.PORT))
}

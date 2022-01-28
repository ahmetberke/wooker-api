package api

import (
	"fmt"
	"github.com/ahmetberke/wooker-api/configs"
	"github.com/ahmetberke/wooker-api/internal/app/v1/controllers"
	"github.com/ahmetberke/wooker-api/internal/auth"
	"github.com/ahmetberke/wooker-api/internal/service"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"time"
)

type api struct {
	PORT string
	DB *gorm.DB
	Engine *gin.Engine
	Router *gin.RouterGroup
	controllers *controllers.Controller
}

func NewAPI(config *configs.Manager, db *gorm.DB) (*api, error)  {

	a := api{
		PORT: config.HostCredentials.PORT,
		DB: db,
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
		User: &controllers.User{
			Service: &service.UserService{
				Database: db,
			},
		},
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

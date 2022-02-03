package api

import (
	"fmt"
	"github.com/ahmetberke/wooker-api/configs"
	"github.com/ahmetberke/wooker-api/internal/app/v1/controllers"
	"github.com/ahmetberke/wooker-api/internal/app/v1/middleware"
	"github.com/ahmetberke/wooker-api/internal/auth"
	"github.com/ahmetberke/wooker-api/internal/repository"
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
}

func NewAPI(db *gorm.DB) (*api, error)  {

	a := api{
		PORT: configs.Manager.HostCredentials.PORT,
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

	a.Router.Use(middleware.Verificate)

	a.InitUser(configs.Manager.Oauth2Credentials.ClientID, configs.Manager.Oauth2Credentials.ClientSecret)

	return &a, nil
}

func (a *api) InitUser(gClientID string, gClientSecret string)  {
	repo := repository.NewUserRepository(a.DB)
	serv := service.NewUserSercive(repo)
	google := auth.NewOauth2(gClientID, gClientSecret)
	userController := controllers.User{Service: serv, GoogleAuth: google}
	a.UserRoutesInitialize(&userController)
}

func (a *api) Run() {
	fmt.Println(a.Engine.Run(":" + a.PORT))
}

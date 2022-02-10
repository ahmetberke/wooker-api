package api

import (
	"fmt"
	"github.com/ahmetberke/wooker-api/configs"
	"github.com/ahmetberke/wooker-api/internal/app/v1/controllers"
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

	userRepository := repository.NewUserRepository(a.DB)
	userService := service.NewUserService(userRepository)

	oauth2 := auth.NewOauth2(configs.Manager.Oauth2Credentials.ClientID, configs.Manager.Oauth2Credentials.ClientSecret, userService)
	a.Router.Use(oauth2.Authorization)

	userController := controllers.UserController{Service: userService, GoogleAuth: oauth2}
	a.UserRoutesInitialize(&userController)

	wordRepository := repository.NewWordRepository(db)
	wordService := service.NewWordService(wordRepository)
	wordController := controllers.WordController{Service: wordService, Auth: oauth2}
	a.WordRoutesInitialize(&wordController)

	return &a, nil
}

func (a *api) Run() {
	fmt.Println(a.Engine.Run(":" + a.PORT))
}

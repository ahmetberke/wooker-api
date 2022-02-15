package api

import (
	"fmt"
	"github.com/ahmetberke/wooker-api/configs"
	"github.com/ahmetberke/wooker-api/internal/app/v1/controllers"
	"github.com/ahmetberke/wooker-api/internal/app/v1/middleware"
	"github.com/ahmetberke/wooker-api/internal/repository"
	"github.com/ahmetberke/wooker-api/internal/service"
	"github.com/ahmetberke/wooker-api/pkg/google"
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
	Middleware *middleware.Manager
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
	wordRepository := repository.NewWordRepository(db)
	languageRepository := repository.NewLanguageRepository(a.DB)

	userService := service.NewUserService(userRepository)
	wordService := service.NewWordService(wordRepository, userRepository, languageRepository)

	a.Middleware = middleware.NewManager(userService)

	googleOa := google.NewGoogleOauth2(configs.Manager.Oauth2Credentials.ClientID, configs.Manager.Oauth2Credentials.ClientSecret)

	a.Router.Use(a.Middleware.Authorization)

	userController := controllers.UserController{Service: userService}
	wordController := controllers.WordController{Service: wordService}
	authController := controllers.AuthController{UserService: userService, Google: googleOa}

	a.UserRoutesInitialize(&userController)
	a.WordRoutesInitialize(&wordController)
	a.AuthRoutesInitialize(&authController)

	return &a, nil
}

func (a *api) Run() {
	fmt.Println(a.Engine.Run(":" + a.PORT))
}

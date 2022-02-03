package controllers

import (
	"fmt"
	"github.com/ahmetberke/wooker-api/internal/app/v1/middleware"
	"github.com/ahmetberke/wooker-api/internal/auth"
	"github.com/ahmetberke/wooker-api/internal/models"
	"github.com/ahmetberke/wooker-api/internal/service"
	"github.com/ahmetberke/wooker-api/pkg/utils"
	"github.com/gin-gonic/gin"
	"net/http"
)

type User struct {
	GoogleAuth *auth.Google
	Service *service.UserService
}

func (u *User) URL(c *gin.Context) {
	url := u.GoogleAuth.GenerateURL()
	c.JSON(http.StatusOK, gin.H{
		"url":url,
	})
}

func (u *User) Auth(c *gin.Context) {

	// If user already logged in, return bad request error because this request is unnecessary
	_, isExists := c.Get("x-user-id")
	if isExists {
		c.AbortWithStatus(http.StatusBadRequest)
	}

	code := c.Query("code")
	state := c.Query("state")

	// pulling the user data using state and code values from google services
	gresp, err := u.GoogleAuth.GetUserData(state, code)
	if err != nil {
		c.AbortWithStatus(http.StatusBadRequest)
		return
	}

	// if the user already registered, it is just authorized
	cUser, err := u.Service.FindByGoogleID(gresp.ID)
	if err == nil {
		token, err := middleware.Authenticate(cUser)
		if err != nil {
			c.AbortWithStatus(http.StatusUnauthorized)
		}
		c.JSON(http.StatusOK, gin.H{
			"user" : models.ToUserDTO(cUser),
			"token" : token,
		})
		return
	}

	// Creating and saving new user
	user := models.User{
		GoogleID: gresp.ID,
		Email: gresp.Email,
		EmailVerified: gresp.VerifiedEmail,
		Picture: gresp.Picture,
	}

	user.Username = utils.GenerateUsernameFromEmail(gresp.Email)

	// If the username generated from the email is already taken, the alternative username is set
	tUser, err := u.Service.MultipleFindByUsername(user.Username)
	if len(tUser) > 0 {
		user.Username = fmt.Sprintf("%v%v", user.Username, len(tUser))
	}

	// saving user to database
	okUser, err := u.Service.Save(&user)
	if err != nil {
		c.AbortWithStatus(http.StatusInternalServerError)
		return
	}

	token, err := middleware.Authenticate(&user)
	if err != nil {
		c.AbortWithStatus(http.StatusUnauthorized)
	}
	c.JSON(http.StatusCreated, gin.H{m
		"user" : models.ToUserDTO(okUser),
		"token" : token,
	})
	return
}
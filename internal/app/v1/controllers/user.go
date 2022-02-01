package controllers

import (
	"github.com/ahmetberke/wooker-api/internal/auth"
	"github.com/ahmetberke/wooker-api/internal/models"
	"github.com/ahmetberke/wooker-api/internal/service"
	"github.com/ahmetberke/wooker-api/pkg/utils"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"time"
)

type User struct {
	GoogleAuth *auth.Google
	Service *service.UserService
}

func (u *User) URL() gin.HandlerFunc {
	return func(c *gin.Context) {
		url := u.GoogleAuth.GenerateURL()
		c.JSON(http.StatusOK, gin.H{
			"url":url,
		})
	}
}

func (u *User) Auth() gin.HandlerFunc {
	return func(c *gin.Context) {

		code := c.Query("code")
		state := c.Query("state")

		gresp, err := u.GoogleAuth.GetUserData(state, code)
		if err != nil {
			c.AbortWithStatus(http.StatusBadRequest)
			return
		}

		cUser, err := u.Service.FindByGoogleID(gresp.ID)
		if err == nil {
			c.JSON(http.StatusOK, models.ToUserDTO(cUser))
			return
		}

		user := models.User{
			GoogleID: gresp.ID,
			Email: gresp.Email,
			EmailVerified: gresp.VerifiedEmail,
			Picture: gresp.Picture,
		}

		t := time.Now()
		user.CreatedDate = t
		user.LastUpdatedDate = t

		user.Username = utils.GenerateUsernameFromEmail(gresp.Email)

		/*
		tUser, err := u.Service.MultipleFindByUsername(user.Username)
		if len(tUser) > 0 {
			user.Username = fmt.Sprintf("%v%v", user.Username, len(tUser))
		}
		*/

		okUser, err := u.Service.Save(&user)
		if err != nil {
			c.AbortWithStatus(http.StatusInternalServerError)
			log.Printf("error: %v", err.Error())
			return
		}

		c.JSON(http.StatusCreated, models.ToUserDTO(okUser))
		return
	}
}
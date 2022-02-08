package controllers

import (
	"errors"
	"fmt"
	"github.com/ahmetberke/wooker-api/internal/auth"
	"github.com/ahmetberke/wooker-api/internal/models"
	"github.com/ahmetberke/wooker-api/internal/service"
	"github.com/ahmetberke/wooker-api/pkg/utils"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"log"
	"net/http"
	"strconv"
	"time"
)

type UserController struct {
	GoogleAuth *auth.Google
	Service *service.UserService
}

type UserResponse struct {
	Success bool `json:"success"`
	Error string `json:"error"`
	Data interface{} `json:"data"`
}

func (u *UserController) URL(c *gin.Context) {
	url := u.GoogleAuth.GenerateURL()
	c.JSON(http.StatusOK, gin.H{
		"url":url,
	})
}

func (u *UserController) Auth(c *gin.Context) {

	// If user already logged in, return bad request error because this request is unnecessary
	userI, isExists := c.Get("x-user")
	if isExists {
		loggedUser := userI.(*models.User)
		log.Printf("id: %v, username: %v", loggedUser.ID, loggedUser.Username)
		c.AbortWithStatusJSON(http.StatusBadRequest, UserResponse{
			Success: false,
			Error:   "user is already logged in",
			Data:    nil,
		})
		return
	}

	code := c.Query("code")
	state := c.Query("state")

	// pulling the user data using state and code values from google services
	token, err := u.GoogleAuth.GetToken(state, code)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, UserResponse{
			Success: false,
			Error:   "something is wrong",
			Data:    nil,
		})
		return
	}
	gresp, err := u.GoogleAuth.GetUserData(token.AccessToken)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, UserResponse{
			Success: false,
			Error:   "something is wrong",
			Data:    nil,
		})
		return
	}

	// if the user already registered, it is just authorized
	cUser, err := u.Service.FindByGoogleID(gresp.ID)
	if err == nil {
		c.JSON(http.StatusOK, UserResponse{
			Success: true,
			Error:   "",
			Data:    gin.H{
				"user" : models.ToUserDTO(cUser),
				"token" : token.AccessToken,
			},
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

	// set created and updated times
	t := time.Now()
	user.CreatedAt = t
	user.UpdatedAt = t

	// If the username generated from the email is already taken, the alternative username is set
	tUser, err := u.Service.MultipleFindByUsername(user.Username)
	if len(tUser) > 0 {
		user.Username = fmt.Sprintf("%v%v", user.Username, len(tUser))
	}

	// saving user to database
	okUser, err := u.Service.Save(&user)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, UserResponse{
			Success: false,
			Error:   "something is wrong",
			Data:    nil,
		})
		return
	}

	c.JSON(http.StatusCreated, UserResponse{
		Success: true,
		Error:   "",
		Data:    gin.H{
			"user" : models.ToUserDTO(okUser),
			"token" : token.AccessToken,
		},
	})

	return
}

func (u *UserController) Get(c *gin.Context)  {
	username := c.Param("username")

	// get user and if there is an error return bad request
	user, err := u.Service.FindByUsername(username)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			c.AbortWithStatusJSON(http.StatusBadRequest, UserResponse{
				Success: false,
				Error:   "user not found",
				Data:    nil,
			})
			return
		}
		c.AbortWithStatusJSON(http.StatusBadRequest, UserResponse{
			Success: false,
			Error:   "something is wrong",
			Data:    nil,
		})
		return
	}

	// checking if signed user is the requested user
	var logged_in bool = false
	id := c.GetString("x-user-id")
	if id == strconv.Itoa(int(user.ID)) {
		logged_in = true
	}


	c.JSON(http.StatusOK, UserResponse{
		Success: true,
		Error:   "",
		Data:    gin.H{
			"user" : user,
			"logged_in" : logged_in,
		},
	})
	return
}

func (u *UserController) All(c *gin.Context)  {
	var limit int = 10
	limit, err := strconv.Atoi(c.Param("limit"))
	if err != nil {
		limit = 10
	}

	users, err := u.Service.GetAll(limit)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, UserResponse{
			Success: false,
			Error:   "something is wrong",
			Data:    nil,
		})
		return
	}

	var usersDTO []models.UserDTO
	for _, u := range users {
		usersDTO = append(usersDTO, *models.ToUserDTO(&u))
	}

	c.JSON(http.StatusOK, UserResponse{
		Success: true,
		Error:   "",
		Data:    usersDTO,
	})

}

func (u *UserController) Update(c *gin.Context) {
	username := c.Param("username")

	var userDTO models.UserDTO
	err := c.ShouldBindJSON(&userDTO)
	if err  != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, UserResponse{
			Success: false,
			Error:   "cannot bind json to object",
			Data:    nil,
		})
		return
	}

	user := models.ToUser(&userDTO)

	updatedUser, err := u.Service.UpdateByUsername(username, user)
	if err != nil {
		c.JSON(http.StatusBadRequest, UserResponse{
			Success: false,
			Error:   "cannot user update",
			Data:    nil,
		})
		return
	}

	c.JSON(http.StatusOK, UserResponse{
		Success: true,
		Error:   "",
		Data:    models.ToUserDTO(updatedUser),
	})
}
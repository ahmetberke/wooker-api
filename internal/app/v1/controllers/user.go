package controllers

import (
	"errors"
	"fmt"
	"github.com/ahmetberke/wooker-api/internal/app/v1/errorss"
	"github.com/ahmetberke/wooker-api/internal/app/v1/response"
	"github.com/ahmetberke/wooker-api/internal/auth"
	"github.com/ahmetberke/wooker-api/internal/models"
	"github.com/ahmetberke/wooker-api/internal/service"
	"github.com/ahmetberke/wooker-api/pkg/utils"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"net/http"
	"strconv"
	"time"
)

type UserController struct {
	GoogleAuth *auth.Google
	Service *service.UserService
}

func (u *UserController) URL(c *gin.Context) {
	url := u.GoogleAuth.GenerateURL()

	var resp response.AnyResponse
	resp.Code = http.StatusOK
	resp.Data = map[string]string{"url" : url}

	c.JSON(resp.Code, resp)
}

func (u *UserController) Auth(c *gin.Context) {

	var resp response.AuthResponse

	// If user already logged in, return bad request error because this request is unnecessary
	_, isExists := c.Get("x-user")
	if isExists {
		resp.Code = http.StatusBadRequest
		resp.Error = errorss.AlreadyLoggedIn
		c.AbortWithStatusJSON(resp.Code, resp)
		return
	}

	code := c.Query("code")
	state := c.Query("state")

	// pulling the user data using state and code values from google services
	token, err := u.GoogleAuth.GetToken(state, code)
	if err != nil {
		resp.Code = http.StatusBadRequest
		resp.Error = errorss.CannotGetGoogleToken
		c.AbortWithStatusJSON(resp.Code, resp)
		return
	}

	googleResponse, err := u.GoogleAuth.GetUserData(token.AccessToken)
	if err != nil {
		resp.Code = http.StatusBadRequest
		resp.Error = errorss.CannotRetrieveDataFromGoogle
		c.AbortWithStatusJSON(resp.Code, resp)
		return
	}

	// if the user already registered, it is just authorized
	cUser, err := u.Service.FindByGoogleID(googleResponse.ID)
	if err == nil {

		resp.Code = http.StatusOK
		resp.Token = token.AccessToken
		resp.User = models.ToUserDTO(cUser)
		resp.LoggedIn = true

		c.AbortWithStatusJSON(resp.Code, resp)

		return
	}

	// Creating and saving new user
	user := models.User{
		GoogleID: googleResponse.ID,
		Email: googleResponse.Email,
		EmailVerified: googleResponse.VerifiedEmail,
		Picture: googleResponse.Picture,
	}

	user.Username = utils.GenerateUsernameFromEmail(googleResponse.Email)

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
		resp.Code = http.StatusInternalServerError
		resp.Error = errorss.UnableUserSaveToDB
		c.AbortWithStatusJSON(resp.Code, resp)
		return
	}

	resp.Code = http.StatusOK
	resp.Token = token.AccessToken
	resp.LoggedIn = true
	resp.User = models.ToUserDTO(okUser)

	c.JSON(resp.Code, resp)

	return
}

func (u *UserController) Get(c *gin.Context)  {

	var resp response.UserResponse

	username := c.Param("username")

	// get user and if there is an error return bad request
	user, err := u.Service.FindByUsername(username)
	if err != nil {

		if errors.Is(err, gorm.ErrRecordNotFound) {
			resp.Code = http.StatusBadRequest
			resp.Error = errorss.UserNotFound
			c.AbortWithStatusJSON(resp.Code, resp)
			return
		}

		resp.Code = http.StatusBadRequest
		resp.Error = errorss.SomethingIsWrong
		c.AbortWithStatusJSON(resp.Code, resp)

		return
	}

	// checking if signed user is the requested user
	userI, isExists := c.Get("x-user")
	if isExists {
		userT := userI.(*models.User)
		if userT.ID == user.ID {
			resp.LoggedIn = true
		}
	}

	resp.Code = http.StatusOK
	resp.User = models.ToUserDTO(user)

	c.JSON(resp.Code, resp)

	return
}

func (u *UserController) All(c *gin.Context)  {

	var resp response.UsersResponse

	limit, err := strconv.Atoi(c.Param("limit"))
	if err != nil {
		limit = 10
	}

	users, err := u.Service.GetAll(limit)
	if err != nil {
		resp.Code = http.StatusBadRequest
		resp.Error = errorss.SomethingIsWrong
		c.AbortWithStatusJSON(resp.Code, resp)
		return
	}

	for _, u := range users {
		resp.Users = append(resp.Users, models.ToUserDTO(&u))
	}

	resp.Code = http.StatusOK
	c.JSON(resp.Code, resp)

}

func (u *UserController) Update(c *gin.Context) {

	var resp response.UserResponse

	username := c.Param("username")

	var userDTO models.UserDTO
	err := c.ShouldBindJSON(&userDTO)
	if err  != nil {
		resp.Code = http.StatusBadRequest
		resp.Error = errorss.CannotBind
		c.AbortWithStatusJSON(resp.Code, resp)
		return
	}

	user := models.ToUser(&userDTO)

	updatedUser, err := u.Service.UpdateByUsername(username, user)
	if err != nil {
		resp.Code = http.StatusBadRequest
		resp.Error = errorss.CannotUserUpdate
		c.AbortWithStatusJSON(resp.Code, resp)
		return
	}

	resp.Code = http.StatusOK
	resp.User = models.ToUserDTO(updatedUser)
	resp.LoggedIn = true
	c.JSON(resp.Code, resp)

}
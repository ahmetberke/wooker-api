package controllers

import (
	"errors"
	"github.com/ahmetberke/wooker-api/internal/app/v1/errorss"
	"github.com/ahmetberke/wooker-api/internal/app/v1/response"
	"github.com/ahmetberke/wooker-api/internal/models"
	"github.com/ahmetberke/wooker-api/internal/service"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"net/http"
	"strconv"
)

type UserController struct {
	Service *service.UserService
}

func NewUserController(userService *service.UserService) *UserController {
	return &UserController{
		Service: userService,
	}
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

	search := c.Query("search")
	limit, err := strconv.Atoi(c.Query("limit"))
	if err != nil {
		limit = 10
	}

	var preAdded bool = false
	if c.Query("preAdded") == "true" {
		preAdded = true
	}

	users, err := u.Service.GetAll(limit, search, preAdded)
	if err != nil {
		resp.Code = http.StatusBadRequest
		resp.Error = errorss.SomethingIsWrong
		c.AbortWithStatusJSON(resp.Code, resp)
		return
	}

	for _, u := range users {
		resp.Users = append(resp.Users, *models.ToUserDTO(&u))
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
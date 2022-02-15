package controllers

import (
	"fmt"
	"github.com/ahmetberke/wooker-api/configs"
	"github.com/ahmetberke/wooker-api/internal/app/v1/errorss"
	"github.com/ahmetberke/wooker-api/internal/app/v1/response"
	"github.com/ahmetberke/wooker-api/internal/google"
	"github.com/ahmetberke/wooker-api/internal/models"
	"github.com/ahmetberke/wooker-api/internal/service"
	"github.com/ahmetberke/wooker-api/pkg/jwt_token"
	"github.com/ahmetberke/wooker-api/pkg/utils"
	"github.com/gin-gonic/gin"
	"net/http"
	"time"
)

type AuthController struct {
	UserService *service.UserService
	Google *google.GoogleOuath2
}

func (a *AuthController) URL(c *gin.Context) {
	url := a.Google.GenerateURL()

	var resp response.AnyResponse
	resp.Code = http.StatusOK
	resp.Data = map[string]string{"url" : url}

	c.JSON(resp.Code, resp)
}

func (a *AuthController) AuthenticationWithGoogle(c *gin.Context)  {
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
	googleToken, err := a.Google.GetToken(state, code)
	if err != nil {
		resp.Code = http.StatusBadRequest
		resp.Error = errorss.CannotGetGoogleToken
		c.AbortWithStatusJSON(resp.Code, resp)
		return
	}

	googleResponse, err := a.Google.GetUserData(googleToken.AccessToken)
	if err != nil {
		resp.Code = http.StatusBadRequest
		resp.Error = errorss.CannotRetrieveDataFromGoogle
		c.AbortWithStatusJSON(resp.Code, resp)
		return
	}

	// if the user already registered, it is just authorized
	cUser, err := a.UserService.FindByGoogleID(googleResponse.ID)
	if err == nil {

		// generating new token
		token, err := jwt_token.GenerateToken(cUser.ID, configs.Manager.JWTSecretKey)
		if err != nil {
			resp.Code = http.StatusBadRequest
			resp.Error = err.Error()
			c.AbortWithStatusJSON(resp.Code, resp)
			return
		}

		resp.Code = http.StatusOK
		resp.Token = token
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
	tUser, err := a.UserService.MultipleFindByUsername(user.Username)
	if len(tUser) > 0 {
		user.Username = fmt.Sprintf("%v%v", user.Username, len(tUser))
	}

	// saving user to database
	okUser, err := a.UserService.Save(&user)
	if err != nil {
		resp.Code = http.StatusInternalServerError
		resp.Error = errorss.UnableUserSaveToDB
		c.AbortWithStatusJSON(resp.Code, resp)
		return
	}

	// generating new token
	token, err := jwt_token.GenerateToken(cUser.ID, configs.Manager.JWTSecretKey)
	if err != nil {
		resp.Code = http.StatusBadRequest
		resp.Error = err.Error()
		c.AbortWithStatusJSON(resp.Code, resp)
		return
	}

	resp.Code = http.StatusOK
	resp.Token = token
	resp.LoggedIn = true
	resp.User = models.ToUserDTO(okUser)

	c.JSON(resp.Code, resp)
	return

}
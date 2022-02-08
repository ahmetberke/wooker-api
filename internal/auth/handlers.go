package auth

import (
	"encoding/json"
	"fmt"
	"github.com/ahmetberke/wooker-api/internal/models"
	"github.com/gin-gonic/gin"
	"golang.org/x/oauth2"
	"net/http"
	"strings"
)

func (g Google) GenerateURL() string {
	return g.Oauth2.AuthCodeURL(g.State)
}

func (g Google) GetToken(state string, code string) (*oauth2.Token, error)  {

	if state != g.State {
		return nil, fmt.Errorf("invalid oauth state")
	}

	token, err := g.Oauth2.Exchange(oauth2.NoContext, code)
	if err != nil {
		return nil, fmt.Errorf("code exchange failed: %s", err.Error())
	}

	return token, nil

}

func (g Google) GetUserData(accessToken string) (*UserResponse, error) {

	resp, err := http.Get(g.DataURL + accessToken)
	if err != nil {
		return nil, fmt.Errorf("failed getting user info: %s", err.Error())
	}

	defer func() {
		_ = resp.Body.Close()
	}()

	var userResponse UserResponse
	err = json.NewDecoder(resp.Body).Decode(&userResponse)
	if err != nil {
		return &userResponse, fmt.Errorf("failed decoding user info: %s", err.Error())
	}

	return &userResponse, err

}

func (g Google) Authorization (c *gin.Context) {

	tokenAr := strings.Split(c.GetHeader("authorization"), " ")
	if len(tokenAr) <= 1 {
		c.Next()
		return
	}
	token := tokenAr[1]

	userResponse, err := g.GetUserData(token)
	if err != nil {
		c.Next()
	}

	user, err := g.UserService.FindByGoogleID(userResponse.ID)
	if err != nil {
		c.Next()
	}

	c.Set("x-user", user)
	c.Next()

}

func (g Google) IsAdmin(c *gin.Context)  {
	userI, isExist := c.Get("x-user")
	if !isExist {
		c.AbortWithStatus(http.StatusUnauthorized)
		return
	}

	loggedUser := userI.(*models.User)
	if !loggedUser.IsAdmin {
		c.AbortWithStatus(http.StatusUnauthorized)
		return
	}

	c.Next()

}
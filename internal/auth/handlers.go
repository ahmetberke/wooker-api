package auth

import (
	"encoding/json"
	"fmt"
	"golang.org/x/oauth2"
	"log"
	"net/http"
)

func (g Google) GenerateURL() string {
	return g.Oauth2.AuthCodeURL(g.State)
}

func (g Google) GetUserData(state string, code string) (UserResponse, error) {
	ur := UserResponse{}
	if state != g.State {
		return ur, fmt.Errorf("invalid oauth state")
	}

	token, err := g.Oauth2.Exchange(oauth2.NoContext, code)
	if err != nil {
		return ur, fmt.Errorf("code exchange failed: %s", err.Error())
	}
	log.Printf(token.AccessToken)

	response, err := http.Get("https://www.googleapis.com/oauth2/v2/userinfo?access_token=" + token.AccessToken)
	if err != nil {
		return ur, fmt.Errorf("failed getting user info: %s", err.Error())
	}

	defer func() {
		_ = response.Body.Close()
	}()

	err = json.NewDecoder(response.Body).Decode(&ur)
	if err != nil {
		return ur, fmt.Errorf("failed decoding user info: %s", err.Error())
	}
	return ur, nil
}
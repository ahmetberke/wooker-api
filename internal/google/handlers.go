package google

import (
	"encoding/json"
	"fmt"
	"golang.org/x/oauth2"
	"net/http"
)

func (g GoogleOuath2) GenerateURL() string {
	// Generate a URL for login page
	return g.Oauth2.AuthCodeURL(g.State)
}

func (g *GoogleOuath2) GetToken(state string, code string) (*oauth2.Token, error)  {

	// Checking whether the state received with the request is correct.
	if state != g.State {
		return nil, fmt.Errorf("invalid oauth state")
	}

	// Generating token with code
	token, err := g.Oauth2.Exchange(oauth2.NoContext, code)
	if err != nil {
		return nil, fmt.Errorf("code exchange failed: %s", err.Error())
	}

	return token, nil

}

func (g *GoogleOuath2) GetUserData(accessToken string) (*UserResponse, error) {

	// Pulling User Data from Google Service with token
	resp, err := http.Get(g.DataURL + accessToken)

	if err != nil {
		return nil, fmt.Errorf("failed getting user info: %s", err.Error())
	}

	defer func() {
		_ = resp.Body.Close()
	}()

	// Decoding Response data to userResponse struct
	var userResponse UserResponse
	err = json.NewDecoder(resp.Body).Decode(&userResponse)
	if err != nil {
		return &userResponse, fmt.Errorf("failed decoding user info: %s", err.Error())
	}

	if userResponse.Error != nil {
		return &userResponse, fmt.Errorf("%s", userResponse.Error.Message)
	}

	return &userResponse, err

}

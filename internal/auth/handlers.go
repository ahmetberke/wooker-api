package auth

import (
	"fmt"
	"golang.org/x/oauth2"
	"io/ioutil"
	"net/http"
)

func (g Google) GenerateURL() string {
	return g.Oauth2.AuthCodeURL(g.State)
}

func (g Google) GetUserData(state string, code string) ([]byte, error) {
	if state != g.State {
		return nil, fmt.Errorf("invalid oauth state")
	}

	token, err := g.Oauth2.Exchange(oauth2.NoContext, code)
	if err != nil {
		return nil, fmt.Errorf("code exchange failed: %s", err.Error())
	}

	response, err := http.Get("https://www.googleapis.com/oauth2/v2/userinfo?access_token=" + token.AccessToken)
	if err != nil {
		return nil, fmt.Errorf("failed getting user info: %s", err.Error())
	}

	defer func() {
		_ = response.Body.Close()
	}()
	contents, err := ioutil.ReadAll(response.Body)
	if err != nil {
		return nil, fmt.Errorf("failed reading response body: %s", err.Error())
	}
	return contents, nil
}
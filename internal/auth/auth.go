package auth

import (
	"github.com/ahmetberke/wooker-api/internal/service"
	"golang.org/x/oauth2"
	"golang.org/x/oauth2/google"
)

type Google struct {
	Oauth2 *oauth2.Config
	State string
	DataURL string
	UserService *service.UserService
}

func NewOauth2(clientID string, clientSecret string, userService *service.UserService) *Google {
	return &Google{
		Oauth2: &oauth2.Config{
			RedirectURL: "http://localhost:3000/auth/",
			ClientID: clientID,
			ClientSecret: clientSecret,
			Scopes: []string{"https://www.googleapis.com/auth/userinfo.email"},
			Endpoint: google.Endpoint,
		},
		State: "AUgmJCg",
		DataURL: "https://www.googleapis.com/oauth2/v2/userinfo?access_token=",
		UserService: userService,
	}
}

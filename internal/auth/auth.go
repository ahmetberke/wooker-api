package auth

import (
	"golang.org/x/oauth2"
	"golang.org/x/oauth2/google"
)

type Google struct {
	Oauth2 *oauth2.Config
	State string
}

func NewOauth2(clientID string, clientSecret string) *Google {
	return &Google{
		Oauth2: &oauth2.Config{
			RedirectURL: "http://localhost:3000/auth/",
			ClientID: clientID,
			ClientSecret: clientSecret,
			Scopes: []string{"https://www.googleapis.com/auth/userinfo.email"},
			Endpoint: google.Endpoint,
		},
		State: "AUgmJCg",
	}
}

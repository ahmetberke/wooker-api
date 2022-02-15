package google

import (
	"golang.org/x/oauth2"
	"golang.org/x/oauth2/google"
)

type GoogleOuath2 struct {
	Oauth2 *oauth2.Config
	State string
	DataURL string
}

func NewGoogleOauth2(clientID string, clientSecret string) *GoogleOuath2 {
	return &GoogleOuath2{
		Oauth2: &oauth2.Config{
			RedirectURL: "http://localhost:3000/auth/",
			ClientID: clientID,
			ClientSecret: clientSecret,
			Scopes: []string{"https://www.googleapis.com/auth/userinfo.email"},
			Endpoint: google.Endpoint,
		},
		State: "AUgmJCg",
		DataURL: "https://www.googleapis.com/oauth2/v2/userinfo?access_token=",
	}
}

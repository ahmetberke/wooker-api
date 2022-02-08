package configs

import (
	"github.com/joho/godotenv"
	"os"
)

var Manager manager

type manager struct {
	DBCredentials *dbCredentials
	HostCredentials *hostCredentials
	Oauth2Credentials *oauth2Credentials
}

type  hostCredentials struct {
	PORT string
}

type  dbCredentials struct {
	Host string
	Port string
	Name string
	User string
	Password string
	SSLMode string
}

type  oauth2Credentials struct {
	ClientID string
	ClientSecret string
}

func (m *manager) EnvInitialise(path string) {

	// migrating .env queries to OS Environment
	err := godotenv.Load(path)
	if err != nil {
		panic(err)
	}

	// Initilazing
	m.DBCredentials = &dbCredentials{
		 Host: os.Getenv("DB_HOST"),
		 Port: os.Getenv("DB_PORT"),
		 Name: os.Getenv("DB_NAME"),
		 User: os.Getenv("DB_USER"),
		 Password: os.Getenv("DB_PASSWORD"),
		 SSLMode: os.Getenv("DB_SSL_MODE"),
	}

	m.HostCredentials = &hostCredentials{
		PORT: os.Getenv("PORT"),
	}

	m.Oauth2Credentials = &oauth2Credentials{
		ClientID: os.Getenv("OAUTH2_GOOGLE_CLIENT_ID"),
		ClientSecret: os.Getenv("OAUTH2_GOOGLE_CLIENT_SECRET"),
	}

}
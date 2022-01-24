package configs

import (
	"github.com/joho/godotenv"
	"os"
)

type Manager struct {
	DBCredentials *dbCredentials
	HostCredentials *hostCredentials
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
}

func (m *Manager) EnvInitialise(path string) {

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
	}

	m.HostCredentials = &hostCredentials{
		PORT: os.Getenv("PORT"),
	}

}
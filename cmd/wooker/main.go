package main

import (
	"github.com/ahmetberke/wooker-api/configs"
	api "github.com/ahmetberke/wooker-api/internal/app/v1"
	"github.com/ahmetberke/wooker-api/internal/database"
)


func init()  {

	configs.Manager.EnvInitialise(".env")
}

func main() {
	dc := configs.Manager.DBCredentials
	db, err := database.ConnectToDB(&database.DBConfig{
		Host: dc.Host,
		Port: dc.Port,
		Name: dc.Name,
		User: dc.User,
		Password: dc.Password,
		SSLMode: dc.SSLMode,
	})
	if err != nil {
		panic(err)
	}

	a, err := api.NewAPI(db)
	if err != nil {
		panic(err)
	}
	a.Run()
}
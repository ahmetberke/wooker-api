package main

import (
	"github.com/ahmetberke/wooker-api/configs"
	api "github.com/ahmetberke/wooker-api/internal/app/v1"
)

var config *configs.Manager

func init()  {
	config = &configs.Manager{}
	config.EnvInitialise(".env")
}

func main() {
	a, err := api.NewAPI(config)
	if err != nil {
		panic(err)
	}
	a.Run()
}
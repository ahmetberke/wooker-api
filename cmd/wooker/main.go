package main

import (
	"github.com/ahmetberke/wooker-api/configs"
)

func main() {
	config := configs.Manager{}
	config.EnvInitialise(".env")
}
package main

import (
	"github.com/JoaoVitor615/go-api-poc/config"
	"github.com/JoaoVitor615/go-api-poc/router"
)

var (
	logger *config.Logger
)

func main() {

	logger = config.GetLogger("main")

	//Initialize configs
	err := config.Init()

	if err != nil {
		logger.Errorf("Error initializing configs: %v", err)
		return
	}

	//Initialize router
	router.Initialize()

}

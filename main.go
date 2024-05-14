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
	db, err := config.Init()

	logger.Infof("Database connection: %v", db)

	if err != nil {
		logger.Errorf("Error initializing configs: %v", err)
		return
	}

	//Initialize router
	router.Initialize()

}

package main

import (
	"fmt"

	"github.com/JoaoVitor615/go-api-poc/config"
	"github.com/JoaoVitor615/go-api-poc/router"
)

func main() {

	//Initialize configs
	_, err := config.Init()

	if err != nil {
		fmt.Println(err)
		return
	}

	//Initialize router
	router.Initialize()

}

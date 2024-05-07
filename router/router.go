package router

import "github.com/gin-gonic/gin"

func Initialize() {
	// Create a new Gin router
	router := gin.Default()

	InitializeRoutes(router)

	router.Run(":8080")
}

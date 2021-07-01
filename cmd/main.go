package main

import (
	"github.com/gin-gonic/gin"

	"github.com/DePavelPo/go_training_web_with_gin/internal/routes"
)

var router *gin.Engine

func main() {

	// Set the router as the default one provided by Gin
	router = gin.Default()

	// Process the templates at the start so that they don't have to be loaded
	// from the disk again. This makes serving HTML pages very fast.
	router.LoadHTMLGlob("templates/*")

	// Define the route for the index page and display the index.html template
	// To start with, we'll use an inline route handler. Later on, we'll create
	// standalone functions that will be used as route handlers.

	routes.InitializeRoutes(router)

	// Start serving the application
	router.Run()

}

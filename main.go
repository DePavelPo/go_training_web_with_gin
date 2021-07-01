// main.go

package main

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"github.com/DePavelPo/go_training_web_with_gin/routes"
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

	routes.initializeRoutes()

	router.GET("/", func(c *gin.Context) {

		c.HTML(

			// Set the HTTP status to 200 (OK)
			http.StatusOK,

			// Use the index.html template
			"index.html",

			// Pass the data that the page uses (in this case, 'title')
			gin.H{
				"title": "Home Page",
			},
		)
	})

	// Start serving the application
	router.Run()

}

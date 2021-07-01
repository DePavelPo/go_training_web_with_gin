package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func initializeRoutes(c *gin.Engine) {
	g := c.Group("/v1/")

	g.GET("/", showIndexPage)
}

func showIndexPage(c *gin.Context) {
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
}

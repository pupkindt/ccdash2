package main

import (
	"net/http"
	//"strconv"

	"github.com/gin-gonic/gin"
)

func showIndexPage(c *gin.Context) {

	var indexTemplate string = "index.html"

	if useLocalBootstrap {
		indexTemplate = "index-local.html"
	}

	// Call the HTML method of the Context to render a template
	c.HTML(
		// Set the HTTP status to 200 (OK)
		http.StatusOK,
		// Use the index.html template
		indexTemplate,
		// Pass the data that the page uses (in this case, 'title')
		gin.H{
			"title": "Home Page",
		},
	)

}

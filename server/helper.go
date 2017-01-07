package main

import (
	"html/template"
	"net/http"

	"github.com/gin-gonic/gin"
)

// Index renders the general template on all routes,
func Index(c *gin.Context) {
	c.HTML(
		http.StatusOK,
		"index.html",
		gin.H{
			"Root":     Config.Server.Root,
			"Endpoint": Config.Server.Endpoint,
		},
	)
}

// Template loads the template to make it parseable.
func Template() *template.Template {
	file, _ := ReadFile("index.html")

	return template.Must(
		template.New(
			"index.html",
		).Parse(
			string(file),
		),
	)
}

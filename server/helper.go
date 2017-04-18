package main

import (
	"html/template"
	"io/ioutil"
	"net/http"
	"path"

	"github.com/Sirupsen/logrus"
	"github.com/gin-gonic/gin"
)

// Index renders the general template on all routes,
func Index(c *gin.Context) {
	c.HTML(
		http.StatusOK,
		"index.html",
		gin.H{
			"Endpoint": Config.Server.Endpoint,
		},
	)
}

// Template loads the template to make it parseable.
func Template() *template.Template {
	var (
		file []byte
		err  error
	)

	if Config.Server.Static != "" {
		file, err = ioutil.ReadFile(
			path.Join(
				Config.Server.Static,
				"index.html",
			),
		)

		if err != nil {
			logrus.Errorf("Failed to read index template. %s", err)
		}
	} else {
		file, err = ReadFile("index.html")

		if err != nil {
			logrus.Errorf("Failed to read index template. %s", err)
		}
	}

	return template.Must(
		template.New(
			"index.html",
		).Parse(
			string(file),
		),
	)
}

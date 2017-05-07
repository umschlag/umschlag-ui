package main

import (
	"html/template"
	"io/ioutil"
	"net/http"
	"os"
	"path"

	"github.com/Sirupsen/logrus"
	"github.com/Unknwon/com"
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

	staticPath := path.Join(
		Config.Server.Static,
		"index.html",
	)

	if Config.Server.Static != "" && com.IsFile(staticPath) {
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

// Assets initializes the static files.
func Assets() http.FileSystem {
	return ChainedFS{}
}

// ChainedFS is a simple HTTP filesystem including custom path.
type ChainedFS struct {
}

// Open just implements the HTTP filesystem interface.
func (c ChainedFS) Open(origPath string) (http.File, error) {
	if Config.Server.Static != "" {
		if com.IsDir(Config.Server.Static) {
			customPath := path.Join(Config.Server.Static, origPath)

			if com.IsFile(customPath) {
				f, err := os.Open(customPath)

				if err != nil {
					return nil, err
				}

				return f, nil
			}
		} else {
			logrus.Warnf("Custom assets directory doesn't exist")
		}
	}

	f, err := FS.OpenFile(CTX, origPath, os.O_RDONLY, 0644)

	if err != nil {
		return nil, err
	}

	return f, nil
}

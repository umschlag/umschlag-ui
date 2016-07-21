package main

import (
	"net/http"
	"os"
	"path"
	"runtime"

	"github.com/Sirupsen/logrus"
	"github.com/elazarl/go-bindata-assetfs"
	"github.com/gin-gonic/gin"
	"github.com/urfave/cli"
)

//go:generate go-bindata -ignore "\\.go" -pkg main -prefix ../ -o bindata.go ../assets/...
//go:generate go fmt bindata.go
//go:generate sed -i "s/Html/HTML/" bindata.go
//go:generate sed -i "s/Css/CSS/" bindata.go

func main() {
	runtime.GOMAXPROCS(runtime.NumCPU())

	app := cli.NewApp()
	app.Name = "umschlag-ui"
	app.Version = Version
	app.Author = "Thomas Boerger <thomas@webhippie.de>"
	app.Usage = "A docker distribution management system"

	app.Flags = []cli.Flag{
		cli.BoolTFlag{
			Name:        "update, u",
			Usage:       "Enable auto update",
			EnvVar:      "UMSCHLAG_UPDATE",
			Destination: &Config.Debug,
		},
		cli.BoolFlag{
			Name:        "debug",
			Usage:       "Activate debug information",
			EnvVar:      "UMSCHLAG_DEBUG",
			Destination: &Config.Debug,
		},
	}

	app.Before = func(c *cli.Context) error {
		logrus.SetOutput(os.Stdout)

		if Config.Debug {
			logrus.SetLevel(logrus.DebugLevel)
		} else {
			logrus.SetLevel(logrus.InfoLevel)
		}

		if Config.Update {
			Update()
		}

		return nil
	}

	app.Commands = []cli.Command{
		cli.Command{
			Name:  "server",
			Usage: "Start the Umschlag UI",
			Flags: []cli.Flag{
				cli.StringFlag{
					Name:        "addr",
					Value:       ":9000",
					Usage:       "Address to bind the server",
					EnvVar:      "UMSCHLAG_UI_ADDR",
					Destination: &Config.Server.Addr,
				},
				cli.StringFlag{
					Name:        "cert",
					Value:       "",
					Usage:       "Path to SSL cert",
					EnvVar:      "UMSCHLAG_UI_CERT",
					Destination: &Config.Server.Cert,
				},
				cli.StringFlag{
					Name:        "key",
					Value:       "",
					Usage:       "Path to SSL key",
					EnvVar:      "UMSCHLAG_UI_KEY",
					Destination: &Config.Server.Key,
				},
				cli.StringFlag{
					Name:        "root",
					Value:       "/",
					Usage:       "Root folder of the app",
					EnvVar:      "UMSCHLAG_UI_ROOT",
					Destination: &Config.Server.Root,
				},
				cli.StringFlag{
					Name:        "endpoint",
					Value:       "http://localhost:8000",
					Usage:       "URL for the API server",
					EnvVar:      "UMSCHLAG_UI_ENDPOINT",
					Destination: &Config.Server.Endpoint,
				},
			},
			Action: func(c *cli.Context) {
				logrus.Infof("Starting the UI on %s", Config.Server.Addr)

				if Config.Debug {
					gin.SetMode(gin.DebugMode)
				} else {
					gin.SetMode(gin.ReleaseMode)
				}

				e := gin.New()

				e.SetHTMLTemplate(
					Template(),
				)

				e.Use(SetLogger())
				e.Use(SetRecovery())

				for _, folder := range []string{"fonts", "images", "scripts", "styles"} {
					e.StaticFS(
						path.Join(Config.Server.Root, folder),
						&assetfs.AssetFS{
							Asset:     Asset,
							AssetDir:  AssetDir,
							AssetInfo: AssetInfo,
							Prefix:    path.Join("assets", folder),
						},
					)
				}

				e.GET(Config.Server.Root, Index)
				e.NoRoute(Index)

				if Config.Server.Cert != "" && Config.Server.Key != "" {
					logrus.Fatal(
						http.ListenAndServeTLS(
							Config.Server.Addr,
							Config.Server.Cert,
							Config.Server.Key,
							e,
						),
					)
				} else {
					logrus.Fatal(
						http.ListenAndServe(
							Config.Server.Addr,
							e,
						),
					)
				}
			},
		},
	}

	cli.HelpFlag = cli.BoolFlag{
		Name:  "help, h",
		Usage: "Show the help, so what you see now",
	}

	cli.VersionFlag = cli.BoolFlag{
		Name:  "version, v",
		Usage: "Print the current version of that tool",
	}

	app.Run(os.Args)
}

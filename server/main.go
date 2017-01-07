package main

import (
	"crypto/tls"
	"fmt"
	"net/http"
	"net/url"
	"os"
	"path"
	"runtime"
	"time"

	"github.com/Sirupsen/logrus"
	"github.com/elazarl/go-bindata-assetfs"
	"github.com/gin-gonic/gin"
	"github.com/urfave/cli"
	"golang.org/x/crypto/acme/autocert"
)

//go:generate go-bindata -ignore "\\.go" -pkg main -o bindata.go ../assets/...
//go:generate go fmt bindata.go
//go:generate sed -i.bak "s/Html/HTML/" bindata.go
//go:generate sed -i.bak "s/Css/CSS/" bindata.go
//go:generate rm -f bindata.go.bak

func main() {
	runtime.GOMAXPROCS(runtime.NumCPU())

	app := cli.NewApp()
	app.Name = "umschlag-ui"
	app.Version = Version
	app.Author = "Thomas Boerger <thomas@webhippie.de>"
	app.Usage = "A docker distribution management system"

	app.Flags = []cli.Flag{
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

		return nil
	}

	app.Commands = []cli.Command{
		cli.Command{
			Name:  "server",
			Usage: "Start the Umschlag UI",
			Flags: []cli.Flag{
				cli.StringFlag{
					Name:        "host",
					Value:       "http://localhost:9000",
					Usage:       "External access to the UI",
					EnvVar:      "KLEISTER_UI_HOST",
					Destination: &Config.Server.Host,
				},
				cli.StringFlag{
					Name:        "addr",
					Value:       ":9000",
					Usage:       "Address to bind the server",
					EnvVar:      "UMSCHLAG_UI_ADDR",
					Destination: &Config.Server.Addr,
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
				cli.StringFlag{
					Name:        "storage",
					Value:       "storage/",
					Usage:       "Folder for storing assets",
					EnvVar:      "KLEISTER_UI_STORAGE",
					Destination: &Config.Server.Storage,
				},
				cli.StringFlag{
					Name:        "ssl-cert",
					Value:       "",
					Usage:       "Path to SSL cert",
					EnvVar:      "KLEISTER_SSL_CERT",
					Destination: &Config.Server.Cert,
				},
				cli.StringFlag{
					Name:        "ssl-key",
					Value:       "",
					Usage:       "Path to SSL key",
					EnvVar:      "KLEISTER_SSL_KEY",
					Destination: &Config.Server.Key,
				},
				cli.BoolFlag{
					Name:        "ssl-letsencrypt",
					Usage:       "Enable Let's Encrypt SSL",
					EnvVar:      "KLEISTER_SSL_LETSENCRYPT",
					Destination: &Config.Server.LetsEncrypt,
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

				var (
					server *http.Server
				)

				if Config.Server.LetsEncrypt || (Config.Server.Cert != "" && Config.Server.Key != "") {
					curves := []tls.CurveID{
						tls.CurveP521,
						tls.CurveP384,
						tls.CurveP256,
					}

					ciphers := []uint16{
						tls.TLS_ECDHE_ECDSA_WITH_AES_256_GCM_SHA384,
						tls.TLS_ECDHE_RSA_WITH_AES_256_GCM_SHA384,
						tls.TLS_ECDHE_ECDSA_WITH_AES_128_GCM_SHA256,
						tls.TLS_ECDHE_RSA_WITH_AES_128_GCM_SHA256,
					}

					cfg := &tls.Config{
						PreferServerCipherSuites: true,
						MinVersion:               tls.VersionTLS12,
						CurvePreferences:         curves,
						CipherSuites:             ciphers,
					}

					if Config.Server.LetsEncrypt {
						parsed, err := url.Parse(Config.Server.Host)

						if err != nil {
							logrus.Fatal("Failed to parse host name. %s", err)
						}

						certManager := autocert.Manager{
							Prompt:     autocert.AcceptTOS,
							HostPolicy: autocert.HostWhitelist(parsed.Host),
							Cache:      autocert.DirCache(Config.Server.Storage),
						}

						cfg.GetCertificate = certManager.GetCertificate
					} else {
						cert, err := tls.LoadX509KeyPair(
							Config.Server.Cert,
							Config.Server.Key,
						)

						if err != nil {
							logrus.Fatal("Failed to load SSL certificates. %s", err)
						}

						cfg.Certificates = []tls.Certificate{
							cert,
						}
					}

					server = &http.Server{
						Addr:         Config.Server.Addr,
						Handler:      e,
						ReadTimeout:  5 * time.Second,
						WriteTimeout: 10 * time.Second,
						TLSConfig:    cfg,
					}
				} else {
					server = &http.Server{
						Addr:         Config.Server.Addr,
						Handler:      e,
						ReadTimeout:  5 * time.Second,
						WriteTimeout: 10 * time.Second,
					}
				}

				if err := startServer(server); err != nil {
					logrus.Fatal(err)
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

	if err := app.Run(os.Args); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

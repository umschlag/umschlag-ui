package main

import (
	"crypto/tls"
	"fmt"
	"net/http"
	"net/url"
	"os"
	"runtime"
	"strings"
	"time"

	"github.com/Sirupsen/logrus"
	"github.com/gin-contrib/pprof"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"golang.org/x/crypto/acme/autocert"
	"golang.org/x/sync/errgroup"
	"gopkg.in/urfave/cli.v2"
)

//go:generate fileb0x ab0x.yaml

var (
	defaultAddr = ":9000"
)

func main() {
	runtime.GOMAXPROCS(runtime.NumCPU())

	if env := os.Getenv("UMSCHLAG_ENV_FILE"); env != "" {
		godotenv.Load(env)
	}

	app := &cli.App{
		Name:     "umschlag-ui",
		Version:  Version.String(),
		Usage:    "A docker distribution management system",
		Compiled: time.Now(),

		Authors: []*cli.Author{
			{
				Name:  "Thomas Boerger",
				Email: "thomas@webhippie.de",
			},
		},

		Flags: []cli.Flag{
			&cli.BoolFlag{
				Name:        "debug",
				Value:       false,
				Usage:       "Activate debug information",
				EnvVars:     []string{"UMSCHLAG_DEBUG"},
				Destination: &Config.Debug,
				Hidden:      true,
			},
		},

		Before: func(c *cli.Context) error {
			logrus.SetOutput(os.Stdout)

			if Config.Debug {
				logrus.SetLevel(logrus.DebugLevel)
			} else {
				logrus.SetLevel(logrus.InfoLevel)
			}

			return nil
		},

		Commands: []*cli.Command{
			{
				Name:  "server",
				Usage: "Start the Umschlag UI",
				Flags: []cli.Flag{
					&cli.StringFlag{
						Name:        "host",
						Value:       "http://localhost:9000",
						Usage:       "External access to the UI",
						EnvVars:     []string{"UMSCHLAG_UI_HOST"},
						Destination: &Config.Server.Host,
					},
					&cli.StringFlag{
						Name:        "addr",
						Value:       defaultAddr,
						Usage:       "Address to bind the server",
						EnvVars:     []string{"UMSCHLAG_UI_ADDR"},
						Destination: &Config.Server.Addr,
					},
					&cli.StringFlag{
						Name:        "endpoint",
						Value:       "http://localhost:8000",
						Usage:       "URL for the API server",
						EnvVars:     []string{"UMSCHLAG_UI_ENDPOINT"},
						Destination: &Config.Server.Endpoint,
					},
					&cli.StringFlag{
						Name:        "static",
						Value:       "",
						Usage:       "Folder for serving assets",
						EnvVars:     []string{"UMSCHLAG_UI_STATIC"},
						Destination: &Config.Server.Static,
					},
					&cli.StringFlag{
						Name:        "storage",
						Value:       "storage/",
						Usage:       "Folder for storing files",
						EnvVars:     []string{"UMSCHLAG_UI_STORAGE"},
						Destination: &Config.Server.Storage,
					},
					&cli.StringFlag{
						Name:        "cert",
						Value:       "",
						Usage:       "Path to SSL cert",
						EnvVars:     []string{"UMSCHLAG_UI_CERT"},
						Destination: &Config.Server.Cert,
					},
					&cli.StringFlag{
						Name:        "key",
						Value:       "",
						Usage:       "Path to SSL key",
						EnvVars:     []string{"UMSCHLAG_UI_KEY"},
						Destination: &Config.Server.Key,
					},
					&cli.BoolFlag{
						Name:        "letsencrypt",
						Value:       false,
						Usage:       "Enable Let's Encrypt SSL",
						EnvVars:     []string{"UMSCHLAG_UI_LETSENCRYPT"},
						Destination: &Config.Server.LetsEncrypt,
					},
					&cli.BoolFlag{
						Name:    "strict-curves",
						Value:   false,
						Usage:   "Use strict SSL curves",
						EnvVars: []string{"UMSCHLAG_STRICT_CURVES"},
					},
					&cli.BoolFlag{
						Name:    "strict-ciphers",
						Value:   false,
						Usage:   "Use strict SSL ciphers",
						EnvVars: []string{"UMSCHLAG_STRICT_CIPHERS"},
					},
					&cli.BoolFlag{
						Name:        "pprof",
						Value:       false,
						Usage:       "Enable pprof debugger",
						EnvVars:     []string{"UMSCHLAG_UI_PPROF"},
						Destination: &Config.Server.Pprof,
					},
				},
				Action: func(c *cli.Context) error {
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

					if Config.Server.Pprof {
						pprof.Register(
							e,
							&pprof.Options{
								RoutePrefix: "/debug/pprof",
							},
						)
					}

					e.StaticFS(
						"/assets",
						Assets(),
					)

					e.NoRoute(Index)

					if Config.Server.LetsEncrypt || (Config.Server.Cert != "" && Config.Server.Key != "") {
						cfg := &tls.Config{
							PreferServerCipherSuites: true,
							MinVersion:               tls.VersionTLS12,
						}

						if Config.Server.StrictCurves {
							cfg.CurvePreferences = []tls.CurveID{
								tls.CurveP521,
								tls.CurveP384,
								tls.CurveP256,
							}
						}

						if Config.Server.StrictCiphers {
							cfg.CipherSuites = []uint16{
								tls.TLS_ECDHE_ECDSA_WITH_AES_256_GCM_SHA384,
								tls.TLS_ECDHE_RSA_WITH_AES_256_GCM_SHA384,
								tls.TLS_ECDHE_ECDSA_WITH_AES_128_GCM_SHA256,
								tls.TLS_ECDHE_RSA_WITH_AES_128_GCM_SHA256,
							}
						}

						if Config.Server.LetsEncrypt {
							if Config.Server.Addr != defaultAddr {
								logrus.Infof("With Let's Encrypt bind port have been overwritten!")
							}

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

							var (
								g errgroup.Group
							)

							splitAddr := strings.SplitN(Config.Server.Addr, ":", 2)
							logrus.Infof("Starting on %s:80 and %s:443", splitAddr[0], splitAddr[0])

							g.Go(func() error {
								return http.ListenAndServe(
									fmt.Sprintf("%s:80", splitAddr[0]),
									http.HandlerFunc(redirect),
								)
							})

							g.Go(func() error {
								return startServer(&http.Server{
									Addr:         fmt.Sprintf("%s:443", splitAddr[0]),
									Handler:      e,
									ReadTimeout:  5 * time.Second,
									WriteTimeout: 10 * time.Second,
									TLSConfig:    cfg,
								})
							})

							if err := g.Wait(); err != nil {
								logrus.Fatal(err)
							}
						} else {
							logrus.Infof("Starting on %s", Config.Server.Addr)

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

							server := &http.Server{
								Addr:         Config.Server.Addr,
								Handler:      e,
								ReadTimeout:  5 * time.Second,
								WriteTimeout: 10 * time.Second,
								TLSConfig:    cfg,
							}

							if err := startServer(server); err != nil {
								logrus.Fatal(err)
							}
						}
					} else {
						logrus.Infof("Starting on %s", Config.Server.Addr)

						server := &http.Server{
							Addr:         Config.Server.Addr,
							Handler:      e,
							ReadTimeout:  5 * time.Second,
							WriteTimeout: 10 * time.Second,
						}

						if err := startServer(server); err != nil {
							logrus.Fatal(err)
						}
					}

					return nil
				},
			},
		},
	}

	cli.HelpFlag = &cli.BoolFlag{
		Name:    "help",
		Aliases: []string{"h"},
		Usage:   "Show the help, so what you see now",
	}

	cli.VersionFlag = &cli.BoolFlag{
		Name:    "version",
		Aliases: []string{"v"},
		Usage:   "Print the current version of that tool",
	}

	if err := app.Run(os.Args); err != nil {
		os.Exit(1)
	}
}

func redirect(w http.ResponseWriter, req *http.Request) {
	target := "https://" + req.Host + req.URL.Path

	if len(req.URL.RawQuery) > 0 {
		target += "?" + req.URL.RawQuery
	}

	logrus.Debugf("Redirecting to %s", target)
	http.Redirect(w, req, target, http.StatusTemporaryRedirect)
}

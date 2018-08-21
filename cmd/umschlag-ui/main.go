package main

import (
	"os"

	"github.com/joho/godotenv"
	"github.com/umschlag/umschlag-ui/pkg/config"
	"github.com/umschlag/umschlag-ui/pkg/version"
	"gopkg.in/urfave/cli.v2"
)

func main() {
	cfg := config.Load()

	if env := os.Getenv("UMSCHLAG_UI_ENV_FILE"); env != "" {
		godotenv.Load(env)
	}

	app := &cli.App{
		Name:     "umschlag-ui",
		Version:  version.Version.String(),
		Usage:    "docker distribution management system",
		Authors:  authorList(),
		Flags:    globalFlags(cfg),
		Commands: globalCommands(cfg),
	}

	cli.HelpFlag = &cli.BoolFlag{
		Name:    "help",
		Aliases: []string{"h"},
		Usage:   "show the help, so what you see now",
	}

	cli.VersionFlag = &cli.BoolFlag{
		Name:    "version",
		Aliases: []string{"v"},
		Usage:   "print the current version of that tool",
	}

	if err := app.Run(os.Args); err != nil {
		os.Exit(1)
	}
}

func authorList() []*cli.Author {
	return []*cli.Author{
		{
			Name:  "Thomas Boerger",
			Email: "thomas@webhippie.de",
		},
	}
}

func globalFlags(cfg *config.Config) []cli.Flag {
	return []cli.Flag{
		&cli.StringFlag{
			Name:        "log-level",
			Value:       "info",
			Usage:       "set logging level",
			EnvVars:     []string{"UMSCHLAG_UI_LOG_LEVEL"},
			Destination: &cfg.Logs.Level,
		},
		&cli.BoolFlag{
			Name:        "log-pretty",
			Value:       true,
			Usage:       "enable pretty logging",
			EnvVars:     []string{"UMSCHLAG_UI_LOG_PRETTY"},
			Destination: &cfg.Logs.Pretty,
		},
		&cli.BoolFlag{
			Name:        "log-color",
			Value:       true,
			Usage:       "enable colored logging",
			EnvVars:     []string{"UMSCHLAG_UI_LOG_COLOR"},
			Destination: &cfg.Logs.Color,
		},
	}
}

func globalCommands(cfg *config.Config) []*cli.Command {
	return []*cli.Command{
		Server(cfg),
		Health(cfg),
	}
}

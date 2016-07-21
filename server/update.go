package main

import (
	"fmt"
	"os"

	"github.com/sanbornm/go-selfupdate/selfupdate"
)

var (
	updates = "http://dl.webhippie.de/"
)

// Update handles automated binary updates in the background.
func Update() {
	if VersionDev == "dev" {
		fmt.Fprintf(os.Stderr, "Updates are disabled for development versions.\n")
	} else {
		updater := &selfupdate.Updater{
			CurrentVersion: fmt.Sprintf(
				"%d.%d.%d",
				VersionMajor,
				VersionMinor,
				VersionPatch,
			),
			ApiURL:  updates,
			BinURL:  updates,
			DiffURL: updates,
			Dir:     "updates/",
			CmdName: "umschlag-ui",
		}

		go updater.BackgroundRun()
	}
}

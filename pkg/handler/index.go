package handler

import (
	"net/http"

	"github.com/rs/zerolog/log"
	"github.com/umschlag/umschlag-ui/pkg/config"
	"github.com/umschlag/umschlag-ui/pkg/templates"
	"github.com/webhippie/fail"
)

// Index renders the general template on all routes.
func Index(cfg *config.Config) http.HandlerFunc {
	logger := log.With().
		Str("handler", "index").
		Logger()

	return func(w http.ResponseWriter, r *http.Request) {
		if err := templates.Load(cfg).ExecuteTemplate(w, "index.html", vars(cfg)); err != nil {
			logger.Warn().
				Err(err).
				Msg("failed to process index template")

			fail.ErrorPlain(w, fail.Cause(err).Unexpected())
			return
		}
	}
}

func vars(cfg *config.Config) map[string]string {
	return map[string]string{
		"Root":     cfg.Server.Root,
		"Endpoint": cfg.API.Endpoint,
	}
}

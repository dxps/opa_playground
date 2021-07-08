package api

import (
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/dxps/opa_playground/pull_data_example/iam_svc/internal/app"
	"github.com/dxps/opa_playground/pull_data_example/iam_svc/internal/infra/repos"
)

type API struct {
	config         app.Config
	logger         *log.Logger
	appVersion     string
	repos          repos.Repos
	signingKeyPair app.SigningKeyPair
}

func NewAPI(config app.Config, logger *log.Logger, appVersion string, repos repos.Repos, signing app.SigningKeyPair) *API {

	return &API{
		config, logger, appVersion, repos, signing,
	}
}

func (api *API) Serve() error {

	srv := http.Server{
		Addr:         fmt.Sprintf(":%d", api.config.Port),
		Handler:      api.Routes(),
		IdleTimeout:  time.Minute,
		ReadTimeout:  10 * time.Second,
		WriteTimeout: 30 * time.Second,
	}
	api.logger.Printf("Listening for HTTP requests on port %s", srv.Addr)
	return srv.ListenAndServe()
}

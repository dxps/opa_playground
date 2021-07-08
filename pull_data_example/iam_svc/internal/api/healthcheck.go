package api

import (
	"net/http"
)

func (api *API) healthcheckHandler(w http.ResponseWriter, r *http.Request) {

	env := envelope{
		"status": "available",
		"system_info": map[string]string{
			"env_stage": api.config.EnvStage,
			"version":   api.appVersion,
		},
	}

	err := api.writeJSON(w, http.StatusOK, env, nil)
	if err != nil {
		api.serverErrorResponse(w, r, err)
	}
}

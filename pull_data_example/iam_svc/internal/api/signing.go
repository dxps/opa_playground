package api

import (
	"encoding/json"
	"net/http"
)

func (api *API) getSigningPublicKeyHandler(w http.ResponseWriter, r *http.Request) {

	bytes, err := json.Marshal(api.signingKeyPair.PublicKey)
	if err != nil {
		api.serverErrorResponse(w, r, err)
		return
	}

	_ = api.writeJSON(w, http.StatusOK, envelope{
		"signingPublicKey": bytes,
	}, nil)
}

package api

import (
	"net/http"

	"github.com/julienschmidt/httprouter"
)

func (api *API) Routes() *httprouter.Router {

	router := httprouter.New()

	router.NotFound = http.HandlerFunc(api.notFoundResponse)

	// Registering the handlers per methods and URL patterns.

	router.HandlerFunc(http.MethodGet, "/v1/healthcheck", api.healthcheckHandler)

	router.HandlerFunc(http.MethodPost, "/v1/subjects", api.registerUserHandler)
	router.HandlerFunc(http.MethodGet, "/v1/subjects/:id/attributes", api.getSubjectAttributesHandler)
	router.HandlerFunc(http.MethodPost, "/v1/subjects/:id/attributes", api.addSubjectAttributeHandler)

	router.HandlerFunc(http.MethodPost, "/v1/authenticate", api.authenticateHandler)

	router.HandlerFunc(http.MethodGet, "/v1/signing/publickey", api.getSigningPublicKeyHandler)

	return router
}

package api

import (
	"fmt"
	"net/http"
)

func (api *API) logError(r *http.Request, err error) {
	api.logger.Printf("Error '%s' encountered on request '%s %s'", err, r.Method, r.URL.String())
}

func (api *API) errorResponse(w http.ResponseWriter, r *http.Request, status int, message interface{}) {
	env := envelope{"error": message}

	err := api.writeJSON(w, status, env, nil)
	if err != nil {
		api.logError(r, err)
		w.WriteHeader(500)
	}
}

func (api *API) serverErrorResponse(w http.ResponseWriter, r *http.Request, err error) {
	api.logError(r, err)

	message := "the server encountered a problem and could not process your request"
	api.errorResponse(w, r, http.StatusInternalServerError, message)
}

func (api *API) notFoundResponse(w http.ResponseWriter, r *http.Request) {
	message := "the requested resource could not be found"
	api.errorResponse(w, r, http.StatusNotFound, message)
}

func (api *API) methodNotAllowedResponse(w http.ResponseWriter, r *http.Request) {
	message := fmt.Sprintf("the %s method is not supported for this resource", r.Method)
	api.errorResponse(w, r, http.StatusMethodNotAllowed, message)
}

func (api *API) badRequestResponse(w http.ResponseWriter, r *http.Request, err error) {
	api.errorResponse(w, r, http.StatusBadRequest, err.Error())
}

func (app *API) failedValidationResponse(w http.ResponseWriter, r *http.Request, errors map[string]string) {
	app.errorResponse(w, r, http.StatusUnprocessableEntity, errors)
}

func (app *API) editConflictResponse(w http.ResponseWriter, r *http.Request) {
	message := "unable to update the record due to an edit conflict, please try again"
	app.errorResponse(w, r, http.StatusConflict, message)
}

func (app *API) rateLimitExceededResponse(w http.ResponseWriter, r *http.Request) {
	message := "rate limit exceeded"
	app.errorResponse(w, r, http.StatusTooManyRequests, message)
}

func (app *API) invalidCredentialsResponse(w http.ResponseWriter, r *http.Request) {
	message := "invalid authentication credentials"
	app.errorResponse(w, r, http.StatusUnauthorized, message)
}

func (app *API) invalidAuthenticationTokenResponse(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("WWW-Authenticate", "Bearer")

	message := "invalid or missing authentication token"
	app.errorResponse(w, r, http.StatusUnauthorized, message)
}

func (app *API) authenticationRequiredResponse(w http.ResponseWriter, r *http.Request) {
	message := "you must be authenticated to access this resource"
	app.errorResponse(w, r, http.StatusUnauthorized, message)
}

func (app *API) inactiveAccountResponse(w http.ResponseWriter, r *http.Request) {
	message := "your user account must be activated to access this resource"
	app.errorResponse(w, r, http.StatusForbidden, message)
}

func (app *API) notPermittedResponse(w http.ResponseWriter, r *http.Request) {
	message := "your user account doesn't have the necessary permissions to access this resource"
	app.errorResponse(w, r, http.StatusForbidden, message)
}

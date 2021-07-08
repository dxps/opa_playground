package api

import (
	"errors"
	"net/http"
	"time"

	"github.com/dxps/opa_playground/pull_data_example/iam_svc/internal/app"
	"github.com/dxps/opa_playground/pull_data_example/iam_svc/internal/domain"
	"github.com/dxps/opa_playground/pull_data_example/iam_svc/internal/validator"
	"github.com/pascaldekloe/jwt"
)

func (api *API) authenticateHandler(w http.ResponseWriter, r *http.Request) {

	var input struct {
		Email    string `json:"email"`
		Password string `json:"password"`
	}

	if err := api.readJSON(w, r, &input); err != nil {
		api.badRequestResponse(w, r, err)
		return
	}

	v := validator.New()

	domain.ValidateEmail(v, input.Email)
	domain.ValidatePasswordPlaintext(v, input.Password)

	if !v.Valid() {
		api.failedValidationResponse(w, r, v.Errors)
		return
	}

	subj, err := api.repos.Subjects.GetByEmail(input.Email)
	if err != nil {
		switch {
		case errors.Is(err, app.ErrRecordNotFound):
			api.invalidCredentialsResponse(w, r)
		default:
			api.serverErrorResponse(w, r, err)
		}
		return
	}

	match, err := subj.Password.Matches(input.Password)
	if err != nil {
		api.serverErrorResponse(w, r, err)
		return
	}
	if !match {
		api.invalidCredentialsResponse(w, r)
		return
	}

	now := time.Now()
	var claims jwt.Claims
	claims.Subject = subj.EID
	claims.Issued = jwt.NewNumericTime(now)
	claims.Issuer = "iam.service"
	claims.NotBefore = jwt.NewNumericTime(now)
	claims.Expires = jwt.NewNumericTime(now.Add(1 * time.Hour))
	claims.Audiences = []string{"anyone"}

	jwtBytes, err := claims.ECDSASign(jwt.ES256, api.signingKeyPair.PrivateKey)
	if err != nil {
		api.serverErrorResponse(w, r, err)
		return
	}
	err = api.writeJSON(w, http.StatusOK, envelope{"access_token": string(jwtBytes)}, nil)
	if err != nil {
		api.serverErrorResponse(w, r, err)
	}
}

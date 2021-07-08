package api

import (
	"errors"
	"net/http"

	"github.com/dxps/opa_playground/pull_data_example/iam_svc/internal/app"
	"github.com/dxps/opa_playground/pull_data_example/iam_svc/internal/domain"
	"github.com/dxps/opa_playground/pull_data_example/iam_svc/internal/validator"
)

func (api *API) registerUserHandler(w http.ResponseWriter, r *http.Request) {

	// The expected data from request payload (used here as an anonymous struct).
	var input struct {
		Name     string `json:"name"`
		Email    string `json:"email"`
		Password string `json:"password"`
	}
	err := api.readJSON(w, r, &input)
	if err != nil {
		api.badRequestResponse(w, r, err)
		return
	}

	subj := &domain.Subject{
		Name:   input.Name,
		Email:  input.Email,
		Active: true, // TODO: Later to activate subject through email confirmation.
	}

	err = subj.Password.Set(input.Password)
	if err != nil {
		api.serverErrorResponse(w, r, err)
		return
	}

	v := validator.New()

	if domain.ValidateSubject(v, subj); !v.Valid() {
		api.failedValidationResponse(w, r, v.Errors)
		return
	}

	err = api.repos.Subjects.Add(subj)

	if err != nil {
		switch {
		case errors.Is(err, app.ErrDuplicateEmail):
			v.AddError("email", "a user with this email address already exists")
			api.failedValidationResponse(w, r, v.Errors)
		default:
			api.serverErrorResponse(w, r, err)
		}
		return
	}

	err = api.writeJSON(w, http.StatusCreated, subj, nil)
	if err != nil {
		api.serverErrorResponse(w, r, err)
	}
}

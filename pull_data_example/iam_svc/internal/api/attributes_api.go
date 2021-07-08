package api

import (
	"net/http"

	"github.com/dxps/opa_playground/pull_data_example/iam_svc/internal/app"
	"github.com/dxps/opa_playground/pull_data_example/iam_svc/internal/domain"
)

func (api *API) getSubjectAttributesHandler(w http.ResponseWriter, r *http.Request) {

	// Getting the subject's external ID provided as URL param.
	eid, err := api.readUUIDParam(r)
	if err != nil {
		api.logger.Print("getSubjectAttributesHandler > readUUIDParam (id) error: ", err)
		api.badRequestResponse(w, r, app.ErrSubjectEIDInvalid)
		return
	}

	id, err := api.repos.Subjects.GetSubjectIDByEID(eid.String())
	if err != nil {
		api.notFoundResponse(w, r)
		return
	}
	api.logger.Printf("getSubjectAttributesHandler > subjectEID=%v", eid)

	attrs, err := api.repos.Attributes.GetAllAttributesBySubjectID(*id)
	if err != nil {
		api.serverErrorResponse(w, r, err)
		return
	}
	err = api.writeJSON(w, http.StatusOK, attrs, nil)
	if err != nil {
		api.serverErrorResponse(w, r, err)
	}
}

func (api *API) addSubjectAttributeHandler(w http.ResponseWriter, r *http.Request) {

	// Getting the subject's external ID provided as URL param.
	eid, err := api.readUUIDParam(r)
	if err != nil {
		api.logger.Print("getSubjectAttributesHandler > readUUIDParam (id) error: ", err)
		api.badRequestResponse(w, r, app.ErrSubjectEIDInvalid)
		return
	}

	id, err := api.repos.Subjects.GetSubjectIDByEID(eid.String())
	if err != nil {
		api.notFoundResponse(w, r)
		return
	}

	var input struct {
		Name  string `json:"name"`
		Value string `json:"value"`
	}
	err = api.readJSON(w, r, &input)
	if err != nil {
		api.badRequestResponse(w, r, err)
		return
	}

	attr := domain.Attribute{
		OwnerID:   *id,
		OwnerType: domain.OWNER_TYPE_SUBJECT,
		Name:      input.Name,
		Value:     input.Value,
	}
	err = api.repos.Attributes.Add(attr)
	if err != nil {
		api.serverErrorResponse(w, r, err)
		return
	}

	api.respondStatus(w, http.StatusCreated)
}

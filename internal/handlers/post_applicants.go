package handlers

import (
	"encoding/json"
	"errors"
	"net/http"

	"github.com/mattisongjj/Go-FAS-API/api"
	"github.com/mattisongjj/Go-FAS-API/internal/tools"
	log "github.com/sirupsen/logrus"
)

func PostApplicants(w http.ResponseWriter, r *http.Request) {
	var newApplicant = tools.Applicant{}
	// Decode and validate data
	err := json.NewDecoder(r.Body).Decode(&newApplicant)
	if err != nil {
		log.Error(err)
		api.RequestErrorHandler(w, err)
		return
	}
	if newApplicant.Id == "" || newApplicant.Name == "" {
		api.RequestErrorHandler(w, errors.New("applicant id and name cannot be empty"))
		return
	}

	database, err := tools.NewDatabase()
	if err != nil {
		api.InternalErrorHandler(w)
		return
	}

	applicant := (*database).CreateApplicant(&newApplicant)
	if applicant == nil {
		api.RequestErrorHandler(w, errors.New("applicant already exists"))
		return
	}

	var response = api.CreateApplicantResponse{
		Code:      http.StatusOK,
		Applicant: *applicant,
	}

	err = json.NewEncoder(w).Encode(response)
	if err != nil {
		log.Error(err)
		api.InternalErrorHandler(w)
	}

}

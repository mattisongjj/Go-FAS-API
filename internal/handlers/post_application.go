package handlers

import (
	"encoding/json"
	"errors"
	"net/http"

	"github.com/mattisongjj/Go-FAS-API/api"
	"github.com/mattisongjj/Go-FAS-API/internal/tools"
	log "github.com/sirupsen/logrus"
)

func PostApplication(w http.ResponseWriter, r *http.Request) {
	var newApplication tools.Application
	// Decode and validate data
	err := json.NewDecoder(r.Body).Decode(&newApplication)
	if err != nil {
		log.Error(err)
		api.RequestErrorHandler(w, err)
		return
	}
	if newApplication.ID == "" || newApplication.ApplicantID == "" || newApplication.SchemeID == "" {
		api.RequestErrorHandler(w, errors.New("application id, applicant id, and scheme id cannot be empty"))
		return
	}

	database, err := tools.NewDatabase()
	if err != nil {
		api.InternalErrorHandler(w)
		return
	}

	application := (*database).CreateApplication(&newApplication)
	if application == nil {
		api.RequestErrorHandler(w, errors.New("application already exists"))
		return
	}

	var response = api.CreateApplicationResponse{
		Code:        http.StatusOK,
		Application: *application,
	}

	err = json.NewEncoder(w).Encode(response)
	if err != nil {
		log.Error(err)
		api.InternalErrorHandler(w)
	}
}

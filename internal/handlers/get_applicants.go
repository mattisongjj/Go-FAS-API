package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/mattisongjj/Go-FAS-API/api"
	"github.com/mattisongjj/Go-FAS-API/internal/tools"
	log "github.com/sirupsen/logrus"
)

func GetApplicants(w http.ResponseWriter, r *http.Request) {
	var err error

	var database *tools.DatabaseInterface
	database, err = tools.NewDatabase()
	if err != nil {
		api.InternalErrorHandler(w)
		return
	}

	var applicants []tools.Applicant
	applicants = (*database).GetApplicants()

	var response = api.ApplicantResponse{
		Code:       http.StatusOK,
		Applicants: applicants,
	}

	w.Header().Set("Content-Type", "application/json")
	err = json.NewEncoder(w).Encode(response)
	if err != nil {
		log.Error(err)
		api.InternalErrorHandler(w)
		return
	}
}

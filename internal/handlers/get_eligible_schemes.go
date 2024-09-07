package handlers

import (
	"encoding/json"
	"errors"
	"net/http"

	"github.com/mattisongjj/Go-FAS-API/api"
	"github.com/mattisongjj/Go-FAS-API/internal/tools"
	log "github.com/sirupsen/logrus"
)

func GetEligibleSchemes(w http.ResponseWriter, r *http.Request) {
	applicantID := r.URL.Query().Get("applicant")
	if applicantID == "" {
		api.RequestErrorHandler(w, errors.New("applicant id cannot be empty"))
		return
	}

	database, err := tools.NewDatabase()
	if err != nil {
		api.InternalErrorHandler(w)
		return
	}

	eligibleSchemes := (*database).GetEligibleSchemes(applicantID)
	if eligibleSchemes == nil {
		api.RequestErrorHandler(w, errors.New("no eligible schemes found or applicant does not exist"))
		return
	}

	var response = api.GetEligibleSchemesResponse{
		Code:    http.StatusOK,
		Schemes: eligibleSchemes,
	}

	err = json.NewEncoder(w).Encode(response)
	if err != nil {
		log.Error(err)
		api.InternalErrorHandler(w)
	}
}

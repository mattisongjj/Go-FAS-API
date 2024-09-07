package handlers

import (
	"encoding/json"
	"errors"
	"net/http"

	"github.com/mattisongjj/Go-FAS-API/api"
	"github.com/mattisongjj/Go-FAS-API/internal/tools"
	log "github.com/sirupsen/logrus"
)

func PostSchemes(w http.ResponseWriter, r *http.Request) {
	var newScheme tools.Scheme
	// Decode and validate data
	err := json.NewDecoder(r.Body).Decode(&newScheme)
	if err != nil {
		log.Error(err)
		api.RequestErrorHandler(w, err)
		return
	}
	if newScheme.ID == "" || newScheme.Name == "" {
		api.RequestErrorHandler(w, errors.New("scheme id and name cannot be empty"))
		return
	}

	database, err := tools.NewDatabase()
	if err != nil {
		api.InternalErrorHandler(w)
		return
	}

	scheme := (*database).CreateScheme(&newScheme)
	if scheme == nil {
		api.RequestErrorHandler(w, errors.New("scheme already exists"))
		return
	}

	var response = api.CreateSchemeResponse{
		Code:   http.StatusOK,
		Scheme: *scheme,
	}

	err = json.NewEncoder(w).Encode(response)
	if err != nil {
		log.Error(err)
		api.InternalErrorHandler(w)
	}
}

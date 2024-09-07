package api

import (
	"encoding/json"
	"net/http"

	"github.com/mattisongjj/Go-FAS-API/internal/tools"
)

// Applicant responses
type GetApplicantResponse struct {
	Code       int
	Applicants []tools.Applicant
}

type CreateApplicantResponse struct {
	Code      int
	Applicant tools.Applicant
}

// Scheme responses
type GetSchemeResponse struct {
	Code    int
	Schemes []tools.Scheme
}

type CreateSchemeResponse struct {
	Code   int
	Scheme tools.Scheme
}

// Application responses
type GetApplicationResponse struct {
	Code         int
	Applications []tools.Application
}

type CreateApplicationResponse struct {
	Code        int
	Application tools.Application
}

// Eligible schemes
type GetEligibleSchemesResponse struct {
	Code    int
	Schemes []tools.Scheme
}

// Error response
type Error struct {

	// Error code
	Code int

	// Error message
	Message string
}

func writeError(w http.ResponseWriter, message string, code int) {
	resp := Error{
		Code:    code,
		Message: message,
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)

	json.NewEncoder(w).Encode(resp)
}

var (
	RequestErrorHandler = func(w http.ResponseWriter, err error) {
		writeError(w, err.Error(), http.StatusBadRequest)
	}
	InternalErrorHandler = func(w http.ResponseWriter) {
		writeError(w, "An Unexpected Error Occured.", http.StatusInternalServerError)
	}
)

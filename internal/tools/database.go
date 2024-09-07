package tools

import (
	log "github.com/sirupsen/logrus"
)

// Enums
type EmploymentStatus string

const (
	Employed   EmploymentStatus = "employed"
	Unemployed EmploymentStatus = "unemployed"
)

type Relation string

const (
	RelationMother   Relation = "mother"
	RelatonFather    Relation = "father"
	RelationSpouse   Relation = "spouse"
	RelationSister   Relation = "sister"
	RelationBrother  Relation = "brother"
	RelationDaughter Relation = "daughter"
	RelationSon      Relation = "son"
	RelationOhter    Relation = "other"
)

// Database models
type LoginDetails struct {
	AuthToken string
	Username  string
}

type Applicant struct {
	Id               string            `json:"id"`
	Name             string            `json:"name"`
	EmploymentStatus EmploymentStatus  `json:"employment_status"`
	Sex              string            `json:"sex"`
	DateOfBirth      string            `json:"date_of_birth"`
	Household        []HouseholdMember `json:"household"`
}

type HouseholdMember struct {
	ID               string           `json:"id"`
	ApplicantID      string           `json:"applicant_id"`
	Name             string           `json:"name"`
	EmploymentStatus EmploymentStatus `json:"employment_status"`
	Sex              string           `json:"sex"`
	DateOfBirth      string           `json:"date_of_birth"`
	Relation         Relation         `json:"relation"`
}

type Scheme struct {
	ID       string                 `json:"id"`
	Name     string                 `json:"name"`
	Criteria map[string]interface{} `json:"criteria"`
	Benefits []Benefit              `json:"benefits"`
}

type Benefit struct {
	ID     string  `json:"id"`
	Name   string  `json:"name"`
	Amount float64 `json:"amount"`
}

type Application struct {
	ID          string `json:"id"`
	ApplicantID string `json:"applicant_id"`
	SchemeID    string `json:"scheme_id"`
	Status      string `json:"status"`
}

type DatabaseInterface interface {
	GetUserLoginDetails(username string) *LoginDetails
	GetApplicants() []Applicant
	CreateApplicant(*Applicant) *Applicant
	GetSchemes() []Scheme
	CreateScheme(*Scheme) *Scheme
	GetApplications() []Application
	CreateApplication(*Application) *Application
	GetEligibleSchemes(applicantID string) []Scheme
	SetupDatabase() error
}

func NewDatabase() (*DatabaseInterface, error) {
	var database DatabaseInterface = &mockDB{}

	var err error = database.SetupDatabase()
	if err != nil {
		log.Error(err)
		return nil, err
	}
	return &database, nil
}

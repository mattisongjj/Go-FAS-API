package tools

import "time"

type mockDB struct{}

var mockLoginDetails = map[string]LoginDetails{
	"john": {
		AuthToken: "123456ABC",
		Username:  "john",
	},
	"doe": {
		AuthToken: "ABC1234",
		Username:  "doe",
	},
}

var mockApplicantDetails = map[string]Applicant{
	"1": {
		Id:               "1",
		Name:             "John Doe",
		EmploymentStatus: Unemployed,
		Sex:              "male",
		DateOfBirth:      time.Date(1985, time.January, 1, 0, 0, 0, 0, time.UTC),
		Household:        []HouseholdMember{},
	},
	"2": {
		Id:               "2",
		Name:             "Jane Smith",
		EmploymentStatus: Employed,
		Sex:              "female",
		DateOfBirth:      time.Date(1990, time.February, 10, 0, 0, 0, 0, time.UTC),
		Household: []HouseholdMember{
			{
				ID:               "3",
				ApplicantID:      "2",
				Name:             "Child One",
				EmploymentStatus: Unemployed,
				Sex:              "female",
				DateOfBirth:      time.Date(2015, time.March, 15, 0, 0, 0, 0, time.UTC),
				Relation:         RelationDaughter,
			},
		},
	},
}

func (d *mockDB) GetUserLoginDetails(username string) *LoginDetails {
	var clientData = LoginDetails{}
	clientData, ok := mockLoginDetails[username]
	if !ok {
		return nil
	}
	return &clientData
}

func (d *mockDB) GetApplicants() []Applicant {
	applicants := make([]Applicant, 0, len(mockApplicantDetails))

	for _, applicant := range mockApplicantDetails {
		applicants = append(applicants, applicant)
	}

	return applicants
}

func (d *mockDB) SetupDatabase() error {
	return nil // Do nothing, for mock db
}

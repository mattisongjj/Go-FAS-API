package tools

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
		DateOfBirth:      "1985-01-01",
		Household:        []HouseholdMember{},
	},
	"2": {
		Id:               "2",
		Name:             "Jane Smith",
		EmploymentStatus: Employed,
		Sex:              "female",
		DateOfBirth:      "1999-02-02",
		Household: []HouseholdMember{
			{
				ID:               "3",
				ApplicantID:      "2",
				Name:             "Child One",
				EmploymentStatus: Unemployed,
				Sex:              "female",
				DateOfBirth:      "2015-03-15",
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

func (d *mockDB) CreateApplicant(applicant *Applicant) *Applicant {
	if _, exists := mockApplicantDetails[applicant.Id]; exists {
		return nil
	}
	mockApplicantDetails[applicant.Id] = *applicant
	return applicant
}

func (d *mockDB) SetupDatabase() error {
	return nil // Do nothing, for mock db
}

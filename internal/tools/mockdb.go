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
		EmploymentStatus: Unemployed,
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

var mockSchemeDetails = map[string]Scheme{
	"1": {
		ID:   "1",
		Name: "Retrenchment Assistance Scheme",
		Criteria: map[string]interface{}{
			"employment_status": Unemployed,
		},
		Benefits: []Benefit{
			{
				ID:       "1",
				SchemeID: "1",
				Name:     "SkillsFuture Credits",
				Amount:   500.00,
			},
			{
				ID:       "2",
				SchemeID: "1",
				Name:     "CDC Vouchers",
				Amount:   100.00,
			},
		},
	},
	"2": {
		ID:   "2",
		Name: "Retrenchment Assistance Scheme (families)",
		Criteria: map[string]interface{}{
			"employment_status": Unemployed,
			"has_children":      true,
		},
		Benefits: []Benefit{
			{
				ID:       "3",
				SchemeID: "2",
				Name:     "School Meal Vouchers",
				Amount:   200.00,
			},
		},
	},
}

var mockApplicationDetails = map[string]Application{
	"1": {
		ID:          "1",
		ApplicantID: "1",
		SchemeID:    "1",
		Status:      "Approved",
	},
	"2": {
		ID:          "2",
		ApplicantID: "2",
		SchemeID:    "2",
		Status:      "Pending",
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

func (d *mockDB) GetSchemes() []Scheme {
	schemes := make([]Scheme, 0, len(mockSchemeDetails))

	for _, scheme := range mockSchemeDetails {
		schemes = append(schemes, scheme)
	}

	return schemes
}

func (d *mockDB) CreateScheme(scheme *Scheme) *Scheme {
	if _, exists := mockSchemeDetails[scheme.ID]; exists {
		return nil
	}
	mockSchemeDetails[scheme.ID] = *scheme
	return scheme
}

func (d *mockDB) GetApplications() []Application {
	applications := make([]Application, 0, len(mockApplicationDetails))

	for _, application := range mockApplicationDetails {
		applications = append(applications, application)
	}

	return applications
}

func (d *mockDB) CreateApplication(application *Application) *Application {
	if _, exists := mockApplicationDetails[application.ID]; exists {
		return nil
	}
	mockApplicationDetails[application.ID] = *application
	return application
}

func (d *mockDB) GetEligibleSchemes(applicantID string) []Scheme {
	applicant, exists := mockApplicantDetails[applicantID]
	if !exists {
		return nil
	}

	var eligibleSchemes []Scheme

	for _, scheme := range mockSchemeDetails {
		if isEligible(applicant, scheme) {
			eligibleSchemes = append(eligibleSchemes, scheme)
		}
	}

	return eligibleSchemes
}

// Given an applicant and scheme, check if applicant qualifies for scheme
func isEligible(applicant Applicant, scheme Scheme) bool {
	for key, value := range scheme.Criteria {
		switch key {
		case "employment_status":
			if applicant.EmploymentStatus != value {
				return false
			}
		case "has_children":
			hasChildren := false
			for _, member := range applicant.Household {
				if member.Relation == RelationSon || member.Relation == RelationDaughter {
					hasChildren = true
					break
				}
			}
			if hasChildren != value.(bool) {
				return false
			}
			// Add other criteria here
		}
	}

	return true
}

func (d *mockDB) SetupDatabase() error {
	return nil // Do nothing, for mock db
}

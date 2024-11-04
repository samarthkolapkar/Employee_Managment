package models

type Employee struct {
	FirstName          string            `json:"first_name" validate:"required"`
	MiddleName         string            `json:"middle_name" validate:"required"`
	LastName           string            `json:"last_name" validate:"required"`
	PanNumber          string            `json:"pan_number" validate:"required"`
	Contact            ContactDetails    `json:"contact"`
	IsExperienced      bool              `json:"is_experienced"`
	ExperiencedDetails ExperienceDetails `json:"experienced_details"`
	Address            AddressDetails    `json:"address"`
	Age                int
}

type ContactDetails struct {
	ContactNumber string `json:"contact_number" validate:"required"`
	Email         string `json:"email_id" validate:"required"`
}

type ExperienceDetails struct {
	Years   int            `json:"years"`
	Details CompanyDetails `json:"details"`
}
type CompanyDetails struct {
	CompanyName string   `json:"Company_name"`
	FromDate    string   `json:"from"`
	ToDate      string   `json:"to"`
	Description []string `json:"description"`
}
type AddressDetails struct {
	Street string `json:"street" validate:"required"`
	City   string `json:"city" validate:"required"`
	State  string `json:"state" validate:"required"`
	Zip    int    `json:"zip" validate:"required"`
}

type JSONresponse struct {
	Status  string `json:"status"`
	Data    any    `json:"data"`
	Message any    `json:"message"`
}

type EmployeeList struct {
	Fields []string `json:"fields"`
	Search string   `json:"search"`
	Status string   `json:"status"`
}

type EmployeeGet struct {
	Id     int64  `json:"id"`
	Status string `json:"status"`
}

type ApproveEmployee struct {
	Id int64 `json:"id"`
}

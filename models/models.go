package models

type Employee struct {
	FirstName  string         `json:"first_name" validate:"required"`
	MiddleName string         `json:"middle_name" validate:"required"`
	LastName   string         `json:"last_name" validate:"required"`
	PanNumber  string         `json:"pan_number" validate:"required"`
	Address    AddressDetails `json:"address"`
	Age        int
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
}

type EmployeeForward struct {
	Id int64 `json:id`
}

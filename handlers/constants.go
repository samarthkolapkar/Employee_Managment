package handlers

const (
	EMPTY_STRING      = ""
	STATUS_DRAFT      = "DRAFT"
	FIRST_NAME        = "first_name"
	MIDDLE_NAME       = "middle_name"
	LAST_NAME         = "last_name"
	EXPERIENCED_YEARS = "experienced_details.years"
	CONTACT_EMAIL     = "contact.email_id"
	CONTACT_NUMBER    = "contact.contact_number"
	PAN_NUMBER        = "pan_number"
	PAN_REGEX         = `^[A-Z]{5}[0-9]{4}[A-Z]{1}$`
	EMAIL_REGEX       = "^[a-zA-Z0-9._%+-]+@[a-zA-Z0-9.-]+[a-zA-Z]{2,}$"
)
const (
	INAVALID_DATA = 501
)

var statusArray = []string{"APPROVED"}

// var statusMakerArray = []string{"DRAFT"}

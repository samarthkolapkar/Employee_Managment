package handlers

import (
	"employee/models"
	"encoding/json"
	"log"
	"net/http"
	"regexp"
	"strings"

	"github.com/remiges-tech/alya/wscutils"
)

var validationErrors []wscutils.ErrorMessage

// This function is used to validate the request
func validateEmployee(employee models.Employee) []wscutils.ErrorMessage {
	log.Println("started custom validation of the employee request while validating the employee")
	if strings.TrimSpace(employee.FirstName) == EMPTY_STRING {
		validationErrors = append(validationErrors, BuildErrorMessages(INAVALID_DATA, FIRST_NAME, "First name cannot be empty"))
	}
	if strings.TrimSpace(employee.MiddleName) == EMPTY_STRING {
		validationErrors = append(validationErrors, BuildErrorMessages(INAVALID_DATA, MIDDLE_NAME, "Middle name cannot be empty"))
	}
	if strings.TrimSpace(employee.LastName) == EMPTY_STRING {
		validationErrors = append(validationErrors, BuildErrorMessages(INAVALID_DATA, LAST_NAME, "Last name cannot be empty"))
	}
	if employee.PanNumber == EMPTY_STRING {
		validationErrors = append(validationErrors, BuildErrorMessages(INAVALID_DATA, PAN_NUMBER, "PAN number should not be empty"))
	} else {
		if !validatePan(employee.PanNumber) {
			validationErrors = append(validationErrors, BuildErrorMessages(INAVALID_DATA, PAN_NUMBER, "PAN number should be valid"))
		}
	}
	if employee.IsExperienced {
		if employee.ExperiencedDetails.Years <= 0 {
			validationErrors = append(validationErrors, BuildErrorMessages(INAVALID_DATA, EXPERIENCED_YEARS, "if you are experienced then experienced years should not 0 or less than 0"))
		}
	}
	if employee.Contact.Email == EMPTY_STRING {
		validationErrors = append(validationErrors, BuildErrorMessages(INAVALID_DATA, CONTACT_EMAIL, "Email field cannot be empty"))
	} else {
		if !validEmailAddress(employee.Contact.Email) {
			validationErrors = append(validationErrors, BuildErrorMessages(INAVALID_DATA, CONTACT_EMAIL, "Email is not valid email address"))
		}
	}
	if len(employee.Contact.ContactNumber) < 10 || len(employee.Contact.ContactNumber) > 10 {
		validationErrors = append(validationErrors, BuildErrorMessages(INAVALID_DATA, CONTACT_NUMBER, "Length of mobile number should be 10"))
	}
	return validationErrors
}

func BuildErrorMessages(errcode int, field string, message string) wscutils.ErrorMessage {
	errormessage := wscutils.BuildErrorMessage(errcode, field, message)
	return errormessage

}

// This function is used to generate the json response
func writeResponse(writer http.ResponseWriter, code int, response models.JSONresponse) {
	writer.Header().Set("Content-Type", "application/json")
	writer.WriteHeader(code)
	json.NewEncoder(writer).Encode(response)
}

// This function is used to validate the pan
func validatePan(pan string) bool {
	pattern := regexp.MustCompile(PAN_REGEX)
	flag := pattern.MatchString(pan)
	return flag
}

// This function is used to validate the email_id
func validEmailAddress(email string) bool {
	pattern := regexp.MustCompile(EMAIL_REGEX)
	flag := pattern.MatchString(email)
	return flag
}

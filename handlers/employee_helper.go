package handlers

import (
	"employee/models"
	"encoding/json"
	"net/http"
	"strings"

	"github.com/remiges-tech/alya/wscutils"
)

var validationErrors []wscutils.ErrorMessage

func validateEmployee(employee models.Employee) []wscutils.ErrorMessage {
	if strings.TrimSpace(employee.FirstName) == EMPTY_STRING {
		validationErrors = append(validationErrors, BuildErrorMessages(11, "first_name", "First name cannot be empty"))
		return validationErrors
	}
	if strings.TrimSpace(employee.MiddleName) == EMPTY_STRING {
		validationErrors = append(validationErrors, BuildErrorMessages(11, "middle_name", "Middle name cannot be empty"))
		return validationErrors
	}
	if strings.TrimSpace(employee.LastName) == EMPTY_STRING {
		validationErrors = append(validationErrors, BuildErrorMessages(11, "last_name", "Last name cannot be empty"))
		return validationErrors
	}
	return validationErrors
}

func BuildErrorMessages(errcode int, field string, message string) wscutils.ErrorMessage {
	errormessage := wscutils.BuildErrorMessage(errcode, field, message)
	return errormessage

}
func writeResponse(writer http.ResponseWriter, code int, response models.JSONresponse) {
	writer.Header().Set("Content-Type", "application/json")
	writer.WriteHeader(code)
	json.NewEncoder(writer).Encode(response)

}

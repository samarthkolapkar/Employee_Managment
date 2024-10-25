package handlers

import (
	"employee/models"
	"encoding/json"
	"fmt"
	"net/http"
)

func HandleEmployeeForward(writer http.ResponseWriter, read http.Request) {
	var forwardRequest models.EmployeeForward

	err := json.NewDecoder(read.Body).Decode(&forwardRequest)
	if err != nil {
		resp := models.JSONresponse{
			Status:  "error",
			Message: fmt.Errorf("error while binding the json request %v", err),
		}
		writeResponse(writer, http.StatusBadRequest, resp)
	}

}

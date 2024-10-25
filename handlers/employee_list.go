package handlers

import (
	"employee/connection"
	"employee/generated"
	"employee/models"
	"encoding/json"
	"fmt"
	"net/http"
)

func HandleEmployeeList(writer http.ResponseWriter, read *http.Request) {
	var listrequest *models.EmployeeList
	db := connection.GetDB()

	err := json.NewDecoder(read.Body).Decode(&listrequest)
	if err != nil {
		resp := models.JSONresponse{
			Status:  "error",
			Message: fmt.Sprintf("Failed to parse the JSON %v", err),
		}
		writeResponse(writer, http.StatusBadRequest, resp)
	}
	queries := generated.New(db)
	list, err := queries.EmployeeList(read.Context(), listrequest.Search)
	if err != nil {
		// http.Error(writer, err.Error(), http.StatusInternalServerError)
		resp := models.JSONresponse{
			Status:  "error",
			Message: fmt.Sprintf("Error while excuting the query %v", err),
		}
		writeResponse(writer, http.StatusInternalServerError, resp)
		return
	}
	resp := models.JSONresponse{
		Status:  "success",
		Message: nil,
		Data:    list,
	}
	writeResponse(writer, http.StatusOK, resp)
}

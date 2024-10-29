package handlers

import (
	"context"
	"database/sql"
	"employee/connection"
	"employee/generated"
	"employee/models"
	"encoding/json"
	"fmt"
	"net/http"
	// "google.golang.org/grpc/profiling/service"
)

func HandleAddEmployee(writer http.ResponseWriter, read *http.Request) {
	var request models.Employee
	connc := connection.GetDB()

	//bind to the body of request to the request structure
	err := json.NewDecoder(read.Body).Decode(&request)
	if err != nil {
		fmt.Println("Error while binding the json!!")
		resp := models.JSONresponse{
			Status:  "error",
			Message: fmt.Sprintf("Failed to bind the json %v", err),
		}
		writeResponse(writer, http.StatusBadRequest, resp)
	}
	validationErrors := validateEmployee(request)
	if len(validationErrors) > 0 {
		resp := models.JSONresponse{
			Status:  "error",
			Message: validationErrors,
		}
		writeResponse(writer, http.StatusBadRequest, resp)
	} else {
		id, err := insertEmployee(request, connc)
		if err != nil {
			// http.Error(writer, err.Error(), http.StatusInternalServerError)
			resp := models.JSONresponse{
				Status:  "error",
				Message: err.Error(),
			}
			writeResponse(writer, http.StatusInternalServerError, resp)
		}
		resp := models.JSONresponse{
			Status:  "success",
			Data:    id,
			Message: "Employee added successfully",
		}
		writeResponse(writer, http.StatusOK, resp)
	}

}

func insertEmployee(request models.Employee, db *sql.DB) (int64, error) {
	querier := generated.New(db)
	var writer http.ResponseWriter
	employeeData, err := json.Marshal(&request)
	if err != nil {
		resp := models.JSONresponse{
			Status:  "error",
			Message: fmt.Sprintf("Error while marshalling %v", err),
			Data:    nil,
		}
		writeResponse(writer, http.StatusInternalServerError, resp)
	}
	id, err := querier.InsertEmployee(context.Background(), generated.InsertEmployeeParams{
		EmployeeData: employeeData,
		Status:       STATUS_DRAFT,
	})

	if err != nil {
		return int64(0), fmt.Errorf("error inserting employee: %v", err)
	}
	return int64(id), nil
}

package handlers

import (
	"context"
	"employee/connection"
	"employee/generated"
	"employee/models"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

func HandleEmployeeInactive(writer http.ResponseWriter, read *http.Request) {
	var (
		inactiverequest models.ApproveEmployee
		conn            = connection.GetDB()
		// employeedata    models.Employee
	)
	log.Println("started the inactive handler")
	err := json.NewDecoder(read.Body).Decode(&inactiverequest)
	if err != nil {
		response := models.JSONresponse{
			Status:  "error",
			Message: fmt.Sprintf("error while decode the json %v", err),
		}
		writeResponse(writer, http.StatusBadRequest, response)
	}
	querier := generated.New(conn)
	log.Println("started fetching the data from the master table")
	data, err := querier.GetEmployeeDataFrom(context.Background(), int32(inactiverequest.Id))
	if err != nil {
		resp := models.JSONresponse{
			Status:  "error",
			Message: fmt.Sprintf("error while getting the data%v", err),
			Data:    nil,
		}
		writeResponse(writer, http.StatusBadRequest, resp)
	}
	log.Println("started marshalling the data")
	makerData, err := json.Marshal(&data)
	if err != nil {
		resp := models.JSONresponse{
			Status:  "error",
			Message: fmt.Sprintf("error while marshalling the data %v", err),
			Data:    nil,
		}
		writeResponse(writer, http.StatusBadRequest, resp)
	}
	// fmt.Println(string(makerData))
	log.Println("started inserting the data to maker request")
	id, err := querier.InsertEmployee(context.Background(), generated.InsertEmployeeParams{
		EmployeeData: makerData,
		Status:       "INACTIVE",
	})
	if err != nil {
		response := models.JSONresponse{
			Status:  "error",
			Data:    id,
			Message: fmt.Sprintf("error while inserting the data %v", err),
		}
		writeResponse(writer, http.StatusBadRequest, response)
	}
	log.Println("successfully inactive the employee")
	response := models.JSONresponse{
		Status: "success",
		Data:   id,
	}
	writeResponse(writer, http.StatusOK, response)
}

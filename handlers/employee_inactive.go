package handlers

import (
	"context"
	"employee/connection"
	"employee/generated"
	"employee/models"
	"encoding/json"
	"fmt"
	"net/http"
)

func HandleEmployeeInactive(writer http.ResponseWriter, read *http.Request) {
	var (
		inactiverequest models.ApproveEmployee
		conn            = connection.GetDB()
		// employeedata    models.Employee
	)
	err := json.NewDecoder(read.Body).Decode(&inactiverequest)
	if err != nil {
		response := models.JSONresponse{
			Status:  "error",
			Message: fmt.Sprintf("error while decode the json"),
		}
		writeResponse(writer, http.StatusBadRequest, response)
	}
	querier := generated.New(conn)
	data, err := querier.GetEmployeeDataFrom(context.Background(), int32(inactiverequest.Id))
	if err != nil {
		fmt.Sprintf("error while getting the data%v", err)
	}
	makerData, err := json.Marshal(&data)
	if err != nil {
		fmt.Sprintf("error while marshalling the data")
	}
	// fmt.Println(string(makerData))
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
	response := models.JSONresponse{
		Status: "success",
		Data:   id,
	}
	writeResponse(writer, http.StatusOK, response)
}

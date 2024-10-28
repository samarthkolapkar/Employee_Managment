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
)

func HandleEmployeeApprove(writer http.ResponseWriter, read *http.Request) {
	var approveRequest models.ApproveEmployee
	var employeeinsert models.Employee
	conn := connection.GetDB()
	err := json.NewDecoder(read.Body).Decode(&approveRequest)
	if err != nil {
		resp := models.JSONresponse{
			Status:  "error",
			Message: fmt.Sprintf("error while decode the request %v", err),
		}
		writeResponse(writer, http.StatusBadRequest, resp)
	}
	querier := generated.New(conn)
	makerData, err := querier.GetEmployeefromMaker(context.Background(), generated.GetEmployeefromMakerParams{
		ID:     int32(approveRequest.Id),
		Status: "DRAFT",
	})
	if err != nil {
		fmt.Printf("error while fetching the data from the maker table %v", err)
	}
	for _, data := range makerData {
		// err := json.Unmarshal(data, &employeeinsert)
		err := json.Unmarshal(data.EmployeeData, &employeeinsert)
		if err != nil {
			resp := models.JSONresponse{
				Status:  "error",
				Message: fmt.Sprintf("error while unmarshalling the data"),
			}
			writeResponse(writer, http.StatusBadRequest, resp)
		}
	}
	id, err := querier.InsertIntoEmployee(context.Background(), generated.InsertIntoEmployeeParams{
		FirstName: employeeinsert.FirstName,
		MiddleName: sql.NullString{
			String: employeeinsert.MiddleName,
			Valid:  true,
		},
		LastName:      employeeinsert.LastName,
		Age:           int32(employeeinsert.Age),
		PanNumber:     employeeinsert.PanNumber,
		AddressStreet: employeeinsert.Address.Street,
		AddressCity:   employeeinsert.Address.City,
		AddressState:  employeeinsert.Address.State,
		AddressZip: sql.NullInt64{
			Int64: int64(employeeinsert.Address.Zip),
			Valid: true,
		},
	})
	if err != nil {
		fmt.Printf("error while inserting the data into master table")
	}
	resp := models.JSONresponse{
		Status:  "success",
		Data:    id,
		Message: nil,
	}
	writeResponse(writer, http.StatusOK, resp)
}

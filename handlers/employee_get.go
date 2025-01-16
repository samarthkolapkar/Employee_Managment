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

func HandleEmployeeForward(writer http.ResponseWriter, read *http.Request) {
	var forwardRequest models.EmployeeGet
	conn := connection.GetDB()

	log.Println("get request started")
	err := json.NewDecoder(read.Body).Decode(&forwardRequest)
	if err != nil {
		resp := models.JSONresponse{
			Status:  "error",
			Message: fmt.Errorf("error while binding the json request %v", err),
		}
		writeResponse(writer, http.StatusBadRequest, resp)
	}
	querier := generated.New(conn)

	data, err := querier.GetEmployeefromMaker(context.Background(), generated.GetEmployeefromMakerParams{
		ID:     int32(forwardRequest.Id),
		Status: forwardRequest.Status,
	})
	if err != nil {
		resp := models.JSONresponse{
			Status:  "error",
			Message: fmt.Sprintf("error while getting the record %v", err),
			Data:    nil,
		}
		writeResponse(writer, http.StatusBadRequest, resp)
	}

	resp := models.JSONresponse{
		Status:  "success",
		Data:    data,
		Message: nil,
	}
	log.Println("get request completed")
	writeResponse(writer, http.StatusBadRequest, resp)

}

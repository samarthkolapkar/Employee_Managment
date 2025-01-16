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
	"slices"
)

func HandleEmployeeList(writer http.ResponseWriter, read *http.Request) {
	log.Println("started the employee list handler")
	var listrequest *models.EmployeeList
	db := connection.GetDB()
	log.Println("started unmarshalling the request")
	err := json.NewDecoder(read.Body).Decode(&listrequest)
	if err != nil {
		resp := models.JSONresponse{
			Status:  "error",
			Message: fmt.Sprintf("Failed to parse the JSON %v", err),
		}
		writeResponse(writer, http.StatusBadRequest, resp)
	}
	queries := generated.New(db)
	if slices.Contains(statusArray, listrequest.Status) {
		log.Println("getting approved list of data from the database")
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
	} else {
		log.Println("getting pending list of data from the database")
		makerlist, err := queries.GetlistFromEmployeeMaker(context.Background(), generated.GetlistFromEmployeeMakerParams{
			Status:  listrequest.Status,
			Column2: listrequest.Search,
		})
		if err != nil {
			resp := models.JSONresponse{
				Status:  "error",
				Message: fmt.Sprintf("error while getting the data from the database!! %v", err),
				Data:    nil,
			}
			writeResponse(writer, http.StatusOK, resp)
		}
		resp := models.JSONresponse{
			Status: "success",
			Data:   makerlist,
		}
		log.Println("successfully fetched the data from the database")
		writeResponse(writer, http.StatusOK, resp)
	}

}

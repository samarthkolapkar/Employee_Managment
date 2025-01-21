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
	var listResponse models.EmployeeResponse
	// var count int64
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

		resp := models.JSONListresponse{
			Status:  "success",
			Count:   len(list),
			Message: nil,
			Data:    list,
		}
		writeListResponse(writer, http.StatusOK, resp)
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
			writeResponse(writer, http.StatusInternalServerError, resp)
		}
		var response []models.EmployeeResponse
		var empData models.Employee
		for i, data := range makerlist {
			err := json.Unmarshal(data.EmployeeData, &empData)
			if err != nil {
				resp := models.JSONresponse{
					Status:  "error",
					Message: fmt.Sprintf("error while unmarshalling the data %v", err),
					Data:    nil,
				}
				writeResponse(writer, http.StatusInternalServerError, resp)
			}
			listResponse.ID = int64(makerlist[i].ID)
			listResponse.Status = makerlist[i].Status
			listResponse.FirstName = empData.FirstName
			listResponse.LastName = empData.LastName
			listResponse.MiddleName = empData.MiddleName
			listResponse.PanNumber = empData.PanNumber
			listResponse.Contact = empData.Contact
			listResponse.IsExperienced = empData.IsExperienced
			listResponse.ExperiencedDetails = empData.ExperiencedDetails
			listResponse.Address = empData.Address
			listResponse.Age = empData.Age
			response = append(response, listResponse)
		}
		statusCount, err := queries.StatusCount(context.Background())
		if err != nil {
			resp := models.JSONresponse{
				Status:  "error",
				Message: fmt.Sprintf("error while getting the data from the database!! %v", err),
				Data:    nil,
			}
			writeResponse(writer, http.StatusInternalServerError, resp)
		}
		resp := models.JSONListresponse{
			Status:      "success",
			Count:       len(makerlist),
			Data:        response,
			StatusCount: statusCount,
		}
		log.Println("successfully fetched the data from the database")
		writeListResponse(writer, http.StatusOK, resp)
	}

}

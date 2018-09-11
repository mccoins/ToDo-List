package controllers

import (
	"encoding/json"
	"io/ioutil"
	"net/http"

	"github.com/google/go-cmp/cmp"
	"github.com/gorilla/mux"
	data "github.com/mccoins/ToDo-List/data"
	m "github.com/mccoins/ToDo-List/models"
	u "github.com/mccoins/ToDo-List/utilities"
)

// GetLists returns all the current lists
func GetLists(w http.ResponseWriter, r *http.Request) {
	if Authenticated(r) {

		// create  DTO
		dto := m.ToDoListsDTO{}
		u.InitializeDTO(&dto)

		// get database instance
		datasource := data.GetDataSourceInstance()

		// check for existing filter criteria
		vars := mux.Vars(r)

		if len(vars) >= 1 {
			filterString := vars["filter"]

			if filterString != "" {
				filteredList := u.FilterLists(datasource.Database, vars["filter"])

				// if the filtered list is empty inform the client
				// an alternative to using the go-cmp package, reflect.DeepEqual will do
				// the same thing but less effectively due to the use of reflection
				if cmp.Equal(*filteredList, []m.TodoList{}) {
					dto.Response = m.ResponseCode{ID: "404", Description: "Item not found"}
					dto.Status = "Item not found"
				}

				dto.Lists = *filteredList
				json.NewEncoder(w).Encode(&dto)
			} else {
				json.NewEncoder(w).Encode(m.ResponseCode{ID: "400", Description: "Bad Request"})
			}
		} else {
			// return the complete list
			dto.Lists = datasource.Database
			json.NewEncoder(w).Encode(&dto)
		}
	} else {
		json.NewEncoder(w).Encode(m.ResponseCode{ID: "401", Description: "Unauthorized"})
	}
}

// CreateList creates a new list
func CreateList(w http.ResponseWriter, r *http.Request) {
	if Authenticated(r) {
		// create  DTO
		dto := m.ToDoListDTO{}
		u.InitializeDTO(&dto)

		// read in the request
		defer r.Body.Close()
		body, _ := ioutil.ReadAll(r.Body)

		// unmarshal into dto object for processing
		err := json.Unmarshal(body, &dto)
		if err != nil {
			dto.Response = m.ResponseCode{ID: "400", Description: "Invalid Request"}
			dto.Status = "Bad Request"
		} else {
			repo := data.ListRepository{}
			err := repo.Add(&dto.List)

			if err != nil {
				dto.Response = m.ResponseCode{ID: "409", Description: "Item already exists"}
				dto.Status = "Conflict"
			} else {
				dto.Response = m.ResponseCode{ID: "201", Description: "Item created"}
				dto.Status = "Success"
			}
		}

		// encode the json response to send back to the client
		json.NewEncoder(w).Encode(dto)
	} else {
		json.NewEncoder(w).Encode(m.ResponseCode{ID: "401", Description: "Unauthorized"})
	}
}

// GetListByID returns a list by ID
func GetListByID(w http.ResponseWriter, r *http.Request) {
	if Authenticated(r) {
		// Create DTO
		dto := m.ToDoListDTO{}
		u.InitializeDTO(&dto)

		vars := mux.Vars(r)

		// check for passed in ID
		if len(vars) < 1 {
			dto.Response = m.ResponseCode{ID: "400", Description: "Invalid Request"}
			dto.Status = "Bad Request"
		} else {
			// get list by id
			dto.List = *data.ListRepository{}.FindByID(vars["id"])

			// if no results are found, return a 404
			if cmp.Equal(dto.List, m.TodoList{}) {
				dto.Response = m.ResponseCode{ID: "404", Description: "Item not found"}
				dto.Status = "Item not found"
			}
		}

		// encode the json response to send back to the client
		json.NewEncoder(w).Encode(dto)
	} else {
		json.NewEncoder(w).Encode(m.ResponseCode{ID: "401", Description: "Unauthorized"})
	}

}

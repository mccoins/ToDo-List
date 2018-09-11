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

// AddTask adds a new task to an existing list
func AddTask(w http.ResponseWriter, r *http.Request) {
	if Authenticated(r) {
		// create  DTO
		dto := m.TaskDTO{}
		u.InitializeDTO(&dto)

		vars := mux.Vars(r)

		// check for passed in ID
		if len(vars) < 1 {
			dto.Response = m.ResponseCode{ID: "400", Description: "Invalid Request"}
			dto.Status = "Bad Request"
			json.NewEncoder(w).Encode(dto)
		} else {
			// find the list by id
			list := *data.ListRepository{}.FindByID(vars["id"])

			// if list not found return a 404
			if cmp.Equal(list, m.TodoList{}) {
				json.NewEncoder(w).Encode(m.ResponseCode{ID: "404", Description: "Item not found"})
			} else {
				// read in the request
				defer r.Body.Close()
				body, _ := ioutil.ReadAll(r.Body)

				// unmarshal into dto object for processing
				err := json.Unmarshal(body, &dto)
				if err != nil {
					dto.Response = m.ResponseCode{ID: "400", Description: "Invalid Request"}
					dto.Status = "Bad Request"
				} else {
					repo := data.TaskRepository{}
					resp := repo.Add(list.ID, &dto.Task)

					if !cmp.Equal(resp, m.TodoList{}) {
						// set response to 201
						dto.Response.ID = "201"
						dto.Response.Description = "Item created"
					}
				}
				json.NewEncoder(w).Encode(dto)
			}
		}
	} else {
		json.NewEncoder(w).Encode(m.ResponseCode{ID: "401", Description: "Unauthorized"})
	}
}

// SetToComplete sets a task lists status to complete
func SetToComplete(w http.ResponseWriter, r *http.Request) {
	if Authenticated(r) {

		// initialize a list DTO to return back to the client
		dto := &m.ToDoListDTO{}
		u.InitializeDTO(dto)

		vars := mux.Vars(r)

		// check for parameters
		if len(vars) < 1 {
			dto.Response = m.ResponseCode{ID: "400", Description: "Invalid Request"}
			dto.Status = "Bad Request"
		} else {
			// update task to a completed status and return the list
			dto.List = *data.TaskRepository{}.Update(vars["id"], vars["taskId"])

			// if the list or the task were not found send 404 otherwise 201
			if cmp.Equal(dto.List, m.TodoList{}) {
				dto.Response = m.ResponseCode{ID: "404", Description: "Item not found"}
				dto.Status = "Item not found"
			} else {
				dto.Response = m.ResponseCode{ID: "201", Description: "Item updated"}
				dto.Status = "Success"
			}
		}

		json.NewEncoder(w).Encode(dto)
	} else {
		json.NewEncoder(w).Encode(m.ResponseCode{ID: "401", Description: "Unauthorized"})
	}
}

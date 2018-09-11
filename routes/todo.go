package routes

import (
	"net/http"

	log "github.com/Sirupsen/logrus"
	"github.com/gorilla/mux"
	c "github.com/mccoins/ToDo-List/controllers"
)

// LoadToDoListRoutes loads the routes for the todolist api
func LoadToDoListRoutes(r *mux.Router) {
	r.HandleFunc("/lists", c.GetLists).Methods("GET")
	r.HandleFunc("/lists/{filter}", c.GetLists).Methods("GET")
	r.HandleFunc("/list", c.CreateList).Methods("POST")
	r.HandleFunc("/list/{id}", c.GetListByID).Methods("GET")
	r.HandleFunc("/list/{id}/task", c.AddTask).Methods("POST")
	r.HandleFunc("/list/{id}/task/{taskId}/complete", c.SetToComplete).Methods("POST")

	r.HandleFunc("/auth/{username}/{password}", c.GetToken).Methods("GET")
	log.Fatal(http.ListenAndServe(":8000", r))
}

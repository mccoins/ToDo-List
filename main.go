package main

import (
	"github.com/gorilla/mux"
	route "github.com/mccoins/ToDo-List/routes"
)

func main() {
	// initialize the router and routes
	router := mux.NewRouter()
	route.LoadToDoListRoutes(router)
}

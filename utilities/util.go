package utilities

import (
	"strings"

	d "github.com/mccoins/ToDo-List/data"
	i "github.com/mccoins/ToDo-List/interfaces"
	m "github.com/mccoins/ToDo-List/models"
)

// InitializeDTO accpets an IDTO interface and initializes the object
func InitializeDTO(dto i.IDTO) {
	dto.Initialize()
}

// FilterLists is used to filter a list of todo lists based ONLY on filtering the Name of the list
func FilterLists(list []m.TodoList, filter string) *[]m.TodoList {
	filteredList := Filter(list, func(v m.TodoList) bool {
		return strings.Contains(v.Name, filter)
	})
	return &filteredList
}

// Filter aids in the filtering process
func Filter(lists []m.TodoList, f func(m.TodoList) bool) []m.TodoList {
	filteredList := make([]m.TodoList, 0)
	for _, list := range lists {
		if f(list) {
			filteredList = append(filteredList, list)
		}
	}
	return filteredList
}

// AddList is a helper function for unit tests
func AddList(listID string) {
	input := m.TodoList{ID: listID}
	listRepo := d.ListRepository{}
	listRepo.Add(&input)
}

// AddFullList is a helper function for unit tests
func AddFullList(list *m.TodoList) {
	listRepo := d.ListRepository{}
	listRepo.Add(list)
}

// AddTask is a helper function for unit tests
func AddTask(listID string, taskID string) {
	// add task to list
	task := m.Task{ID: taskID, Completed: false}
	taskRepo := d.TaskRepository{}
	taskRepo.Add(listID, &task)
}

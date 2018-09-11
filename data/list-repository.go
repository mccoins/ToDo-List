package data

import (
	"errors"

	"github.com/google/go-cmp/cmp"
	m "github.com/mccoins/ToDo-List/models"
)

// ListRepository represents the repo adapter for lists
type ListRepository struct {
}

// Add function adds a todo list to the datasource
func (r *ListRepository) Add(list *m.TodoList) error {
	datasource := GetDataSourceInstance()

	// check if list already exissts (by list id)
	lst := ListRepository{}.FindByID(list.ID)

	// if the list does not already exist, add it
	if !cmp.Equal(list, lst) {
		datasource.Database = append(datasource.Database, *list)
	} else {
		return errors.New("already exists")
	}

	return nil
}

// FindByID returns a specified list by its corrisponding id
func (r ListRepository) FindByID(id string) *m.TodoList {
	datasource := GetDataSourceInstance()
	list := m.TodoList{}

	// simple linear search
	for _, lst := range datasource.Database {
		if lst.ID == id {
			list = lst
			break
		}
	}
	return &list
}

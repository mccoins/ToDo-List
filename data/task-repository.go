package data

import (
	m "github.com/mccoins/ToDo-List/models"
)

// TaskRepository represents the repo adapter for tasks
type TaskRepository struct {
}

// Add function adds a task to the datasource based on list id
func (r *TaskRepository) Add(listID string, task *m.Task) *m.TodoList {
	datasource := GetDataSourceInstance()
	list := m.TodoList{}

	// check existing lists
	for idx, lst := range datasource.Database {
		if lst.ID == listID {
			// list found, now add task
			lst.Tasks = append(lst.Tasks, *task)

			// update the list
			datasource.Database[idx].Tasks = lst.Tasks
			list = lst
			break
		}
	}

	return &list
}

// Update function updates an existing task (in this case it just sets completed = true)
func (r TaskRepository) Update(listID string, taskID string) *m.TodoList {
	// find the list based on ID
	list := *ListRepository{}.FindByID(listID)
	updated := false

	// find the specified task within the list
	for idx, tsk := range list.Tasks {
		if tsk.ID == taskID {
			list.Tasks[idx].Completed = true
			updated = true
		}
	}

	// if task was not updated, return an empty list
	if !updated {
		list = m.TodoList{}
	}

	return &list
}

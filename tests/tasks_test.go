package tests_test

import (
	"testing"

	d "github.com/mccoins/ToDo-List/data"
	u "github.com/mccoins/ToDo-List/utilities"
)

// TestAddTask tests adding a task to an existing list
func TestAddTask(t *testing.T) {

	// clear database
	d.GetDataSourceInstance().Clear()

	// add list to db
	u.AddList("1")

	// add task to list
	u.AddTask("1", "1")

	// test validity
	if len(d.GetDataSourceInstance().Database[0].Tasks) != 1 {
		t.Error("Unable to add task", len(d.GetDataSourceInstance().Database))
	}
}

// TestUpdateTaskToComplete tests updating a task and setting completed = true
func TestUpdateTaskToComplete(t *testing.T) {

	// clear database
	d.GetDataSourceInstance().Clear()

	u.AddList("1")
	u.AddTask("1", "1")

	list := *d.TaskRepository{}.Update("1", "1")

	if list.Tasks[0].Completed != true {
		t.Error("unable to update task to completed")
	}
}

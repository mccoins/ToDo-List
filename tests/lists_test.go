package tests_test

import (
	"testing"

	"github.com/google/go-cmp/cmp"

	d "github.com/mccoins/ToDo-List/data"
	m "github.com/mccoins/ToDo-List/models"
	u "github.com/mccoins/ToDo-List/utilities"
)

// TestAddList tests adding a list to the datasource
func TestAddList(t *testing.T) {

	// clear database
	d.GetDataSourceInstance().Clear()

	u.AddList("1")

	if len(d.GetDataSourceInstance().Database) != 1 {
		t.Error("Unable to add List")
	}
}

// TestAddDuplicateList tests whether duplicate lists can be created
func TestAddDuplicateList(t *testing.T) {

	// clear database
	d.GetDataSourceInstance().Clear()

	u.AddList("1")
	u.AddList("1")

	if len(d.GetDataSourceInstance().Database) != 1 {
		t.Error("Duplicate list added to database")
	}
}

// TestFindListByID tests finding a list by a specific id
func TestFindListByID(t *testing.T) {

	// clear database
	d.GetDataSourceInstance().Clear()

	// add 1 entry with an id of 1
	u.AddList("1")

	input := m.TodoList{ID: "1"}

	// check validity
	valid := d.ListRepository{}.FindByID("1")

	if !cmp.Equal(&input, valid) {
		t.Error("unable to find list by id", valid)
	}

	// check invalidity
	invalid := d.ListRepository{}.FindByID("2")
	if cmp.Equal(&input, invalid) {
		t.Error("unable to find list by id", invalid)
	}
}

// TestFilteredList is used to test the ability to get a filtered result set
func TestFilteredList(t *testing.T) {

	// clear database
	d.GetDataSourceInstance().Clear()

	// create 3 entries, 2 with "li" in the name
	u.AddFullList(&m.TodoList{
		ID:   "1",
		Name: "list 1",
	})
	u.AddFullList(&m.TodoList{
		ID:   "2",
		Name: "Temporary",
	})
	u.AddFullList(&m.TodoList{
		ID:   "3",
		Name: "limabean",
	})

	filterString := "li"
	filteredList := u.FilterLists(d.GetDataSourceInstance().Database, filterString)

	if len(*filteredList) != 2 {
		t.Error("Unable to filter list")
	}
}

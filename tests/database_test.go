package tests_test

import (
	"testing"

	d "github.com/mccoins/ToDo-List/data"
	u "github.com/mccoins/ToDo-List/utilities"
)

// TestClearDatabase tests clearing the database
func TestClearDatabase(t *testing.T) {
	// ensure the database is empty
	d.GetDataSourceInstance().Clear()

	// add 1 entry
	u.AddList("1")

	// clear the database
	d.GetDataSourceInstance().Clear()

	if len(d.GetDataSourceInstance().Database) != 0 {
		t.Error("Unable to clear database")
	}
}

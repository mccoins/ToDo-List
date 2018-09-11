package data

import (
	"sync"

	models "github.com/mccoins/ToDo-List/models"
)

// InMemoryDataSource represents an in memory version of the data in session
// using a singleton ensures only one instance will get created
type InMemoryDataSource struct {
	Database []models.TodoList
}

// init initializes the database
func (ds InMemoryDataSource) init() {
	ds.Database = []models.TodoList{}
}

var dbInstance *InMemoryDataSource
var once sync.Once

// GetDataSourceInstance is used to return the singleton InMemoryDataSource
func GetDataSourceInstance() *InMemoryDataSource {
	once.Do(func() {
		dbInstance = &InMemoryDataSource{}
		dbInstance.init()
	})
	return dbInstance
}

// Clear reinitializes the database, primarily for unit testing only
func (ds *InMemoryDataSource) Clear() {
	ds.Database = []models.TodoList{}
}

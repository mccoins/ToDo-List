package interfaces

// IRepository is used to define a repository for use with multiple datastores
type IRepository interface {
	Add(interface{}) error
}

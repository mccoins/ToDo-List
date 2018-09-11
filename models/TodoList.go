package models

// TodoList describes a specific to do list which includes id, name, description, and a list of related tasks
type TodoList struct {
	ID          string `json:"Id,omitempty"`
	Name        string `json:"Name,omitempty"`
	Description string `json:"Description,omitempty"`
	Tasks       []Task `json:"Tasks,omitempty"`
}

package models

// Task used to identify a specific task
type Task struct {
	ID        string `json:"Id,omitempty"`
	Name      string `json:"Name,omitempty"`
	Completed bool   `json:"Completed"`
}

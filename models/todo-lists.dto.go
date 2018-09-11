package models

// ToDoListsDTO is the data transfer object used to encapsulate the request and response from the client
type ToDoListsDTO struct {
	Lists    []TodoList   `json:"Lists,omitempty"`
	Response ResponseCode `json:"ResponseCode,omitempty"`
	Status   string       `json:"Status,omitempty"`
}

// Initialize is used to set default values
func (dto *ToDoListsDTO) Initialize() {
	dto.Lists = []TodoList{}
	dto.Response = ResponseCode{
		ID:          "200",
		Description: "OK",
	}
	dto.Status = "Success"
}

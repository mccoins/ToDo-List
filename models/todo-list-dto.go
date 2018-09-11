package models

// ToDoListDTO is the data transfer object used to encapsulate the request and response from the client
type ToDoListDTO struct {
	List     TodoList     `json:"List,omitempty"`
	Response ResponseCode `json:"ResponseCode,omitempty"`
	Status   string       `json:"Status,omitempty"`
}

// Initialize is used to set default values
func (dto *ToDoListDTO) Initialize() {
	dto.List = TodoList{}
	dto.Response = ResponseCode{
		ID:          "200",
		Description: "OK",
	}
	dto.Status = "Success"
}

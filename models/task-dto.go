package models

// TaskDTO is the data transfer object used to encapsulate the request and response from the client
type TaskDTO struct {
	Task     Task         `json:"Task,omitempty"`
	Response ResponseCode `json:"ResponseCode,omitempty"`
	Status   string       `json:"Status,omitempty"`
}

// Initialize is used to set default values
func (dto *TaskDTO) Initialize() {
	dto.Task = Task{}
	dto.Response = ResponseCode{
		ID:          "200",
		Description: "OK",
	}
	dto.Status = "Success"
}

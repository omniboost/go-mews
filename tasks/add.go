package tasks

import (
	"time"

	"github.com/omniboost/go-mews/json"
)

const (
	endpointAdd = "tasks/add"
)

// List all products
func (s *Service) Add(requestBody *AddRequest) (*AddResponse, error) {
	// @TODO: create wrapper?
	if err := s.Client.CheckTokens(); err != nil {
		return nil, err
	}

	apiURL, err := s.Client.GetApiURL(endpointAdd)
	if err != nil {
		return nil, err
	}

	responseBody := &AddResponse{}
	httpReq, err := s.Client.NewRequest(apiURL, requestBody)
	if err != nil {
		return nil, err
	}

	_, err = s.Client.Do(httpReq, responseBody)
	return responseBody, err
}

func (s *Service) NewAddRequest() *AddRequest {
	return &AddRequest{}
}

type AddRequest struct {
	json.BaseRequest
	Name           string    `json:"Name"`                     // Name (or title) of the task.
	Description    string    `json:"Description,omitempty"`    // Further decription of the task.
	DeadlineUTC    time.Time `json:"DeadlineUtc"`              // Deadline of the task in UTC timezone in ISO 8601 format.
	ServiceOrderID string    `json:"ServiceOrderId,omitempty"` // Unique identifier of the service order (reservation or product service order) the task is linked with.
	DepartmentID   string    `json:"DepartmentId,omitempty"`   // Unique identifier of the Department the task is addressed to.
}

type AddResponse struct {
}

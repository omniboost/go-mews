package tasks

import (
	"time"

	"github.com/omniboost/go-mews/configuration"
	base "github.com/omniboost/go-mews/json"
	"github.com/omniboost/go-mews/omitempty"
)

const (
	endpointAll = "tasks/getAll"
)

// List all tasks
func (s *Service) All(requestBody *AllRequest) (*AllResponse, error) {
	// @TODO: create wrapper?
	if err := s.Client.CheckTokens(); err != nil {
		return nil, err
	}

	apiURL, err := s.Client.GetApiURL(endpointAll)
	if err != nil {
		return nil, err
	}

	responseBody := &AllResponse{}
	httpReq, err := s.Client.NewRequest(apiURL, requestBody)
	if err != nil {
		return nil, err
	}

	_, err = s.Client.Do(httpReq, responseBody)
	return responseBody, err
}

type AllResponse struct {
	Tasks Tasks `json:"Tasks"` // The filtered tasks.
}

func (s *Service) NewAllRequest() *AllRequest {
	return &AllRequest{}
}

type AllRequest struct {
	base.BaseRequest

	TaskIds         []string                   `json:"TaskIds,omitempty"`
	DepartmentIds   []string                   `json:"DepartmentIds,omitempty"`
	ServiceOrderIds []string                   `json:"ServiceOrderIds,omitempty"`
	CreatedUtc      configuration.TimeInterval `json:"CreatedUtc,omitempty"`
	ClosedUtc       configuration.TimeInterval `json:"ClosedUtc,omitempty"`
	DeadlineUtc     configuration.TimeInterval `json:"DeadlineUtc,omitempty"`
}

func (r AllRequest) MarshalJSON() ([]byte, error) {
	return omitempty.MarshalJSON(r)
}

type Tasks []Task

type Task struct {
	ID             string    `json:"Id"`             // Unique identifier of the task.
	Name           string    `json:"Name"`           // Name (or title) of the task.
	State          TaskState `json:"State"`          // State of the task.
	Description    string    `json:"Description"`    // Further decription of the task.
	DepartmentID   string    `json:"DepartmentId"`   // Unique identifier of the Department the task is addressed to.
	ServiceOrderID string    `json:"ServiceOrderId"` // Unique identifier of the service order (reservation or product service order) the task is linked with.
	CreatedUTC     time.Time `json:"CreatedUtc"`     // Creation date and time of the task in UTC timezone in ISO 8601 format.
	DeadlineUTC    time.Time `json:"DeadlineUtc"`    // Deadline date and time of the task in UTC timezone in ISO 8601 format.
	ClosedUTC      time.Time `json:"ClosedUtc"`      // Last update date and time of the task in UTC timezone in ISO 8601 format.
}

type TaskState string

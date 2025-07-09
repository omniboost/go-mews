package fiscalmachinecommands

import (
	"time"

	"github.com/omniboost/go-mews/bills"
	"github.com/omniboost/go-mews/commands"
	"github.com/omniboost/go-mews/configuration"
	"github.com/omniboost/go-mews/json"
	"github.com/omniboost/go-mews/omitempty"
)

const (
	endpointAll = "fiscalMachineCommands/getAll"
)

// List all commands
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

func (s *Service) NewAllRequest() *AllRequest {
	return &AllRequest{
		DeviceIDs: []string{},
	}
}

type AllRequest struct {
	json.BaseRequest
	DeviceIDs  []string                   `json:"DeviceIds"`
	States     []commands.CommandState             `json:"States,omitempty"`
	UpdatedUTC configuration.TimeInterval `json:"UpdatedUtc,omitempty"`
	Limitation json.Limitation            `json:"Limitation,omitempty"`
}

func (r AllRequest) MarshalJSON() ([]byte, error) {
	return omitempty.MarshalJSON(r)
}

type AllResponse struct {
	Commands Commands `json:"Commands"`
	Cursor   string   `json:"Cursor"`
}

type Commands []Command

type Command struct {
	ID         string                `json:"Id"`             // Unique identifier of the command.
	CreatedUTC time.Time             `json:"CreatedUtc"`     // Date and time of the command was created in UTC timezone in ISO 8601 format.
	Bill       bills.Bill            `json:"Bill,omitempty"` // If available add Bill informaion
	Device     commands.Device       `json:"Device"`         // Device information
	State      commands.CommandState `json:"State"`          // State of the command.
}

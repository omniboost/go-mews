package cashiers

import (
	"time"

	"github.com/omniboost/go-mews/json"
	"github.com/omniboost/go-mews/omitempty"
)

const (
	endpointAll = "cashiers/getAll"
)

// List all cashiers
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
	return &AllRequest{}
}

type AllRequest struct {
	json.BaseRequest
	Limitation json.Limitation `json:"Limitation,omitempty"`
}

func (r AllRequest) MarshalJSON() ([]byte, error) {
	return omitempty.MarshalJSON(r)
}

type AllResponse struct {
	Cashiers Cashiers `json:"Cashiers"`
	Cursor   string   `json:"Cursor"`
}

type Cashiers []Cashier

type Cashier struct {
	ID         string    `json:"Id"`         // Unique identifier of the transaction.
	IsActive   bool      `json:"IsActive"`   // Whether the cashier is still active.
	Name       string    `json:"Name"`       // Name of the cashier.
	CreatedUTC time.Time `json:"CreatedUtc"` // Creation date and time of the Cashier in UTC timezone in ISO 8601 format.
	UpdatedUTC time.Time `json:"UpdatedUtc"` // Last update date and time of the Cashier in UTC timezone in ISO 8601 format.
}

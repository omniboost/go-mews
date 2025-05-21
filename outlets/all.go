package outlets

import (
	base "github.com/omniboost/go-mews/json"
)

const (
	endpointAll = "outlets/getAll"
)

// List all outlets
func (s *APIService) All(requestBody *AllRequest) (*AllResponse, error) {
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
	Outlets Outlets
	Cursor  string `json:"Cursor"`
}

func (s *APIService) NewAllRequest() *AllRequest {
	return &AllRequest{}
}

type AllRequest struct {
	base.BaseRequest
	Limitation base.Limitation `json:"Limitation,omitempty"`
}

type Outlets []Outlet

type Outlet struct {
	ID       string `json:"Id"`       // Unique identifier of the outlet
	IsActive bool   `json:"IsActive"` // Whether the outlet is still active.
	Name     string `json:"Name"`     // Name of the outlet.
}

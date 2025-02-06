package commands

import (
	"github.com/omniboost/go-mews/json"
)

const (
	endpointAllByIDs = "commands/getAllByIDs"
)

// List all commands
func (s *Service) AllByIDs(requestBody *AllByIDsRequest) (*AllByIDsResponse, error) {
	// @TODO: create wrapper?
	if err := s.Client.CheckTokens(); err != nil {
		return nil, err
	}

	apiURL, err := s.Client.GetApiURL(endpointAllByIDs)
	if err != nil {
		return nil, err
	}

	responseBody := &AllByIDsResponse{}
	httpReq, err := s.Client.NewRequest(apiURL, requestBody)
	if err != nil {
		return nil, err
	}

	_, err = s.Client.Do(httpReq, responseBody)
	return responseBody, err
}

func (s *Service) NewAllByIDsRequest() *AllByIDsRequest {
	return &AllByIDsRequest{}
}

type AllByIDsRequest struct {
	json.BaseRequest
	// Unique identifiers of Commands to be returned.
	CommandIDs []string `json:"CommandIds"`
}

type AllByIDsResponse struct {
	Commands Commands `json:"Commands"` // The closed bills.
}

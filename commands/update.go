package commands

import "github.com/omniboost/go-mews/json"

const (
	endpointUpdate = "commands/update"
)

// List all products
func (s *Service) Update(requestBody *UpdateRequest) (*UpdateResponse, error) {
	// @TODO: create wrapper?
	if err := s.Client.CheckTokens(); err != nil {
		return nil, err
	}

	apiURL, err := s.Client.GetApiURL(endpointUpdate)
	if err != nil {
		return nil, err
	}

	responseBody := &UpdateResponse{}
	httpReq, err := s.Client.NewRequest(apiURL, requestBody)
	if err != nil {
		return nil, err
	}

	_, err = s.Client.Do(httpReq, responseBody)
	return responseBody, err
}

func (s *Service) NewUpdateRequest() *UpdateRequest {
	return &UpdateRequest{}
}

type UpdateRequest struct {
	json.BaseRequest
	CommandID string       `json:"CommandId"`
	State     CommandState `json:"State"`
	Notes     string       `json:"Notes"`
}

type UpdateResponse struct {
}

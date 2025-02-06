package creditcards

import "github.com/omniboost/go-mews/json"

const (
	endpointAllByIDs = "creditCards/getAllByIds"
)

// List all products
func (s *Service) AllByIDs(requestBody *AllByIDsRequest) (*AllResponse, error) {
	// @TODO: create wrapper?
	if err := s.Client.CheckTokens(); err != nil {
		return nil, err
	}

	apiURL, err := s.Client.GetApiURL(endpointAllByIDs)
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

func (s *Service) NewAllByIDsRequest() *AllByIDsRequest {
	return &AllByIDsRequest{}
}

type AllByIDsRequest struct {
	json.BaseRequest
	CreditCardIDs []string `json:"CreditCardIds"`
}

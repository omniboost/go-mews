package creditcards

import (
	"github.com/omniboost/go-mews/configuration"
	"github.com/omniboost/go-mews/json"
	"github.com/omniboost/go-mews/omitempty"
)

const (
	endpointAll = "creditCards/getAll"
)

// List all products
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
	EnterpriseIDs []string                   `json:"EnterpriseIds"`        // Unique identifiers of the Enterprises. If not specified, the operation returns data for all enterprises within scope of the Access Token.
	CreditCardIDs []string                   `json:"CreditCardIds"`        // Unique identifiers of the Credit cards. Required if no other filter is provided.
	CustomerIDs   []string                   `json:"CustomerIds"`          // Unique identifiers of the Customers.
	UpdatedUTC    configuration.TimeInterval `json:"UpdatedUtc,omitempty"` // Interval in which Credit card was updated.
	Limitation    json.Limitation            `json:"Limitation,omitempty"`
}

func (r AllRequest) MarshalJSON() ([]byte, error) {
	return omitempty.MarshalJSON(r)
}

type AllResponse struct {
	CreditCards CreditCards `json:"CreditCards"` // The credit cards.
	Cursor      string      `json:"Cursor"`      // Unique identifier of the item one newer in time order than the items to be returned. If Cursor is not specified, i.e. null, then the latest or most recent items will be returned.
}

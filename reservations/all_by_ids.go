package reservations

import "github.com/omniboost/go-mews/json"

const (
	endpointAllByIDs = "reservations/getAllByIds"
)

// List all products
func (s *APIService) AllByIDs(requestBody *AllByIDsRequest) (*AllResponse, error) {
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

func (s *APIService) NewAllByIDsRequest() *AllByIDsRequest {
	return &AllByIDsRequest{}
}

type AllByIDsRequest struct {
	json.BaseRequest
	ReservationIDs []string          `json:"ReservationIds"`
	Extent         ReservationExtent `json:"Extent,omitempty"`
}

package enterprises

import (
	"github.com/omniboost/go-mews/json"
	"github.com/omniboost/go-mews/omitempty"
)

const (
	endpointCountersGetAll = "counters/getAll"
)

// List all products
func (s *APIService) CountersGetAll(requestBody *CountersGetAllRequest) (*CountersGetAllResponse, error) {
	// @TODO: create wrapper?
	if err := s.Client.CheckTokens(); err != nil {
		return nil, err
	}

	apiURL, err := s.Client.GetApiURL(endpointCountersGetAll)
	if err != nil {
		return nil, err
	}

	responseBody := &CountersGetAllResponse{}
	httpReq, err := s.Client.NewRequest(apiURL, requestBody)
	if err != nil {
		return nil, err
	}

	_, err = s.Client.Do(httpReq, responseBody)
	return responseBody, err
}

func (s *APIService) NewCountersGetAllRequest() *CountersGetAllRequest {
	return &CountersGetAllRequest{}
}

type CountersGetAllRequest struct {
	json.BaseRequest
}

func (r CountersGetAllRequest) MarshalJSON() ([]byte, error) {
	return omitempty.MarshalJSON(r)
}

type CountersGetAllResponse struct {
	BillCounters             Counters `json:"BillCounters"`
	ProformaCounters         Counters `json:"ProformaCounters"`
	ServiceOrderCounters     Counters `json:"ServiceOrderCounters"`
	RegistrationCardCounters Counters `json:"RegistrationCardCounters"`
}

type Counters []Counter

type Counter struct {
	ID        string `json:"Id"`
	Name      string `json:"Name"`
	IsDefault bool   `json:"IsDefault"`
	Value     int    `json:"Value"`
	Format    string `json:"Format"`
}

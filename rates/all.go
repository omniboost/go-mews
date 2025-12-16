package rates

import (
	"github.com/omniboost/go-mews/configuration"
	"github.com/omniboost/go-mews/json"
)

const (
	endpointAll = "rates/getAll"
)

// List all products
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
	Rates      Rates      `json:"Rates"`      // Rates of the default service.
	RateGroups RateGroups `json:"RateGroups"` // Rate groups of the default service.
	Cursor     string     `json:"Cursor"`     // Cursor for pagination.
}

type Rates []Rate

type Rate struct {
	// Unique identifier of the rate.
	ID string `json:"Id"`
	// Unique identifier of Rate group where the rate belongs.
	GroupID string `json:"GroupId"`
	// Unique identifier of the base Rate.
	BaseRateID string `json:"BaseRateId"`
	// Unique identifier of the Service.
	ServiceID string `json:"ServiceId"`
	// Whether the rate is still active.
	IsActive bool `json:"IsActive"`
	// Whether the rate is currently available to customers.
	IsEnabled bool `json:"IsEnabled"`
	// Whether the rate is publicly available.
	IsPublic bool `json:"IsPublic"`
	// Name of the rate.
	Name string `json:"Name"`
	// Short name of the rate.
	ShortName string `json:"ShortName"`
	// All translations of the external name of the rate.
	ExternalNames configuration.LocalizedText `json:"ExternalNames"`
}

type RateGroups []RateGroup

type RateGroup struct {
	// Unique identifier of the group.
	ID string `json:"Id"`
	// Unique identifier of the Service.
	ServiceID string `json:"ServiceId"`
	// Whether the rate group is still active.
	IsActive bool `json:"IsActive"`
	// Name of the rate group.
	Name string `json:"Name"`
}

func (s *APIService) NewAllRequest() *AllRequest {
	return &AllRequest{}
}

type AllRequest struct {
	json.BaseRequest
	Limitation json.Limitation `json:"Limitation"`
	// Unique identifiers of the Services from which the rates are requested.
	ServiceIDs []string `json:"ServiceIds"`
	// Extent of data to be returned.
	Extent RateExtent `json:"Extent"`
}

type RateExtent struct {
	// Whether the response should contain rates.
	Rates bool `json:"Rates"`
	// Whether the response should contain rate groups.
	RateGroups bool `json:"RateGroups"`
	// Whether the response should contain rate restrictions.
	RateRestrictions bool `json:"RateRestrictions"`
}

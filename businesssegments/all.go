package businesssegments

import (
	"github.com/omniboost/go-mews/configuration"
	"github.com/omniboost/go-mews/json"
)

const (
	endpointAll = "businesssegments/getAll"
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

	IDs            []string                    `json:"Ids,omitempty"`            // Unique identifiers of the requested Bussiness segment.
	ServiceIDs     []string                    `json:"ServiceIds,omitempty"`     // Unique identifiers of the Services from which the business segments are requested.
	UpdatedUTC     *configuration.TimeInterval `json:"UpdatedUtc,omitempty"`     // Interval in which Business segment was updated.
	ActivityStates ActivityStates              `json:"ActivityStates,omitempty"` // Whether to return only active, only deleted or both records.
	Limitation     json.Limitation             `json:"Limitation,omitempty"`
}

type ActivityStates []ActivityState

type ActivityState string

type AllResponse struct {
	BusinessSegments BusinessSegments `json:"BusinessSegments"`
	Cursor           string           `json:"Cursor"`
}

type BusinessSegments []BusinessSegment

type BusinessSegment struct {
	ID       string `json:"Id"`       // Unique identifier of the segment.
	IsActive bool   `json:"IsActive"` // Whether the business segment is still active.
	Name     string `json:"Name"`     // Name of the segment.
}

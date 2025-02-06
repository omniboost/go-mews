package bills

import (
	"github.com/omniboost/go-mews/configuration"
	base "github.com/omniboost/go-mews/json"
	"github.com/omniboost/go-mews/omitempty"
)

const (
	endpointAll = "bills/getAll"
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
	base.BaseRequest
	Limitation base.Limitation `json:"Limitation,omitempty"`
	// Unique identifiers of the Bills.
	BillIDs []string `json:"BillIds,omitempty"`
	// Unique identifiers of the Customers.
	CustomerIDs []string `json:"CustomerIds,omitempty"`
	// Bill state the bills should be in. If not specified Open and Closed bills are returned.
	State string `json:"State,omitempty"`
	// Interval in which the Bill was closed.
	ClosedUTC configuration.TimeInterval `json:"ClosedUtc,omitempty"`
	// Interval in which the Bill was created.
	CreatedUTC configuration.TimeInterval `json:"CreatedUtc,omitempty"`
	// Interval in which the Bill was updated.
	UpdatedUTC configuration.TimeInterval `json:"UpdatedUtc,omitempty"`
	// Interval in which the Bill is due to be paid.
	DueUTC configuration.TimeInterval `json:"DueUtc,omitempty"`
	// Interval in which the Bill was paid.
	PaidUTC configuration.TimeInterval `json:"PaidUtc,omitempty"`
	// Extent of data to be returned. E.g. it is possible to specify that together with the bills, payments and revenue items should be also returned. If not specified, no extent is used.
	Extent BillExtent `json:"Extent,omitempty"`
}

func (r AllRequest) MarshalJSON() ([]byte, error) {
	return omitempty.MarshalJSON(r)
}

type AllResponse struct {
	Bills  Bills  `json:"Bills"` // The closed bills.
	Cursor string `json:"Cursor"`
}

type BillExtent struct {
	Items bool `json:"Items"`
}

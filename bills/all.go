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
	// Interval in which the Bill was issued.
	IssuedUTC configuration.TimeInterval `json:"IssuedUtc,omitempty"`
	// Interval in which the Bill was paid.
	PaidUTC configuration.TimeInterval `json:"PaidUtc,omitempty"`
	// Interval in which the Bill is due to be paid.
	DueUTC configuration.TimeInterval `json:"DueUtc,omitempty"`
	// Interval in which the Bill was created.
	CreatedUTC configuration.TimeInterval `json:"CreatedUtc,omitempty"`
	// Interval in which the Bill was updated.
	UpdatedUTC configuration.TimeInterval `json:"UpdatedUtc,omitempty"`
	// Unique identifiers of the Bills.
	BillIDs []string `json:"BillIds,omitempty"`
	// Unique identifiers of the Customers.
	CustomerIDs []string `json:"CustomerIds,omitempty"`
	// Bill state the bills should be in. If not specified Open and Closed bills are returned.
	State BillState `json:"State,omitempty"`
	// Type of the bills. If not specified, all types are returned.
	Type BillType `json:"Type,omitempty"`
	// Whether to return regular bills, corrective bills, or both. If BillIds are specified, defaults to both, otherwise defaults to Bill.
	CorrectionState []BillCorrectionType `json:"CorrectionState,omitempty"`
	// Interval in which the Bill was closed.
	ClosedUTC configuration.TimeInterval `json:"ClosedUtc,omitempty"`
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

type BillCorrectionType string

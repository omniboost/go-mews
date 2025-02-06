package productserviceorders

import (
	"time"

	"github.com/omniboost/go-mews/configuration"
	base "github.com/omniboost/go-mews/json"
	"github.com/omniboost/go-mews/omitempty"
)

const (
	endpointAll = "productserviceorders/getAll"
)

// List all productserviceorders
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
	ProductServiceOrders ProductServiceOrders
	Cursor    string    `json:"Cursor"`
}

func (s *APIService) NewAllRequest() *AllRequest {
	return &AllRequest{}
}

type AllRequest struct {
	base.BaseRequest
	Limitation             base.Limitation            `json:"Limitation,omitempty"`
	EnterpriseIDs          []string                   `json:"EnterpriseIds,omitempty"`          // Unique identifiers of the Enterprises
	ProductServiceOrderIDs []string                   `json:"ProductServiceOrderIds,omitempty"` // Unique identifiers of the Product Service Orders
	ServiceIDs             []string                   `json:"ServiceIds,omitempty"`             // Unique identifiers of the Services
	States                 []string                   `json:"States,omitempty"`                 // A list of product service order states to filter by.
	UpdatedUTC             configuration.TimeInterval `json:"UpdatedUtc,omitempty"`             // Interval in wich the product service orders were updated
}

func (r AllRequest) MarshalJSON() ([]byte, error) {
	return omitempty.MarshalJSON(r)
}

type ProductServiceOrders []ProductServiceOrder

type ProductServiceOrder struct {
	ID                string              `json:"Id"`                // Unique identifier of the product service order.
	ServiceID         string              `json:"ServiceId"`         // Unique identifier of the Service that the product service order is made against.
	AccountID         string              `json:"AccountId"`         // Unique identifier of the Customer or a Company who owns the product service order.
	AccountType       string              `json:"AccountType"`       // A discriminator specifying the type of the Account, e.g. customer or company.
	CreatorProfileID  string              `json:"CreatorProfileId"`  // Unique identifier of the user who created the order item.
	UpdateProfileID   string              `json:"UpdateProfileId"`   // Unique identifier of the user who updated the order item.
	StartUTC          time.Time           `json:"StartUtc"`          // Product service order start in UTC timezone in ISO 8601 format.
	EndUTC            time.Time           `json:"EndUtc"`            // Product service order end in UTC timezone in ISO 8601 format.
	BookerID          string              `json:"BookerId"`          // Unique identifier of the customer on whose behalf the product service order was made.
	Number            string              `json:"Number"`            // Confirmation number of the service order in Mews.
	State             ServiceOrderState   `json:"State"`             // State of the product service order.
	Origin            ServiceOrderOrigin  `json:"Origin"`            // Origin of the product service order.
	OriginDetails     string              `json:"OriginDetails"`     // Details about the product service order origin.
	CreatedUTC        time.Time           `json:"CreatedUtc"`        // Creation date and time of the product service order in UTC timezone in ISO 8601 format.
	UpdatedUTC        time.Time           `json:"UpdatedUtc"`        // Last update date and time of the product service order in UTC timezone in ISO 8601 format.
	CancelledUTC      time.Time           `json:"CancelledUtc"`      // Cancellation date and time of the product service order in UTC timezone in ISO 8601 format.
	VoucherID         string              `json:"VoucherId"`         // Unique identifier of the voucher that has been used to create the product service order.
	BusinessSegmentID string              `json:"BusinessSegmentId"` // Unique identifier of the product service order Business segment.
	Options           ServiceOrderOptions `json:"Options"`           // Options of the service order.
}

type ServiceOrderOptions struct {
	OwnerCheckedIn         bool `json:"OwnerCheckedIn"`         // Owner of the reservation checked in.
	AllCompanionsCheckedIn bool `json:"AllCompanionsCheckedIn"` // All companions of the reservation checked in.
	AnyCompanionCheckedIn  bool `json:"AnyCompanionCheckedIn"`  // Any of the companions of the reservation checked in.
	ConnectorCheckIn       bool `json:"ConnectorCheckIn"`       // Check in was done via Connector API.
}

type ServiceOrderState string
type ServiceOrderOrigin string

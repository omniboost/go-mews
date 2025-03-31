package counters

import (
	"time"

	"github.com/omniboost/go-mews/configuration"
	"github.com/omniboost/go-mews/json"
)

const (
	endpointAll = "counters/getAll"
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

type AllResponse struct {
	Counters []Counter
	Cursor   string `json:"Cursor"`
}

func (s *Service) NewAllRequest() *AllRequest {
	return &AllRequest{
		Limitation: &json.Limitation{},
	}
}

type AllRequest struct {
	json.BaseRequest

	EnterpriseIDs []string                    `json:"EnterpriseIds,omitempty"` // Unique identifiers of the Enterprises. If not specified, the operation returns data for all enterprises within scope of the Access Token.
	UpdatedUTC    *configuration.TimeInterval `json:"UpdatedUtc,omitempty"`    // Interval in which Counter was updated.
	Type          CounterType                 `json:"Type,omitempty"`          //Type of the counter. If not specified, the operation returns all types.
	Limitation    *json.Limitation            `json:"Limitation,omitempty"`    // Limitation on the quantity of data returned.
}

type Counter struct {
	ID         string      `json:"Id"`         // Unique identifier of the counter.
	Name       string      `json:"Name"`       // Name of the counter
	CreatedUTC time.Time   `json:"CreatedUtc"` // Creation date and time of the counter in UTC timezone in ISO 8601 format.
	UpdatedUTC time.Time   `json:"UpdatedUtc"` // Last update date and time of the counter in UTC timezone in ISO 8601 format.
	IsDefault  bool        `json:"IsDefault"`  // Whether the counter is used by default.
	Value      int         `json:"Value"`      // Current value the counter.
	Format     string      `json:"Format"`     // Format the counter is displayed in.
	Type       CounterType `json:"Type"`       // Type of the counter
}

type CounterType string

var (
	CounterCounterType                 CounterType = "Counter"
	AvailabilityBlockCounterType       CounterType = "AvailabilityBlockCounter"
	BillCounterType                    CounterType = "BillCounter"
	BillPreviewCounterType             CounterType = "BillPreviewCounter"
	FiscalCounterType                  CounterType = "FiscalCounter"
	ProformaCounterType                CounterType = "ProformaCounter"
	RegistrationCardCounterType        CounterType = "RegistrationCardCounter"
	CorrectionBillCounterType          CounterType = "CorrectionBillCounter"
	ServiceOrderCounterType            CounterType = "ServiceOrderCounter"
	PaymentConfirmationBillCounterType CounterType = "PaymentConfirmationBillCounter"
	CreditNoteBillCounterType          CounterType = "CreditNoteBillCounter"
)

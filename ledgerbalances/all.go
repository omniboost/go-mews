package ledgerbalances

import (
	"github.com/omniboost/go-mews/configuration"
	base "github.com/omniboost/go-mews/json"
	"github.com/omniboost/go-mews/omitempty"
)

const (
	endpointAll = "ledgerBalances/getAll"
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
	EnterpriseIDs []string                   `json:"EnterpriseIds,omitempty"` // Unique identifiers of the Enterprises. If not specified, the operation returns data for all enterprises within scope of the Access Token.
	LedgerTypes   []LedgerType               `json:"LedgerTypes"`
	Date          configuration.DateInterval `json:"Date,omitempty"` // Interval in which Credit card was updated.
	Limitation    base.Limitation            `json:"Limitation,omitempty"`
}

func (r AllRequest) MarshalJSON() ([]byte, error) {
	return omitempty.MarshalJSON(r)
}

type AllResponse struct {
	LedgerBalances LedgerBalances `json:"LedgerBalances"`
	Cursor         string         `json:"Cursor"`
}

type LedgerType string

type LedgerBalances []LedgerBalance

type LedgerBalance struct {
	EnterpriseID   string    `json:"EnterpriseId"`
	Date           base.Date `json:"Date"`
	LedgerType     string    `json:"LedgerType"`
	OpeningBalance Balance   `json:"OpeningBalance"`
	ClosingBalance Balance   `json:"ClosingBalance"`
}

type Balance struct {
	Currency   string        `json:"Currency"`
	NetValue   float64       `json:"NetValue"`
	GrossValue float64       `json:"GrossValue"`
	TaxValues  []interface{} `json:"TaxValues"`
	Breakdown  struct {
		Items []struct {
			TaxRateCode string  `json:"TaxRateCode"`
			NetValue    float64 `json:"NetValue"`
			TaxValue    float64 `json:"TaxValue"`
		} `json:"Items"`
	} `json:"Breakdown"`
}

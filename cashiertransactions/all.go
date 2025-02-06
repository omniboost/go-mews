package cashiertransactions

import (
	"time"

	"github.com/omniboost/go-mews/accountingitems"
	"github.com/omniboost/go-mews/configuration"
	"github.com/omniboost/go-mews/json"
	"github.com/omniboost/go-mews/omitempty"
)

const (
	endpointAll = "cashierTransactions/getAll"
)

// List all cashier transactions
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
	CreatedUTC configuration.TimeInterval `json:"CreatedUtc"`
}

func (r AllRequest) MarshalJSON() ([]byte, error) {
	return omitempty.MarshalJSON(r)
}

type AllResponse struct {
	CashierTransactions CashierTransactions `json:"CashierTransactions"`
}

type CashierTransactions []CashierTransaction

type CashierTransaction struct {
	ID         string                 `json:"Id"`         // Unique identifier of the transaction.
	CashierID  string                 `json:"CashierId"`  // Unique identifier of the Cashier.
	PaymentID  string                 `json:"PaymentId"`  // Unique identifier of the corresponding payment item.
	CreatedUTC time.Time              `json:"CreatedUtc"` // Creation date and time of the transaction.
	Number     string                 `json:"Number"`     // Number of the transaction.
	Notes      string                 `json:"Notes"`      // Additional notes of the transaction.
	Amount     accountingitems.Amount `json:"Amount"`     // Value of the transaction.
}

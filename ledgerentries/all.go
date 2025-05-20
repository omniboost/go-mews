package ledgerentries

import (
	"time"

	base "github.com/omniboost/go-mews/json"
	"github.com/omniboost/go-mews/omitempty"
)

const (
	endpointAll = "ledgerEntries/getAll"
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
	EnterpriseIDs []string     `json:"EnterpriseIds,omitempty"` // Unique identifiers of the Enterprises. If not specified, the operation returns data for all enterprises within scope of the Access Token.
	LedgerTypes   []LedgerType `json:"LedgerTypes"`
	PostingDate   struct {
		Start base.Date `json:"Start"`
		End   base.Date `json:"End"`
	} `json:"PostingDate"` // Interval in which Credit card was updated.
	Limitation base.Limitation `json:"Limitation,omitempty"`
}

func (r AllRequest) MarshalJSON() ([]byte, error) {
	return omitempty.MarshalJSON(r)
}

type AllResponse struct {
	LedgerEntries LedgerEntries `json:"LedgerEntries"`
	Cursor        string        `json:"Cursor"`
}

type LedgerType string

var (
	LedgerTypeRevenue    LedgerType = "Revenue"
	LedgerTypeTax        LedgerType = "Tax"
	LedgerTypePayment    LedgerType = "Payment"
	LedgerTypeDeposit    LedgerType = "Deposit"
	LedgerTypeGuest      LedgerType = "Guest"
	LedgerTypeCity       LedgerType = "City"
	LedgerTypeNonRevenue LedgerType = "NonRevenue"
)

type LedgerEntries []LedgerEntry

type LedgerEntry struct {
	ID                   string    `json:"Id"`
	EnterpriseID         string    `json:"EnterpriseId"`
	AccountID            string    `json:"AccountId"`
	BillID               string    `json:"BillId"`
	AccountingCategoryID string    `json:"AccountingCategoryId"`
	AccountingItemID     string    `json:"AccountingItemId"`
	AccountingItemType   string    `json:"AccountingItemType"`
	LedgerType           string    `json:"LedgerType"`
	LedgerEntryType      string    `json:"LedgerEntryType"`
	PostingDate          string    `json:"PostingDate"`
	Value                float64   `json:"Value"`
	NetBaseValue         any       `json:"NetBaseValue"`
	TaxRateCode          any       `json:"TaxRateCode"`
	CreatedUTC           time.Time `json:"CreatedUtc"`
}

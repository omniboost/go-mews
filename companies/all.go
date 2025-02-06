package companies

import (
	"time"

	"github.com/omniboost/go-mews/configuration"
	base "github.com/omniboost/go-mews/json"
	"github.com/omniboost/go-mews/omitempty"
)

const (
	endpointAll = "companies/getAll"
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
	Limitation          base.Limitation            `json:"Limitation,omitempty"`
	IDs                 []string                   `json:"Ids,omitempty"`   // Unique identifiers of Companies.
	Names               []string                   `json:"Names,omitempty"` // Names of Companies.
	ExternalIdentifiers []string                   `json:"ExternalIdentifiers,omitempty"`
	CreatedUTC          configuration.TimeInterval `json:"CreatedUtc,omitempty"` // Interval of Company creation date and time.
	UpdatedUTC          configuration.TimeInterval `json:"UpdatedUtc,omitempty"` // Interval of Company last update date and time.
}

func (r AllRequest) MarshalJSON() ([]byte, error) {
	return omitempty.MarshalJSON(r)
}

type AllResponse struct {
	Companies Companies `json:"companies"`
	Cursor    string    `json:"Cursor"`
}

type CompanyOptions struct {
	Invoiceable                     bool `json:"Invoiceable"`
	AddFeesToInvoices               bool `json:"AddFeesToInvoices"`
	AddTaxDeductedPaymentToInvoices bool `json:"AddTaxDeductedPaymentToInvoices"` // Whether tax-deducted payments should be automatically added to invoices.
}

type Companies []Company

type Company struct {
	ID                          string                `json:"Id"`                          // Unique identifier of the company.
	Name                        string                `json:"Name"`                        // Name of the company.
	Number                      int                   `json:"Number"`                      // Unique number of the company.
	IsActive                    bool                  `json:"IsActive"`                    // Whether the company is still active.
	Identifier                  string                `json:"Identifier"`                  // Identifier of the company (e.g. legal identifier).
	TaxIdentifier               string                `json:"TaxIdentifier"`               // Tax identification number of the company.
	AdditionalTaxIdentifier     string                `json:"AdditionalTaxIdentifier"`     // Additional tax identifer of the company.
	ElectronicInvoiceIdentifier string                `json:"ElectronicInvoiceIdentifier"` // Electronic invoice identifer of the company.
	ExternalIdentifier          string                `json:"ExternalIdentifier"`
	AccountingCode              string                `json:"AccountingCode"` // Accounting code of the company.
	MotherCompanyID             string                `json:"MotherCompanyId"`
	BillingCode                 string                `json:"BillingCode"` // Billing code of the company.
	Address                     configuration.Address `json:"Address"`     // Address of the company (if it is non-empty, otherwise null).
	InvoiceDueInterval          base.Duration         `json:"InvoiceDueInterval"`
	CreatedUtc                  time.Time             `json:"CreatedUtc"`
	UpdatedUtc                  time.Time             `json:"UpdatedUtc"`
	Iata                        string                `json:"Iata"`
	Telephone                   string                `json:"Telephone"`
	InvoicingEmail              string                `json:"InvoicingEmail"`
	ContactPerson               string                `json:"ContactPerson"` // Email for issuing invoices to the company.
	Contact                     string                `json:"Contact"`
	Notes                       string                `json:"Notes"`
	TaxIdentificationNumber     string                `json:"TaxIdentificationNumber"`
	Options                     CompanyOptions        `json:"Options"`
}

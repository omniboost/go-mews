package companies

import (
	"github.com/omniboost/go-mews/configuration"
	"github.com/omniboost/go-mews/json"
)

const (
	endpointAdd = "companies/add"
)

// Add customer
func (s *Service) Add(requestBody *AddRequest) (*AddResponse, error) {
	// @TODO: create wrapper?
	if err := s.Client.CheckTokens(); err != nil {
		return nil, err
	}

	apiURL, err := s.Client.GetApiURL(endpointAdd)
	if err != nil {
		return nil, err
	}

	responseBody := &AddResponse{}
	httpReq, err := s.Client.NewRequest(apiURL, requestBody)
	if err != nil {
		return nil, err
	}

	_, err = s.Client.Do(httpReq, responseBody)
	return responseBody, err
}

func (s *Service) NewAddRequest() *AddRequest {
	return &AddRequest{}
}

type AddRequest struct {
	json.BaseRequest

	// Unique identifier of the chain
	ChainID string `json:"string,omitempty"`
	// Name of the company
	Name string `json:"Name"`
	// Options of the company
	Options CompanyOptions `json:"Options"`
	// Unique identifier of the mother company
	MotherCompanyID string `json:"MotherCompanyId,omitempty"`
	// Identifier of the company
	Identifier string `json:"Identifier,omitempty"`
	// Tax identification number of the company
	TaxIdentifier string `json:"TaxIdentifier,omitempty"`
	// Additional tax identifier of the company
	AdditionalTaxIdentifier string `json:"AdditionalTaxIdentifier,omitempty"`
	// Billing code of the company
	BillingCode string `json:"BillingCode,omitempty"`
	// Accounting code of the company
	AccountingCode string `json:"AccountingCode,omitempty"`
	// Address of the company
	Address *configuration.Address `json:"Address,omitempty"`
	// The maximum time, when the invoice has to be paid in ISO 8601 duration format
	InvoiceDueInterval *json.Duration `json:"InvoiceDueInterval,omitempty"`
	// Contact person of the company
	ContactPerson string `json:"ContactPerson,omitempty"`
	// Other contact details
	Contact string `json:"Contact,omitempty"`
	// Notes of the company
	Notes string `json:"Notes,omitempty"`
	// Iata of the company
	Iata string `json:"Iata,omitempty"`
	// The internal segmentation of a company
	Department string `json:"Department,omitempty"`
	// The DUN & bradstreet unique 9-digit DUNS number
	DunsNumber string `json:"DunsNumber,omitempty"`
	// Credit rating to define creditworthiness of the company
	CreditRating string `json:"CreditRating,omitempty"`
	// Identifier of the company from external system
	ExternalIdentifier string `json:"ExternalIdentifier,omitempty"`
	// External system identifier
	ReferenceIdentifier string `json:"ReferenceIdentifier,omitempty"`
	// The website URL of the company
	WebsiteURL string `json:"WebsiteUrl,omitempty"`
}

type AddResponse struct {
	Companies Companies `json:"companies"`
}

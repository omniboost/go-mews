package customers

import (
	"time"

	"github.com/omniboost/go-mews/configuration"
	base "github.com/omniboost/go-mews/json"
	"github.com/omniboost/go-mews/omitempty"
	"github.com/omniboost/go-mews/services"
)

const (
	endpointAll = "customers/getAll"
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
	return &AllRequest{
		Extent: CustomersExtent{
			Customers: true,
			Addresses: true,
		},
	}
}

type AllRequest struct {
	base.BaseRequest
	Limitation     base.Limitation            `json:"Limitation,omitempty"`
	CustomerIDs    []string                   `json:"CustomerIds,omitempty"`    // Unique identifiers of Customers. Required if no other filter is provided.
	Emails         []string                   `json:"Emails,omitempty"`         // Emails of the Customers.
	FirstNames     []string                   `json:"FirstNames,omitempty"`     // First names of the Customers.
	LastNames      []string                   `json:"LastNames,omitempty"`      // Last names of the Customers.
	LoyaltyCodes   []string                   `json:"LoyaltyCodes,omitempty"`   // Loyalty codes of the Customers.
	CreatedUTC     configuration.TimeInterval `json:"CreatedUtc,omitempty"`     // Interval in which Customer was created.
	UpdatedUTC     configuration.TimeInterval `json:"UpdatedUtc,omitempty"`     // Interval in which Customer was updated.
	DeletedUTC     configuration.TimeInterval `json:"DeletedUtc,omitempty"`     // Interval in which Customer was deleted. ActivityStates value Deleted should be provided with this filter to get expected results.
	ActivityStates services.ActivityStates    `json:"ActivityStates,omitempty"` // Whether return only active, only deleted or both records.
	Extent         CustomersExtent            `json:"Extent,omitempty"`         // Extent of data to be returned.
}

func (r AllRequest) MarshalJSON() ([]byte, error) {
	return omitempty.MarshalJSON(r)
}

type AllResponse struct {
	Customers Customers `json:"customers"`
	Cursor    string    `json:"Cursor"`
}

type CustomersExtent struct {
	Customers bool `json:"Customers"` // Whether the response should contain information about customers.
	Documents bool `json:"Document"`  // Whether the response should contain identity documents of customers.
	Addresses bool `json:"Addresses"` // Whether the response should contain addresses of customers.
}

type Customers []Customer

type Customer struct {
	ID                      string                `json:"Id"`                      // Unique identifier of the customer.
	Number                  string                `json:"Number"`                  // Number of the customer.
	FirstName               string                `json:"FirstName"`               // First name of the customer.
	LastName                string                `json:"LastName"`                // Last name of the customer.
	SecondLastName          string                `json:"SecondLastName"`          // Second last name of the customer.
	Title                   Title                 `json:"Title"`                   // Title prefix of the customer.
	Sex                     Sex                   `json:"Sex"`                     // Sex of the customer.
	NationalityCode         string                `json:"NationalityCode"`         // ISO 3166-1 alpha-2 country code (two letter country code) of the nationality.
	LanguageCode            string                `json:"LanguageCode"`            // Language and culture code of the customers preferred language. E.g. en-US or fr-FR.
	BirthDate               string                `json:"BirthDate"`               // Date of birth in ISO 8601 format.
	BirthDateUTC            time.Time             `json:"BirthDateUtc"`            // ??
	BirthPlace              string                `json:"BirthPlace"`              // Place of birth.
	Email                   string                `json:"Email"`                   // Email address of the customer.
	Phone                   string                `json:"Phone"`                   // Phone number of the customer (possibly mobile).
	TaxIdentificationNumber string                `json:"TaxIdentificationNumber"` // tax id customer
	LoyaltyCode             string                `json:"LoyaltyCode"`             // Loyalty code of the customer.
	AccountingCode          string                `json:"AccountingCode"`          // Accounting code of the customer.
	BillingCode             string                `json:"BillingCode"`             // Billing code of the customer.
	Notes                   string                `json:"Notes"`                   // ??
	Classifications         []Classification      `json:"Classifications"`         // Classifications of the customer.
	Options                 Options               `json:"Options"`                 // Options of the customer.
	Address                 configuration.Address `json:"Address"`                 // Address of the customer.
	CreatedUTC              time.Time             `json:"CreatedUtc"`              // Creation date and time of the customer in UTC timezone in ISO 8601 format.
	UpdatedUTC              time.Time             `json:"UpdatedUtc"`              // Last update date and time of the customer in UTC timezone in ISO 8601 format.
	ItalianDestinationCode  string                `json:"ItalianDestinationCode"`  // Value of Italian destination code.
	ItalianFiscalCode       string                `json:"ItalianFiscalCode"`       // Value of Italian fiscal code.
	CompanyID               string                `json:"CompanyId"`               // Unique identifier of Company the customer is associated with.

	// Deprecated
	Passport       Document       `json:"Passport"`       // Passport details of the customer.
	IdentityCard   Document       `json:"IdentityCard"`   // IdentityCard details for Customer.
	Visa           Document       `json:"Visa"`           // Visa details for Customer.
	CategoryID     string         `json:"CategoryId"`     // ??
	CitizenNumber  string         `json:"CitizenNumber"`  // ??
	FatherName     string         `json:"FatherName"`     // ??
	MotherName     string         `json:"MotherName"`     // ??
	Occupation     string         `json:"Occupation"`     // ??
	DriversLicense DriversLicense `json:"DriversLicense"` // Drivers license  details of the customer.
}

type Title string

const (
	TitleMister Title = "Mister"
	TitleMiss   Title = "Miss"
	TitleMisses Title = "Missed"
)

type Sex string

const (
	SexMale   Title = "Male"
	SexFemale Title = "Female"
)

type Gender string

const (
	GenderMale   Gender = "Male"
	GenderFemale Gender = "Female"
)

type Document struct {
	Number             string    `json:"Number"`             // Number of the document (e.g. passport number).
	Issuance           base.Date `json:"Issuance"`           // Date of issuance in ISO 8601 format.
	Expiration         base.Date `json:"Expiration"`         // Expiration date in ISO 8601 format.
	ExpirationUTC      time.Time `json:"ExpirationUtc"`      // ??
	IssuanceUTC        time.Time `json:"IssuanceUtc"`        // ??
	IssuingCountryCode string    `json:"IssuingCountryCode"` // ISO 3166-1 code of the Country.
}

type Classification string

type Options []string

type DriversLicense struct {
	Expiration         base.Date `json:"Expiration"`
	ExpirationUTC      time.Time `json:"ExpirationUtc"`
	Issuance           base.Date `json:"Issuance"`
	IssuanceUTC        time.Time `json:"IssuanceUtc"`
	IssuingCountryCode string    `json:"IssuingCountryCode"`
	Number             string    `json:"Number"`
}

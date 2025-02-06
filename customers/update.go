package customers

import (
	"github.com/omniboost/go-mews/configuration"
	"github.com/omniboost/go-mews/json"
)

const (
	endpointUpdate = "customers/update"
)

// Update customer
func (s *Service) Update(requestBody *UpdateRequest) (*UpdateResponse, error) {
	// @TODO: create wrapper?
	if err := s.Client.CheckTokens(); err != nil {
		return nil, err
	}

	apiURL, err := s.Client.GetApiURL(endpointUpdate)
	if err != nil {
		return nil, err
	}

	responseBody := &UpdateResponse{}
	httpReq, err := s.Client.NewRequest(apiURL, requestBody)
	if err != nil {
		return nil, err
	}

	_, err = s.Client.Do(httpReq, responseBody)
	return responseBody, err
}

func (s *Service) NewUpdateRequest() *UpdateRequest {
	return &UpdateRequest{}
}

type UpdateRequest struct {
	json.BaseRequest

	// Unique identifier of the Customer.
	CustomerID string `json:"CustomerId"`
	// New first name.
	FirstName string `json:"FirstName"`
	// New last name.
	LastName string `json:"LastName"`
	// New second last name.
	SecondLastName string `json:"SecondLastName"`
	// New title.
	Title string `json:"Title"`
	// New birth date in ISO 8601 format.
	BirthDate *json.Date `json:"BirthDate"`
	// New birth place.
	BirthPlace string `json:"BirthPlace"`
	// ISO 3166-1 code of the Country.
	NationalityCode string `json:"NationalityCode"`
	// New email address.
	Email string `json:"Email"`
	// New phone number.
	Phone string `json:"Phone"`
	// Loyalty code of the customer.
	LoyaltyCode string `json:"LocaltyCode"`
	// Internal notes about the customer. Old value will be overwritten.
	Notes string `json:"Notes"`
	// New identity card details.
	IdentityCard *Document `json:"IdentityCard"`
	// New passport details.
	Passport *Document `json:"Passport"`
	// New visa details.
	Visa *Document `json:"Visa"`
	// New drivers license details.
	DriversLicense *Document `json:"DriversLicense"`
	// New address details.
	Address *configuration.Address `json:"Address"`
	// New classifications of the customers
	Classifications []Classification `json:"Classifications"`
	// Options of the customer.
	Options *Options `json:"Options"`
}

type UpdateResponse Customer

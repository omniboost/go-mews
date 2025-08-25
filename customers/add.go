package customers

import (
	"github.com/omniboost/go-mews/configuration"
	"github.com/omniboost/go-mews/json"
)

const (
	endpointAdd = "customers/add"
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

	// Overwrite existing
	OverwriteExisting bool `json:"OverwriteExisting,omitempty"`
	// New first name.
	FirstName string `json:"FirstName,omitempty"`
	// New last name.
	LastName string `json:"LastName"`
	// New second last name.
	SecondLastName string `json:"SecondLastName,omitempty"`
	// New title.
	Title string `json:"Title,omitempty"`
	// Sex
	Sex string `json:"Sex,omitempty"`
	// ISO 3166-1 code of the Country.
	NationalityCode string `json:"NationalityCode,omitempty"`
	// New birth date in ISO 8601 format.
	BirthDate *json.Date `json:"BirthDate,omitempty"`
	// New birth place.
	BirthPlace string `json:"BirthPlace,omitempty"`
	// New email address.
	Email string `json:"Email,omitempty"`
	// New phone number.
	Phone string `json:"Phone,omitempty"`
	// Loyalty code of the customer.
	LoyaltyCode string `json:"LocaltyCode,omitempty"`
	// Internal notes about the customer. Old value will be overwritten.
	Notes string `json:"Notes,omitempty"`
	// New address details.
	Address *configuration.Address `json:"Address,omitempty"`
	// Italian Destination Code
	ItalianDestinationCode string `json:"ItalianDestinationCode,omitempty"`
}

type AddResponse Customer

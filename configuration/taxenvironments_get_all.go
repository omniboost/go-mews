package configuration

import (
	"time"

	"github.com/omniboost/go-mews/json"
)

const (
	endpointTaxenvironmentsGetAll = "taxenvironments/getAll"
)

var (
	TaxRateStrategyDiscriminatorFlat     TaxRateStrategyDiscriminator = "Flat"
	TaxRateStrategyDiscriminatorRelative TaxRateStrategyDiscriminator = "Relative"
)

// List all products
func (s *Service) TaxenvironmentsGetAll(requestBody *TaxenvironmentsGetAllRequest) (*TaxenvironmentsGetAllResponse, error) {
	// @TODO: create wrapper?
	if err := s.Client.CheckTokens(); err != nil {
		return nil, err
	}

	apiURL, err := s.Client.GetApiURL(endpointTaxenvironmentsGetAll)
	if err != nil {
		return nil, err
	}

	responseBody := &TaxenvironmentsGetAllResponse{}
	httpReq, err := s.Client.NewRequest(apiURL, requestBody)
	if err != nil {
		return nil, err
	}

	_, err = s.Client.Do(httpReq, responseBody)
	return responseBody, err
}

type TaxenvironmentsGetAllResponse struct {
	TaxEnvironments TaxEnvironments `json:"TaxEnvironments"` // The supported tax environments.
}

type TaxEnvironments []TaxEnvironment

type TaxEnvironment struct {
	Code             string        `json:"Code"`             // Code of the tax environment.
	CountryCode      string        `json:"CountryCode"`      // ISO 3166-1 alpha-3 code, e.g. USA or GBR.
	TaxationCodes    TaxationCodes `json:"TaxationCodes"`    // Codes of the Taxations that are used by this environment.
	ValidityStartUTC time.Time     `json:"ValidityStartUtc"` // If specified, marks the start of the validity interval in UTC timezone in ISO 8601 format.
	ValidityEndUTC   time.Time     `json:"ValidityEndUtc"`   // If specified, marks the end of the validity interval in UTC timezone in ISO 8601 format.

}

type TaxationCodes []string

type TaxationCode string

func (s *Service) NewTaxenvironmentsGetAllRequest() *TaxenvironmentsGetAllRequest {
	return &TaxenvironmentsGetAllRequest{}
}

type TaxenvironmentsGetAllRequest struct {
	json.BaseRequest
}

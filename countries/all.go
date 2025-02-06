package countries

import (
	"github.com/omniboost/go-mews/json"
	"github.com/omniboost/go-mews/omitempty"
)

const (
	endpointAll = "countries/getAll"
)

// List all countries
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
}

func (r AllRequest) MarshalJSON() ([]byte, error) {
	return omitempty.MarshalJSON(r)
}

type AllResponse struct {
	Countries           Countries           `json:"Countries"`
	CountrySubdivisions CountrySubdivisions `json:"CountrySubdivisions"`
	CountryAlliances    CountryAlliances    `json:"CountryAlliances"`
}

type Countries []Country

type Country struct {
	// SO 3166-1 alpha-2 code, e.g. US or GB.
	Code string `json:"Code"`
	// English name of the country.
	EnglishName string `json:"EnglishName"`
}

type CountrySubdivisions []CountrySubdivision

type CountrySubdivision struct {
	// ISO 3166-2 code of the administrative division, e.g AU-QLD.
	Code string `json:"Code"`
	// ISO 3166-1 code of the Country
	CountryCode string `json:"CountryCode"`
	// English name of the country subdivision.
	EnglishName string `json:"EnglishName"`
}

type CountryAlliances []CountryAlliance

type CountryAlliance struct {
	// Code of the alliance, e.g EU, SCHENGEN, SCHENGEN-VISA.
	Code string `json:"Code"`
	// English name of the alliance.
	EnglishName string `json:"EnglishName"`
	// ISO 3166-1 codes of the member countries.
	CountryCodes []string `json:"CountryCodes"`
}

package finance

import (
	"github.com/omniboost/go-mews/json"
)

const (
	endpointExchangeRatesGetAll = "exchangeRates/getAll"
)

// Returns configuration of the enterprise and the client.
func (s *Service) ExchangeRatesGetAll(requestBody *ExchangeRatesGetAllRequest) (*ExchangeRatesGetAllResponse, error) {
	// @TODO: create wrapper?
	if err := s.Client.CheckTokens(); err != nil {
		return nil, err
	}

	apiURL, err := s.Client.GetApiURL(endpointExchangeRatesGetAll)
	if err != nil {
		return nil, err
	}

	responseBody := &ExchangeRatesGetAllResponse{}
	httpReq, err := s.Client.NewRequest(apiURL, requestBody)
	if err != nil {
		return nil, err
	}

	_, err = s.Client.Do(httpReq, responseBody)
	return responseBody, err
}

func (s *Service) NewExchangeRatesGetAllRequest() *ExchangeRatesGetAllRequest {
	return &ExchangeRatesGetAllRequest{}
}

type ExchangeRatesGetAllRequest struct {
	json.BaseRequest
}

type ExchangeRatesGetAllResponse struct {
	ExchangeRates ExchangeRates `json:"ExchangeRates"`
}

type ExchangeRates []ExchangeRate

type ExchangeRate struct {
	SourceCurrency string  `json:"SourceCurrency"` // ISO-4217 code of the source Currency.
	TargetCurrency string  `json:"TargetCurrency"` // ISO-4217 code of the target Currency.
	Value          float64 `json:"Value"`          // The exchange rate from the source currency to the target currency.
}

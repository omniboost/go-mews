package configuration

import (
	"time"

	base "github.com/omniboost/go-mews/json"
)

const (
	endpointGet = "configuration/get"
)

// Returns configuration of the enterprise and the client.
func (s *Service) Get(requestBody *GetRequest) (*GetResponse, error) {
	// @TODO: create wrapper?
	if err := s.Client.CheckTokens(); err != nil {
		return nil, err
	}

	apiURL, err := s.Client.GetApiURL(endpointGet)
	if err != nil {
		return nil, err
	}

	responseBody := &GetResponse{}
	httpReq, err := s.Client.NewRequest(apiURL, requestBody)
	if err != nil {
		return nil, err
	}

	_, err = s.Client.Do(httpReq, responseBody)
	return responseBody, err
}

func (s *Service) NewGetRequest() *GetRequest {
	return &GetRequest{}
}

type GetRequest struct {
	base.BaseRequest
}

type GetResponse struct {
	Enterprise Enterprise `json:"Enterprise"`
	NowUtc     time.Time  `json:"NowUtc"` // NEW
}

type Enterprise struct {
	ID                      string        `json:"Id"`                      // Unique identifier of the enterprise.
	CreatedUTC              time.Time     `json:"CreatedUtc"`              // Creation date and time of the enterprise in UTC timezone in ISO 8601 format.
	Name                    string        `json:"Name"`                    // Name of the enterprise.
	WebsiteURL              string        `json:"WebsiteUrl"`              // URL of the enterprise website.
	Email                   string        `json:"Email"`                   // Email address of the enterprise.
	Phone                   string        `json:"Phone"`                   // Phone number of the enterprise.
	TimeZoneIdentifier      string        `json:"TimeZoneIdentifier"`      // IANA timezone identifier of the enterprise.
	EditableHistoryInterval base.Duration `json:"EditableHistoryInterval"` // Editable history interval in ISO 8601 duration format.
	Address                 Address       `json:"Address"`                 // Address of the enterprise.
	Currencies              Currencies    `json:"Currencies"`              // Currencies accepted by the enterprise.
	ChainID                 string        `json:"ChainId"`                 // NEW
	CoverImageId            string        `json:"CoverImageId"`            // NEW
	DefaultLanguageCode     string        `json:"DefaultLanguageCode"`     // NEW
	LegalEnvironmentCode    string        `json:"LegalEnvironmentCode"`    // NEW
	LogoImageID             string        `json:"LogoImageId"`             // NEW
}

type Currencies []Currency

type Currency struct {
	Currency  string `json:"Currency"`
	IsDefault bool   `json:"IsDefault"`
	IsEnabled bool   `json:"IsEnabled"`
}

type Address struct {
	Line1                  string      `json:"Line1,omitempty"`                  // First line of the address.
	Line2                  string      `json:"Line2,omitempty"`                  // Second line of the address.
	City                   string      `json:"City,omitempty"`                   // The city.
	PostalCode             string      `json:"PostalCode,omitempty"`             // Postal code.
	CountryCode            string      `json:"CountryCode,omitempty"`            // ISO 3166-1 alpha-2 country code (two letter country code).
	CountrySubdivisionCode string      `json:"CountrySubdivisionCode,omitempty"` // ISO 3166-2 code of the administrative division, e.g. DE-BW.
	Latitude               interface{} `json:"Latitude,omitempty"`
	Longitude              interface{} `json:"Longitude,omitempty"`
}

type LocalizedText map[string]string

func (t LocalizedText) Default() string {
	if v, ok := t["en-US"]; ok {
		return v
	}

	for _, v := range t {
		return v
	}

	return ""
}

type CurrencyValue struct {
	Currency string  `json:"Currency"` // ISO-4217 currency code, e.g. EUR or USD.
	Value    float64 `json:"Value"`    // Amount in the currency (including tax if taxed).
	TaxRate  float64 `json:"TaxRate"`  // Tax rate in case the item is taxed (e.g. 0.21).
	Tax      float64 `json:"Tax"`      // Tax value in case the item is taxed.
}

type TimeInterval struct {
	// Start of the interval in UTC timezone in ISO 8601 format.
	StartUTC time.Time `json:"StartUtc"`
	// End of the interval in UTC timezone in ISO 8601 format.
	EndUTC time.Time `json:"EndUtc"`
}

func (i TimeInterval) IsEmpty() bool {
	return i.StartUTC.IsZero() && i.EndUTC.IsZero()
}

type DateInterval struct {
	Start base.Date `json:"Start"`
	End   base.Date `json:"End"`
}

func (i DateInterval) IsEmpty() bool {
	return i.Start.IsZero() && i.End.IsZero()
}

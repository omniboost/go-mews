package products

import (
	"github.com/omniboost/go-mews/configuration"
	"github.com/omniboost/go-mews/json"
	base "github.com/omniboost/go-mews/json"
	"github.com/omniboost/go-mews/services"
)

const (
	endpointAll = "products/getAll"
)

// List all products
func (s *APIService) All(requestBody *AllRequest) (*AllResponse, error) {
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

type AllResponse struct {
	Products Products
	Cursor   string `json:"Cursor"`
}

func (s *APIService) NewAllRequest() *AllRequest {
	return &AllRequest{
		IncludeDefault: true,
	}
}

type AllRequest struct {
	base.BaseRequest

	ProductIDs     []string                    `json:"ProductIds,omitempty"` // Unique identifiers of the products.
	ServiceIDs     []string                    `json:"ServiceIds"`           // Unique identifiers of the Services.
	UpdatedUTC     *configuration.TimeInterval `json:"UpdatedUtc,omitempty"` // Interval in which the products were updated.
	IncludeDefault bool                        `json:"IncludeDefault"`       // Whether or not to include default products for the service, i.e. products which are standard includes and not true extras. For example, the night's stay would be the default product for a room reservation. These may be useful for accounting purposes but should not be displayed to customers for selection. If ProductIds are provided, IncludeDefault defaults to true, otherwise it defaults to false.
	Limitation     json.Limitation             `json:"Limitation,omitempty"`
}

type Products []Product

type Product struct {
	ID                 string                      `json:"Id"`                     // Unique identifier of the product.
	ServiceID          string                      `json:"ServiceId"`              // Unique identifier of the Service.
	CategoryID         string                      `json:"CategoryId"`             // Unique identifier of the Product category.
	IsActive           bool                        `json:"IsActive"`               // Whether the product is still active.
	Name               string                      `json:"Name"`                   // Name of the product.
	ExternalName       string                      `json:"ExternalName"`           // Name of the product meant to be displayed to customer.
	ShortName          string                      `json:"ShortName"`              // Short name of the product.
	Description        string                      `json:"Description"`            // Description of the product.
	ExternalIdentifier string                      `json:"ExternalIdentifier"`     // Identifier of the product from external system.
	Charging           ProductCharging             `json:"Charging"`               // Charging of the product.
	Posting            ProductPosting              `json:"Posting"`                // Posting of the product.
	Promotions         services.Promotions         `json:"Promotions"`             // Promotions of the service.
	Classifications    ProductClassifications      `json:"ProductClassifications"` // Classifications of the service.
	Price              configuration.CurrencyValue `json:"Price"`                  // Price of the product.
}

type ProductCharging string

var (
	ProductChargingOnce                 ProductCharging = "Once"
	ProductChargingPerTimeUnit          ProductCharging = "PerTimeUnit"
	ProductChargingPerPersonPerTimeUnit ProductCharging = "PerPersonPerTimeUnit"
	ProductChargingPerPerson            ProductCharging = "PerPerson"
)

type ProductPosting string

var (
	ProductPostingOnce  ProductPosting = "Once"
	ProductPostingDaily ProductPosting = "Daily"
)

type ProductClassifications struct {
	Food     bool `json:"Food"`     // Product is classified as food.
	Beverage bool `json:"Beverage"` // Product is classified as beverage.
	Wellness bool `json:"Wellness"` // Product is classified as wellness.
	CityTax  bool `json:"CityTax"`  // Product is classified as city tax.
}

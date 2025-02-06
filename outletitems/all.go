package outletitems

import (
	"encoding/json"
	"time"

	"github.com/tim-online/go-errors"
	base "github.com/omniboost/go-mews/json"
	"github.com/omniboost/go-mews/omitempty"
)

const (
	endpointAll = "outletItems/getAll"

	Revenue     OutletItemType = "Revenue"
	NoneRevenue OutletItemType = "NonRevenue"
	Payment     OutletItemType = "Payment"
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

type AllResponse struct {
	OutletItems []OutletItem
	OutletBills []OutletBill
	Cursor      string `json:"Cursor"`
}

func (s *Service) NewAllRequest() *AllRequest {
	return &AllRequest{}
}

type AllRequest struct {
	base.BaseRequest
	Limitation base.Limitation       `json:"Limitation,omitempty"`
	StartUTC   *time.Time            `json:"StartUtc,omitempty"`
	EndUTC     *time.Time            `json:"EndUtc,omitempty"`
	TimeFilter OutletItemsTimeFilter `json:"TimeFilter,omitempty"`
}

func (r AllRequest) MarshalJSON() ([]byte, error) {
	return omitempty.MarshalJSON(r)
}

type OutletItemsTimeFilter string

const (
	TimeFilterClosed   OutletItemsTimeFilter = "Closed"
	TimeFilterConsumed OutletItemsTimeFilter = "Consumed"
)

func (f *OutletItemsTimeFilter) UnmarshalJSON(data []byte) error {
	var s string
	err := json.Unmarshal(data, &s)
	if err != nil {
		return err
	}

	switch s {
	case string(TimeFilterClosed):
		*f = TimeFilterClosed
		return nil
	case string(TimeFilterConsumed):
		*f = TimeFilterConsumed
		return nil
	}

	return errors.Errorf("Unknown outlet items time filter: %s", s)
}

// 	"AccountingCategoryId": "4ac8ce68-5732-4f1d-bf0d-e557072c926f",
// 	"Amount": {
// 		"Currency": "GBP",
// 		"Tax": 0.42,
// 		"TaxRate": 0.2,
// 		"Value": 2.5
// 	},
// 	"BillId": null,
// 	"ConsumptionUtc": "2016-07-27T12:48:39Z",
// 	"Id": "89b93f7c-5c63-4de2-bd17-ec5fee5e3120",
// 	"Name": "Caramel, Pepper & Chilli Popcorn",
// 	"Notes": null,
// 	"OrderId": "810b8c3a-d358-4378-84a9-534c830016fc",
// 	"ProductId": null,
// 	"Type": "ServiceRevenue"
// }

type OutletItems []OutletItem

type OutletItem struct {
	ID                   string         `json:"Id"`                   // Unique identifier of the item.
	BillID               string         `json:"BillId"`               // Unique identifier of the bill the item is assigned to.
	AccountingCategoryID string         `json:"AccountingCategoryId"` // Unique identifier of the Accounting Category the item belongs to.
	Type                 OutletItemType `json:"Type"`                 // Type of the item.
	Name                 string         `json:"Name"`                 // Name of the item.
	UnitCount            int            `json:"UnitCount"`            // Amount the item costs, negative amount represents either rebate or a payment.
	UnitCost             UnitCost       `json:"UnitCost"`             // Amount the item costs, negative amount represents either rebate or a payment.
	UnitAmount           UnitAmount     `json:"UnitAmount"`           // Unit amount of the item.
	CreatedUTC           time.Time      `json:"CreatedUtc"`           // Date and time of the item creation in UTC timezone in ISO 8601 format.
	ConsumptionUTC       time.Time      `json:"ConsumedUtc"`          // Date and time of the item consumption in UTC timezone in ISO 8601 format.
	ExternalIdentifier   string         `json:"ExternalIdentifier"`   // An identifier of this item from another system.
	Notes                string         `json:"Notes"`                // Additional notes.
}

type OutletItemType string

type OutletBill struct {
	ID        string    `json:"Id"`        // Unique identifier of the bill.
	OutletID  string    `json:"OutletId"`  // Unique identifier of the Outlet where the bill was issued.
	Number    string    `json:"Number"`    // Number of the bill.
	ClosedUTC time.Time `json:"ClosedUtc"` // Date and time of the bill closure in UTC timezone in ISO 8601 format.
	Notes     string    `json:"Notes"`
}

type UnitCost Amount

type UnitAmount Amount

type Amount struct {
	Currency   string    `json:"Currency"`   // ISO-4217 code of the Currency.
	NetValue   float64   `json:"NetValue"`   // Net value in case the item is taxed.
	GrossValue float64   `json:"GrossValue"` // Gross value including all taxes.
	TaxValues  TaxValues `json:"TaxValues"`  // The tax values applied.

	// Deprecated?
	Net     float64  `json:"Net"`     // Net value in case the item is taxed.
	Tax     float64  `json:"Tax"`     // Tax value in case the item is taxed.
	TaxRate *float64 `json:"TaxRate"` // Tax rate in case the item is taxed (e.g. 0.21).
	Value   float64  `json:"Value"`   // Amount in the currency (including tax if taxed).
}

type TaxValues []TaxValue

type TaxValue struct {
	Code  string  `json:"Code"`  // Code corresponding to tax type.
	Value float64 `json:"Value"` // Amount of tax applied.
}

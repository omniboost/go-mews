package accountingitems

import (
	"encoding/json"
	"time"

	"github.com/omniboost/go-mews/configuration"
	base "github.com/omniboost/go-mews/json"
	"github.com/omniboost/go-mews/omitempty"
	"github.com/tim-online/go-errors"
)

const (
	endpointAll = "accountingItems/getAll"

	ServiceRenue      AccountingItemType = "ServiceRevenue"
	ProductRevenue    AccountingItemType = "ProductRevenue"
	AdditionalRevenue AccountingItemType = "AdditionalRevenue"
	Payment           AccountingItemType = "Payment"
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
	AccountingItems        []AccountingItem
	OrderItems             OrderItems
	PaymentItems           PaymentItems
	CreditCardTransactions CreditCardTransactions
}

func (s *APIService) NewAllRequest() *AllRequest {
	return &AllRequest{
		Extent: AccountingItemsExtent{
			AccountingItems:        true,
			CreditCardTransactions: false,
		},
	}
}

type AllRequest struct {
	base.BaseRequest

	StartUTC   *time.Time                `json:"StartUtc,omitempty"`
	EndUTC     *time.Time                `json:"EndUtc,omitempty"`
	TimeFilter AccountingItemsTimeFilter `json:"TimeFilter,omitempty"`

	ConsumedUTC    configuration.TimeInterval `json:"ConsumedUtc,omitempty"`    // Interval in which the accounting item was consumed. Required if no other filter is provided.
	ClosedUTC      configuration.TimeInterval `json:"ClosedUtc,omitempty"`      // Interval in which the accounting item was closed. Required if no other filter is provided.
	UpdatedUTC     configuration.TimeInterval `json:"UpdatedUtc,omitempty"`     // Interval in which the accounting item was updated. Required if no other filter is provided.
	ItemIDs        []string                   `json:"ItemIds,omitempty"`        // Unique identifiers of the Accounting items. Required if no other filter is provided.
	RebatedItemIDs []string                   `json:"RebatedItemIds,omitempty"` // Unique identifiers of the Accounting items we are finding rebates for. Required if no other filter is provided.
	Currency       string                     `json:"Currency,omitempty"`       // ISO-4217 code of the Currency the item costs should be converted to.
	Extent         AccountingItemsExtent      `json:"Extent,omitempty"`         // Extent of data to be returned. E.g. it is possible to specify that together with the accounting items, credit card transactions should be also returned.
	States         []AccountingItemsState     `json:"States,omitempty"`         // States the accounting items should be in. If not specified, accounting items in Open or Closed states are returned.
}

func (r AllRequest) MarshalJSON() ([]byte, error) {
	return omitempty.MarshalJSON(r)
}

type AccountingItemsTimeFilter string

const (
	TimeFilterClosed   AccountingItemsTimeFilter = "Closed"
	TimeFilterConsumed AccountingItemsTimeFilter = "Consumed"
)

func (f *AccountingItemsTimeFilter) UnmarshalJSON(data []byte) error {
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

	return errors.Errorf("Unknown accounting items time filter: %s", s)
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

type AccountingItems []AccountingItem

type AccountingItem struct {
	ID                     string                `json:"Id"`                     // Unique identifier of the item.
	CustomerID             string                `json:"CustomerId"`             // Unique identifier of the Customer whose account the item belongs to.
	ProductID              string                `json:"ProductId"`              // Unique identifier of the Product.
	ServiceID              string                `json:"ServiceId"`              // Unique identifier of the Service the item belongs to.
	OrderID                string                `json:"OrderId"`                // Unique identifier of the order (or Reservation) the item belongs to.
	BillID                 string                `json:"BillId"`                 // Unique identifier of the bill the item is assigned to.
	CreditCardID           string                `json:"CreditCardId,omitempty"` // Unique identifier of the Credit card the item is associated to.
	InvoiceID              string                `json:"InvoiceId"`              // Unique identifier of the invoiced Bill the item is receivable for.
	AccountingCategoryID   string                `json:"AccountingCategoryId"`   // Unique identifier of the Accounting Category the item belongs to.
	Amount                 Amount                `json:"Amount"`                 // Amount the item costs, negative amount represents either rebate or a payment.
	Type                   AccountingItemType    `json:"Type"`                   // Type of the item.
	Name                   string                `json:"Name"`                   // Name of the item.
	Notes                  string                `json:"Notes"`                  // Additional notes.
	ConsumptionUTC         time.Time             `json:"ConsumptionUtc"`         // Date and time of the item consumption in UTC timezone in ISO 8601 format.
	ClosedUTC              time.Time             `json:"ClosedUtc"`              // Date and time of the item bill closure in UTC timezone in ISO 8601 format.
	SubType                AccountingItemSubtype `json:"SubType"`                // subtype of the item. Note that the subtype depends on the Type of the item.
	State                  string                `json:"State"`
	RebatedItemID          string                `json:"RebatedItemId"` // Unique identifier of Order item which has been rebated by current item.
	TaxExemptionReasonCode string                `json:"TaxExemptionReasonCode,omitempty"`
}

type OrderItems []OrderItem

type OrderItem struct {
	ID                   string          `json:"Id"`                   // Unique identifier of the item.
	AccountID            string          `json:"AccountID"`            // Unique identifier of the account (for example Customer) the item belongs to.
	OrderID              string          `json:"OrderId"`              // Unique identifier of the order (or Reservation) the item belongs to.
	BillID               string          `json:"BillId"`               // Unique identifier of the bill the item is assigned to.
	AccountingCategoryID string          `json:"AccountingCategoryId"` // Unique identifier of the Accounting Category the item belongs to.
	UnitCount            int             `json:"UnitCount"`            // Unit count of item, i.e. the number of sub-items or units, if applicable.
	UnitAmount           Amount          `json:"UnitAmount"`           // Unit amount of item, i.e. the amount of each individual sub-item or unit, if applicable.
	Amount               Amount          `json:"Amount"`               // Amount the item costs, negative amount represents either rebate or a payment.
	RevenueType          RevenueType     `json:"RevenueType"`          // Revenue type of the item.
	ConsumedUTC          time.Time       `json:"ConsumedUtc"`          // Date and time of the item consumption in UTC timezone in ISO 8601 format.
	ClosedUTC            time.Time       `json:"ClosedUtc"`            // Date and time of the item bill closure in UTC timezone in ISO 8601 format.
	AccountingState      AccountingState `json:"AccountingState"`      // Accounting state of the item.
	Data                 OrderItemData   `json:"Data"`                 // Additional data specific to particular order item.
}

type PaymentItems []PaymentItem

type PaymentItem struct {
	ID                   string           `json:"Id"`                   // Unique identifier of the item.
	AccountID            string           `json:"AccountId"`            // Unique identifier the account (for example Customer) the item belongs to
	BillID               string           `json:"BillId"`               // Unique identifier of the the Bill the item is assigned to.
	AccountingCategoryID string           `json:"AccountingCategoryId"` // Unique identifier of the Accounting Category the item belongs to
	Amount               Amount           `json:"Amount"`               // Item's amount, negative amount represents either rebate or a payment.
	OriginalAmount       Amount           `json:"OriginalAmount"`       // Amount of item; note a negative amount represents a rebate or payment. Contains the earliest known value in conversion chain.
	Notes                string           `json:"Notes"`                // Additional notes.
	SettlementID         string           `json:"SettlementId"`         // Identifier of the settled payment from the external system (ApplePay/GooglePay).
	ConsumedUTC          time.Time        `json:"ConsumedUtc"`          // Date and time of the item consumption in UTC timezone in ISO 8601 format.
	ClosedUTC            time.Time        `json:"ClosedUtc"`            // Date and time of the item bill closure in UTC timezone in ISO 8601 format.
	AccountingState      AccountingState  `json:"AccountingState"`      // Accounting state of the item.
	State                PaymentItemState `json:"State"`                // Payment state of the item.
	Data                 PaymentItemData  `json:"Data"`                 // Additional data specific to particular payment item.
}

type PaymentItemData struct {
	Discriminator PaymentItemDataDiscriminator `json:"Discriminator"` // Type of the payment item (e.g. CreditCard).
	Value         map[string]interface{}       `json:""`              // Based on order item discriminator, e.g. Credit card payment item data or null for types without any additional data
}

type PaymentItemDataDiscriminator string

type PaymentItemState string

type CreditCardTransactions []CreditCardTransaction

type CreditCardTransaction struct {
	ID            string    `json:"Id"`
	PaymentID     string    `json:"PaymentId"`     // Unique identifier of the
	SettlementID  string    `json:"SettlementId"`  // Identifier of the settlement.
	SettledUTC    time.Time `json:"SettledUtc"`    // Settlement date and time in UTC timezone in ISO 8601 format.
	Fee           Amount    `json:"Fee"`           // Transaction fee - this includes an estimate of bank charges.
	AdjustedFee   Amount    `json:"AdjustedFee"`   // Transaction fee (adjusted) - this is the final confirmed transaction fee, including confirmed bank charges.
	ChargedAmount Amount    `json:"ChargedAmount"` // Charged amount of the transaction.
	SettledAmount Amount    `json:"SettledAmount"` // Settled amount of the transaction.
}

type AccountingItemType string

type AccountingItemsExtent struct {
	AccountingItems        bool `json:"AccountingItems"`
	OrderItems             bool `json:"OrderItems"`
	PaymentItems           bool `json:"PaymentItems"`
	CreditCardTransactions bool `json:"CreditCardTransactions"`
}

type AccountingItemsState string

const (
	AccountingItemsStateClosed   AccountingItemsState = "Closed"
	AccountingItemsStateOpen     AccountingItemsState = "Open"
	AccountingItemsStateInactive AccountingItemsState = "Inactive"
	AccountingItemsStateCanceled AccountingItemsState = "Canceled"
)

type Cost struct {
	Currency string   `json:"Currency"` // ISO-4217 code of the Currency.
	Net      float64  `json:"Net"`      // Net value in case the item is taxed.
	Tax      float64  `json:"Tax"`      // Tax value in case the item is taxed.
	TaxRate  *float64 `json:"TaxRate"`  // Tax rate in case the item is taxed (e.g. 0.21).
	Value    float64  `json:"Value"`    // Amount in the currency (including tax if taxed).
}

type AccountingItemSubtype string

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

type AccountingState string

type OrderItemData struct {
	Discriminator        OrderItemDataDiscriminator `json:"Discriminator"` // Type of the order item (e.g. ProductOrder).
	Value                json.RawMessage            `json:"Value"`         // Based on order item discriminator, e.g. Product order item data or null for types without any additional data.
	ProductOrderItemData ProductOrderItemData       `json:"-"`
	RebateOrderItemData  RebateOrderItemData        `json:"-"`
}

type OrderItemDataDiscriminator string

func (d *OrderItemData) UnmarshalJSON(data []byte) error {
	type alias OrderItemData
	a := alias(*d)
	err := json.Unmarshal(data, &a)
	if err != nil {
		return err
	}

	err = json.Unmarshal(a.Value, &a.ProductOrderItemData)
	if err != nil {
		return err
	}

	err = json.Unmarshal(a.Value, &a.RebateOrderItemData)
	if err != nil {
		return err
	}

	*d = OrderItemData(a)
	return nil
}

type RebateOrderItemData struct {
	RebatedItemID string `json:"RebatedItemId"` // Unique identifier of Order item which has been rebated by current item.
}

type ProductOrderItemData struct {
	ProductID     string `json:"ProductId"`     // Unique identifier of the Product
	AgeCategoryID string `json:"AgeCategoryId"` // Unique identifier of the Age Category
}

type RevenueType string

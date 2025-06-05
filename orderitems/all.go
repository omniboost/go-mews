package orderitems

import (
	"time"

	"github.com/cydev/zero"
	"github.com/omniboost/go-mews/configuration"
	base "github.com/omniboost/go-mews/json"
	"github.com/omniboost/go-mews/omitempty"
)

const (
	endpointAll = "orderItems/getAll"
)

// List all orderitems
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
	OrderItems OrderItems
	Cursor     string `json:"Cursor"`
}

func (s *Service) NewAllRequest() *AllRequest {
	return &AllRequest{}
}

type AllRequest struct {
	base.BaseRequest
	EnterpriseIDs    []string                   `json:"EnterpriseIDs,omitempty"`
	OrderItemIDs     []string                   `json:"OrderItemIds,omitempty"`
	ServiceOrderIDs  []string                   `json:"ServiceOrderIds,omitempty"`
	ServiceIDs       []string                   `json:"ServiceIds,omitempty"`
	BillIDs          []string                   `json:"BillIDs,omitempty"`
	CreatedUTC       configuration.TimeInterval `json:"CreatedUtc,omitempty"`
	UpdatedUTC       configuration.TimeInterval `json:"UpdatedUtc,omitempty"`
	ConsumedUTC      configuration.TimeInterval `json:"ConsumedUtc,omitempty"`
	CanceledUTC      configuration.TimeInterval `json:"CanceledUtc,omitempty"`
	ClosedUTC        configuration.TimeInterval `json:"ClosedUtc,omitempty"`
	AccountingStates []AccountingState          `json:"AccountingStates,omitempty"`
	States           []OrderItemState           `json:"States,omitempty"`
	Types            OrderItemType              `json:"Types,omitempty"`
	Currency         string                     `json:"Currency,omitempty"`
	Limitation       base.Limitation            `json:"Limitation,omitempty"`
}

func (r AllRequest) MarshalJSON() ([]byte, error) {
	return omitempty.MarshalJSON(r)
}

type OrderItems []OrderItem

type OrderItem struct {
	ID                         string           `json:"ID,omitempty"`
	EnterpriseID               string           `json:"EnterpriseID,omitempty"`
	AccountID                  string           `json:"AccountID,omitempty"`
	AccountType                string           `json:"AccountType,omitempty"`
	ServiceID                  string           `json:"ServiceID,omitempty"`
	ServiceOrderID             string           `json:"ServiceOrderID,omitempty"`
	BillID                     string           `json:"BillID,omitempty"`
	AccountingCategoryID       string           `json:"AccountingCategoryID,omitempty"`
	BillingName                string           `json:"BillingName,omitempty"`
	UnitCount                  int              `json:"UnitCount,omitempty"`
	UnitAmount                 Amount           `json:"UnitAmount,omitempty"`
	Amount                     Amount           `json:"Amount,omitempty"`
	OriginalAmount             Amount           `json:"OriginalAmount,omitempty"`
	Notes                      string           `json:"Notes,omitempty"`
	RevenueType                RevenueType      `json:"RevenueType,omitempty"`
	CreatorProfileID           string           `json:"CreatorProfileID,omitempty"`
	UpdatedProfileID           string           `json:"UpdatedProfileID,omitempty"`
	ConsumedUTC                time.Time        `json:"ConsumedUtc,omitempty"`
	ClosedUTC                  time.Time        `json:"ClosedUtc,omitempty"`
	ChargedUTC                 time.Time        `json:"ChargedUtc,omitempty"`
	CreatedUTC                 time.Time        `json:"CreatedUtc,omitempty"`
	UpdatedUTC                 time.Time        `json:"UpdatedUtc,omitempty"`
	StartUTC                   time.Time        `json:"StartUtc,omitempty"`
	AccountingState            AccountingState  `json:"AccountingState,omitempty"`
	Type                       OrderItemType    `json:"Type,omitempty"`
	Options                    OrderItemOptions `json:"Options,omitempty"`
	Data                       OrderItemData    `json:"Data,omitempty"`
	TaxExemptionReason         string           `json:"TaxExemptionReason,omitempty"`
	TaxExemptionLegalReference string           `json:"TaxExemptionLegalReference,omitempty"`
}

type RevenueType string

var (
	RevenueTypeService    RevenueType = "Service"
	RevenueTypeProduct    RevenueType = "Product"
	RevenueTypeAdditional RevenueType = "Additional"
)

type OrderItemOptions struct {
	CanceledWithReservation bool `json:"CanceledWithReservation,omitempty"`
}

type AccountingState string

var (
	AccountingStateOpen     AccountingState = "Open"
	AccountingStateClosed   AccountingState = "Closed"
	AccountingStateInactive AccountingState = "Inactive"
	AccountingStateCanceled AccountingState = "Canceled"
)

type OrderItemState string

var (
	OrderItemStateCharged   OrderItemState = "Charged"
	OrderItemStateCanceled  OrderItemState = "Canceled"
	OrderItemStatePending   OrderItemState = "Pending"
	OrderItemStateFailed    OrderItemState = "Failed"
	OrderItemStateVerifying OrderItemState = "Verifying"
)

type Amount struct {
	Currency   string    `json:"Currency"`   // ISO-4217 code of the Currency.
	NetValue   float64   `json:"NetValue"`   // Net value in case the item is taxed.
	GrossValue float64   `json:"GrossValue"` // Gross value including all taxes.
	TaxValues  TaxValues `json:"TaxValues"`  // The tax values applied.
	TaxCodes   []string  `json:"TaxCodes"`   // Codes of Tax rates to be applied to the item. (Note, you can only define one tax when sending GrossValue. For multiple taxes, use NetValue)

	// Deprecated?
	Net     float64  `json:"Net"`     // Net value in case the item is taxed.
	Tax     float64  `json:"Tax"`     // Tax value in case the item is taxed.
	TaxRate *float64 `json:"TaxRate"` // Tax rate in case the item is taxed (e.g. 0.21).
	Value   float64  `json:"Value"`   // Amount in the currency (including tax if taxed).
}

func (a Amount) MarshalJSON() ([]byte, error) {
	return omitempty.MarshalJSON(a)
}

func (a Amount) IsEmpty() bool {
	return zero.IsZero(a)
}

type TaxValues []TaxValue

type TaxValue struct {
	Code  string  `json:"Code"`  // Code corresponding to tax type.
	Value float64 `json:"Value"` // Amount of tax applied.
}

type OrderItemData struct {
	Discriminator OrderItemDataDiscriminator `json:"Discriminator,omitempty"`
	Rebate        RebateData                 `json:"Rebate,omitempty"`
	Product       ProductData                `json:"Product,omitempty"`
}

type OrderItemDataDiscriminator string

var (
	OrderItemDataDiscriminatorRebate  OrderItemDataDiscriminator = "Rebate"
	OrderItemDataDiscriminatorProduct OrderItemDataDiscriminator = "Product"
)

type RebateData struct {
	RebatedItemID string `json:"RebatedItemId,omitempty"`
}

type ProductData struct {
	ProductID     string `json:"ProductId,omitempty"`
	AgeCategoryID string `json:"AgeCategoryId,omitempty"`
}

type OrderItemType string

var (
	OrderItemTypeCancellationFee         OrderItemType = "CancellationFee"
	OrderItemTypeNightRebate             OrderItemType = "NightRebate"
	OrderItemTypeProductOrderRebate      OrderItemType = "ProductOrderRebate"
	OrderItemTypeAdditionalExpenseRebate OrderItemType = "AdditionalExpenseRebate"
	OrderItemTypeDeposit                 OrderItemType = "Deposit"
	OrderItemTypeExchangeRateDifference  OrderItemType = "ExchangeRateDifference"
	OrderItemTypeCustomItem              OrderItemType = "CustomItem"
	OrderItemTypeServiceCharge           OrderItemType = "ServiceCharge"
	OrderItemTypeCityTax                 OrderItemType = "CityTax"
	OrderItemTypeCityTaxDiscount         OrderItemType = "CityTaxDiscount"
	OrderItemTypeSpaceOrder              OrderItemType = "SpaceOrder"
	OrderItemTypeProductOrder            OrderItemType = "ProductOrder"
	OrderItemTypeSurcharge               OrderItemType = "Surcharge"
	OrderItemTypeTaxCorrection           OrderItemType = "TaxCorrection"
)

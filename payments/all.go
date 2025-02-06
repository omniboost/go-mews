package payments

import (
	"time"

	"github.com/omniboost/go-mews/configuration"
	base "github.com/omniboost/go-mews/json"
	"github.com/omniboost/go-mews/omitempty"
)

const (
	endpointAll = "payments/getAll"
)

// List all Payments
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
	Payments Payments
	Cursor   string `json:"Cursor"`
}

func (s *Service) NewAllRequest() *AllRequest {
	return &AllRequest{}
}

type AllRequest struct {
	base.BaseRequest
	EnterpriseIDs    []string                   `json:"EnterpriseIDs,omitempty"`
	PaymentIDs       []string                   `json:"PaymentIDs,omitempty"`
	BillIDs          []string                   `json:"BillIDs,omitempty"`
	CreatedUTC       configuration.TimeInterval `json:"CreatedUtc,omitempty"`
	UpdatedUTC       configuration.TimeInterval `json:"UpdatedUtc,omitempty"`
	ChargedUTC       configuration.TimeInterval `json:"ChargedUtc,omitempty"`
	ClosedUTC        configuration.TimeInterval `json:"ClosedUtc,omitempty"`
	SettlementUTC    configuration.TimeInterval `json:"SettlementUtc,omitempty"`
	AccountingStates []AccountingState          `json:"AccountingStates,omitempty"`
	States           []PaymentState             `json:"States,omitempty"`
	Type             PaymentType                `json:"Type,omitempty"`
	Currency         string                     `json:"Currency,omitempty"`
	Limitation       base.Limitation            `json:"Limitation,omitempty"`
}

func (r AllRequest) MarshalJSON() ([]byte, error) {
	return omitempty.MarshalJSON(r)
}

type Payments []Payment

type Payment struct {
	ID                   string          `json:"ID,omitempty"`
	EnterpriseID         string          `json:"EnterpriseID,omitempty"`
	AccountID            string          `json:"AccountID,omitempty"`
	AccountType          string          `json:"AccountType,omitempty"`
	BillID               string          `json:"BillID,omitempty"`
	AccountingCategoryID string          `json:"AccountingCategoryID,omitempty"`
	Amount               Amount          `json:"Amount,omitempty"`
	OriginalAmount       Amount          `json:"OriginalAmount,omitempty"`
	Notes                string          `json:"Notes,omitempty"`
	SettlementID         string          `json:"SettlementID,omitempty"`
	ConsumedUTC          time.Time       `json:"ConsumedUtc,omitempty"`
	ClosedUTC            time.Time       `json:"ClosedUtc,omitempty"`
	ChargedUTC           time.Time       `json:"ChargedUtc,omitempty"`
	CreatedUTC           time.Time       `json:"CreatedUtc,omitempty"`
	UpdatedUTC           time.Time       `json:"UpdatedUtc,omitempty"`
	SettlementUTC        time.Time       `json:"SettlementUtc,omitempty"`
	AccountingState      AccountingState `json:"AccountingState,omitempty"`
	State                PaymentState    `json:"State,omitempty"`
	Identifier           string          `json:"Identifier,omitempty"`
	Type                 PaymentType     `json:"Type,omitempty"`
	Kind                 PaymentKind     `json:"Kind,omitempty"`
	Data                 PaymentData     `json:"Data,omitempty"`
}

type AccountingState string

var (
	AccountingStateOpen     AccountingState = "Open"
	AccountingStateClosed   AccountingState = "Closed"
	AccountingStateInactive AccountingState = "Inactive"
	AccountingStateCanceled AccountingState = "Canceled"
)

type PaymentState string

var (
	PaymentStateCharged   PaymentState = "Charged"
	PaymentStateCanceled  PaymentState = "Canceled"
	PaymentStatePending   PaymentState = "Pending"
	PaymentStateFailed    PaymentState = "Failed"
	PaymentStateVerifying PaymentState = "Verifying"
)

type PaymentType string

var (
	PaymentTypePayment            PaymentType = "Payment"
	PaymentTypeCreditCardPayment  PaymentType = "CreditCardPayment"
	PaymentTypeAlternativePayment PaymentType = "AlternativePayment"
	PaymentTypeCashPayment        PaymentType = "CashPayment"
	PaymentTypeInvoicePayment     PaymentType = "InvoicePayment"
	PaymentTypeExternalPayment    PaymentType = "ExternalPayment"
)

type PaymentKind string

var (
	PaymentKindPayment            PaymentKind = "Payment"
	PaymentKindChargeback         PaymentKind = "Chargeback"
	PaymentKindChargebackReversal PaymentKind = "ChargebackReversal"
	PaymentKindRefund             PaymentKind = "Refund"
)

type PaymentData struct {
	Discriminator string              `json:"Discriminator,omitempty"`
	CreditCard    CreditCard          `json:"CreditCard,omitempty"`
	Invoice       InvoicePayment      `json:"Invoice,omitempty"`
	External      ExternalPaymentData `json:"External,omitempty"`
}

type CreditCard struct {
	CreditCardID string          `json:"CreditCardID,omitempty"`
	Transaction  TransactionData `json:"Transaction,omitempty"`
}

type TransactionData struct {
	PaymentID     string `json:"PaymentID,omitempty"`
	SettlementID  string `json:"SettlementID,omitempty"`
	SettledUTC    string `json:"SettledUtc,omitempty"`
	Fee           Amount `json:"Fee,omitempty"`
	AdjustedFee   Amount `json:"AdjustedFee,omitempty"`
	ChargedAmount Amount `json:"ChargedAmount,omitempty"`
	SettledAmount Amount `json:"SettledAmount,omitempty"`
}

type InvoicePayment struct {
	InvoiceID string `json:"InvoiceID,omitempty"`
}

type ExternalPaymentData struct {
	Type               ExternalPaymentType `json:"Type,omitempty"`
	ExternalIdentifier string              `json:"ExternalIdentifier,omitempty"`
}

type ExternalPaymentType string

var (
	Unspecified                ExternalPaymentType = "Unspecified"
	BadDebts                   ExternalPaymentType = "BadDebts"
	Bacs                       ExternalPaymentType = "Bacs"
	WireTransfer               ExternalPaymentType = "WireTransfer"
	Invoice                    ExternalPaymentType = "Invoice"
	ExchangeRateDifference     ExternalPaymentType = "ExchangeRateDifference"
	Complimentary              ExternalPaymentType = "Complimentary"
	Reseller                   ExternalPaymentType = "Reseller"
	ExchangeRoundingDifference ExternalPaymentType = "ExchangeRoundingDifference"
	Barter                     ExternalPaymentType = "Barter"
	Commission                 ExternalPaymentType = "Commission"
	BankCharges                ExternalPaymentType = "BankCharges"
	CrossSettlement            ExternalPaymentType = "CrossSettlement"
	Cash                       ExternalPaymentType = "Cash"
	Prepayment                 ExternalPaymentType = "Prepayment"
	Cheque                     ExternalPaymentType = "Cheque"
	Bancontact                 ExternalPaymentType = "Bancontact"
	IDeal                      ExternalPaymentType = "IDeal"
	PayPal                     ExternalPaymentType = "PayPal"
	GiftCard                   ExternalPaymentType = "GiftCard"
	LoyaltyPoints              ExternalPaymentType = "LoyaltyPoints"
	ChequeVacances             ExternalPaymentType = "ChequeVacances"
	OnlinePayment              ExternalPaymentType = "OnlinePayment"
	CardCheck                  ExternalPaymentType = "CardCheck"
	PaymentHubRedirection      ExternalPaymentType = "PaymentHubRedirection"
	Voucher                    ExternalPaymentType = "Voucher"
	MasterCard                 ExternalPaymentType = "MasterCard"
	Visa                       ExternalPaymentType = "Visa"
	Amex                       ExternalPaymentType = "Amex"
	Discover                   ExternalPaymentType = "Discover"
	DinersClub                 ExternalPaymentType = "DinersClub"
	Jcb                        ExternalPaymentType = "Jcb"
	UnionPay                   ExternalPaymentType = "UnionPay"
	Twint                      ExternalPaymentType = "Twint"
	Reka                       ExternalPaymentType = "Reka"
	LoyaltyCard                ExternalPaymentType = "LoyaltyCard"
	PosDiningAndSpaReward      ExternalPaymentType = "PosDiningAndSpaReward"
	DepositCheck               ExternalPaymentType = "DepositCheck"
	DepositCash                ExternalPaymentType = "DepositCash"
	DepositCreditCard          ExternalPaymentType = "DepositCreditCard"
	DepositWireTransfer        ExternalPaymentType = "DepositWireTransfer"
)

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

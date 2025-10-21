package configuration

import (
	"time"

	base "github.com/omniboost/go-mews/json"
	"github.com/omniboost/go-mews/services"
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
	NowUtc                           time.Time          `json:"NowUtc"`                           // Current server date and time in UTC timezone in ISO 8601 format.
	Enterprise                       Enterprise         `json:"Enterprise"`                       // The enterprise (e.g. hotel, hostel) associated with the access token.
	Service                          services.Service   `json:"Service"`                          // The reservable service (e.g. accommodation, parking) associated with the access token of the service scoped integration.
	PaymentCardStorage               PaymentCardStorage `json:"PaymentCardStorage"`               // Contains information about payment card storage.
	IsIdentityDocumentNumberRequired bool               `json:"IsIdentityDocumentNumberRequired"` // Whether the identity documents for this enterprise include the value of identity document number as required by the legal environment. When false, the number is not required, and an empty string can be used in write operations. In read operations, an empty string is returned when an empty string was provided for the number.
}

type Enterprise struct {
	ID                                 string                  `json:"Id"`                                 // Unique identifier of the enterprise.
	ExternalIdentifier                 string                  `json:"ExternalIdentifier"`                 // Identifier of the enterprise from external system.
	HoldingKey                         string                  `json:"HoldingKey"`                         // Identifies an enterprise in the external system of a holding company. The holding company may administer multiple portfolios.
	ChainID                            string                  `json:"ChainId"`                            // Unique identifier of the chain to which the enterprise belongs.
	ChainName                          string                  `json:"ChainName"`                          // Name of the Chain to which the enterprise belongs.
	CreatedUTC                         time.Time               `json:"CreatedUtc"`                         // Creation date and time of the enterprise in UTC timezone in ISO 8601 format.
	UpdatedUTC                         time.Time               `json:"UpdatedUtc"`                         // Update date and time of the enterprise in UTC timezone in ISO 8601 format.
	Name                               string                  `json:"Name"`                               // Name of the enterprise.
	TimeZoneIdentifier                 string                  `json:"TimeZoneIdentifier"`                 // IANA timezone identifier of the enterprise.
	LegalEnvironmentCode               string                  `json:"LegalEnvironmentCode"`               // Unique identifier of the legal environment where the enterprise resides.
	AccommodationEnvironmentCode       string                  `json:"AccommodationEnvironmentCode"`       // Unique code of the accommodation environment where the enterprise resides.
	AccountingEnvironmentCode          string                  `json:"AccountingEnvironmentCode"`          // Unique code of the accounting environment where the enterprise resides.
	TaxEnvironmentCode                 string                  `json:"TaxEnvironmentCode"`                 // Unique code of the tax environment where the enterprise resides.
	DefaultLanguageCode                string                  `json:"DefaultLanguageCode"`                // Language-culture codes of the enterprise default Language.
	AccountingEditableHistoryInterval  base.Duration           `json:"AccountingEditableHistoryInterval"`  // Editable history interval for accounting data in ISO 8601 duration format.
	OperationalEditableHistoryInterval base.Duration           `json:"OperationalEditableHistoryInterval"` // Editable history interval for operational data in ISO 8601 duration format.
	BusinessDayClosingOffset           base.Duration           `json:"BusinessDayClosingOffset"`           // The offset value for the business day closing time, in ISO 8601 duration format.
	WebsiteURL                         string                  `json:"WebsiteUrl"`                         // URL of the enterprise website.
	Email                              string                  `json:"Email"`                              // Email address of the enterprise.
	Phone                              string                  `json:"Phone"`                              // Phone number of the enterprise.
	LogoImageID                        string                  `json:"LogoImageId"`                        // Unique identifier of the Image of the enterprise logo.
	CoverImageID                       string                  `json:"CoverImageId"`                       // Unique identifier of the Image of the enterprise cover.
	Pricing                            Pricing                 `json:"Pricing"`                            // Pricing of the enterprise.
	TaxPrecision                       int                     `json:"TaxPrecision"`                       // Tax precision used for financial calculations in the enterprise. If null, Currency precision is used.
	AddressID                          string                  `json:"AddressId"`                          // Unique identifier of the address of the enterprise.
	Address                            Address                 `json:"Address"`                            // Address of the enterprise.
	GroupNames                         []string                `json:"GroupNames"`                         // A list of the group names of the enterprise.
	Subscription                       EnterpriseSubscription  `json:"Subscription"`                       // Subscription information of the enterprise.
	Currencies                         Currencies              `json:"Currencies"`                         // Currencies accepted by the enterprise.
	AccountingConfiguration            AccountingConfiguration `json:"AccountingConfiguration"`            // Configuration information containing financial information about the property.
	IsPortfolio                        bool                    `json:"IsPortfolio"`                        // Whether the enterprise is a Portfolio enterprise (see Multi-property guidelines).

	// Deprecated
	EditableHistoryInterval base.Duration `json:"EditableHistoryInterval"` // Editable history interval in ISO 8601 duration format.
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

type PaymentCardStorage struct {
	PublicKey string `json:"PublicKey"` // Key for accessing PCI proxy storage.
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

// Gross - The enterprise shows amount with gross prices.
// Net - The enterprise shows amount with net prices.
type Pricing string

type EnterpriseSubscription struct {
	TaxIdentifier string `json:"TaxIdentifier"` // Tax identifier of the Enterprise.
}

type AccountingConfiguration struct {
	AdditionalTaxIdentifier     string                           `json:"AdditionalTaxIdentifier"` // Organization number.
	CompanyName                 string                           `json:"CompanyName"`             // Legal name of the company.
	BankAccountNumber           string                           `json:"BankAccountNumber"`       // Bank account number.
	BankName                    string                           `json:"BankName"`                // Name of the bank.
	IBAN                        string                           `json:"Iban"`                    // International Bank Account Number.
	BIC                         string                           `json:"Bic"`                     // Business Identification Code.
	SurchargeConfiguration      SurchargingFeesConfiguration     `json:"SurchargeConfiguration"`  // Configuration for surcharging fees.
	EnabledExternalPaymentTypes []ExternalPaymentType            `json:"EnabledExternalPaymentTypes"`
	Options                     []AccountingConfigurationOptions `json:"Options"` // Accounting configuration options.
}

type SurchargingFeesConfiguration struct {
	SurchargeFees      map[string]float64 `json:"SurchargeFees"`      // Dictionary keys are CreditCardType and values are surcharging fees as a percentage.
	SurchargeServiceID string             `json:"SurchargeServiceId"` // Unique identifier of the surcharging Service.
	SurchargeTaxCode   string             `json:"SurchargeTaxCode"`   // Surcharging fee TaxCode.
}

// Unspecified - Unspecified (unavailable in French Legal Environment)
// BadDebts - Bad debts
// Bacs - Bacs payment
// WireTransfer - Wire transfer
// Invoice - Invoice
// ExchangeRateDifference - Exchange rate difference
// Complimentary - Complimentary
// Reseller - Reseller
// ExchangeRoundingDifference - Exchange rounding difference
// Barter - Barter
// Commission - Commission
// BankCharges - Bank charges
// CrossSettlement - Cross settlement
// Cash - Cash
// CreditCard - Credit card – deprecated, only for existing partners
// Prepayment - Prepayment
// Cheque - Cheque
// Bancontact - Bancontact
// IDeal - iDeal – deprecated, only for existing partners
// PayPal - PayPal – deprecated, only for existing partners
// GiftCard - Gift card
// LoyaltyPoints - Loyalty points
// ChequeVacances - Chèque-Vacances
// OnlinePayment - Online payment – deprecated, only for existing partners
// CardCheck - Card check
// PaymentHubRedirection - Payment hub redirection
// Voucher - Voucher
// MasterCard - MasterCard – deprecated, only for existing partners
// Visa - Visa – deprecated, only for existing partners
// Amex - American Express – deprecated, only for existing partners
// Discover - Discover – deprecated, only for existing partners
// DinersClub - Diners Club – deprecated, only for existing partners
// Jcb - JCB – deprecated, only for existing partners
// UnionPay - UnionPay – deprecated, only for existing partners
// Twint - TWINT
// Reka - Reka
// LoyaltyCard - Loyalty card
// PosDiningAndSpaReward - POS Dining & Spa Reward
// DirectDebit - Direct debit
// DepositCheck - Deposit - check
// DepositCash - Deposit - cash
// DepositCreditCard - Deposit - credit card
// DepositWireTransfer - Deposit - wire transfer
type ExternalPaymentType string

// OptionalCreditCardPaymentDetails - Optional credit card payment details
// ReceivableTrackingEnabled - Receivable tracking enabled
// SeparateDepositsOnBill - Separate deposits on bill
// AllowModifyingClosedBills - Allow modifying closed bills
// RequireAccountingCategorySetup - Require accounting category setup
// GroupTaxesOnBill - Group taxes on bill
// DisplayEmployeeNameOnBill - Display employee name on bill
// TaxDeclarationOnDeposit - Tax declaration on deposit
type AccountingConfigurationOptions string

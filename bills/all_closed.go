package bills

import (
	"time"

	"encoding/json"

	"github.com/omniboost/go-mews/accountingitems"
	base "github.com/omniboost/go-mews/json"
)

const (
	endpointAllClosed = "bills/getAllClosed"
)

// List all products
func (s *Service) AllClosed(requestBody *AllRequest) (*AllResponse, error) {
	// @TODO: create wrapper?
	if err := s.Client.CheckTokens(); err != nil {
		return nil, err
	}

	apiURL, err := s.Client.GetApiURL(endpointAllClosed)
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

func (s *Service) NewAllClosedRequest() *AllRequest {
	return &AllRequest{}
}

type AllClosedRequest struct {
	base.BaseRequest
	StartUTC *time.Time `json:"StartUtc,omitempty"`
	EndUTC   *time.Time `json:"EndUtc,omitempty"`
}

type AllClosedResponse struct {
	Bills Bills `json:"Bills"` // The closed bills.
}

type Bills []Bill

type Bill struct {
	ID                    string                     `json:"Id"`                    // Unique identifier of the bill.
	CustomerID            string                     `json:"CustomerId"`            // Unique identifier of the Customer the bill is issued to.
	CompanyID             string                     `json:"CompanyId"`             // Unique identifier of the Company the bill is issued to.
	CounterID             string                     `json:"CounterId"`             // Unique identifier of the bill Counter.
	State                 BillState                  `json:"State"`                 // State of the bill.
	Type                  BillType                   `json:"Type"`                  // Type of the bill.
	Number                string                     `json:"Number"`                // Number of the bill.
	VariableSymbol        string                     `json:"VariableSymbol"`        // Unique identifier of the bill.
	IssuedUTC             time.Time                  `json:"IssuedUtc"`             // Date and time of the bill issuance in UTC timezone in ISO 8601 format.
	TaxedUTC              base.Date                  `json:"TaxedUtc"`              // Taxation date of the bill in UTC timezone in ISO 8601 format.
	DueUTC                *time.Time                 `json:"DueUtc"`                // Bill due date and time in UTC timezone in ISO 8601 format.
	Notes                 string                     `json:"Notes"`                 // Additional notes.
	OrderItems            accountingitems.OrderItems `json:"OrderItems"`            // The revenue items on the bill.
	Revenue               Revenue                    `json:"Revenue"`               // The revenue items on the bill.
	Payments              Payments                   `json:"Payments"`              // The payments on the bill.
	AssociatedAccountData []AssociatedAccountData    `json:"AssociatedAccountData"` // Account data of the associated account on a bill. Currently one object is supported and only populated when the bill is closed.
	OwnerData             BillOwnerData              `json:"OwnerData"`             // Additional information about owner of the bill. Can be a Customer or Company. Persisted at the time of closing of the bill.
	CompanyDetails        BillCompanyData            `json:"CompanyDetails"`        // Additional information about the company assigned to the bill. Not the same as the owner. Persisted at the time of closing of the bill.
	EnterpriseData        BillEnterpriseData         `json:"EnterpriseData"`        // Additional information about the enterprise issuing the bill, including bank account details. Persisted at the time of closing of the bill.
	PurchaseOrderNumber   string                     `json:"PurchaseOrderNumber"`   // Unique number of the purchase order from the buyer.
	Options               BillOptions                `json:"Options"`               // Options of the bill.
}

type BillType string

const (
	BillTypeReceipt BillType = "Receipt"
	BillTypeInvoice BillType = "Invoice"
)

type BillState string

const (
	BillStateOpen   BillState = "open"
	BillStateClosed BillState = "closed"
)

type Revenue []accountingitems.AccountingItem
type Payments []accountingitems.AccountingItem

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

type AssociatedAccountData struct {
	Discriminator    string            `json:"Discriminator"`              // Determines type of value.
	BillCustomerData *BillCustomerData `json:"BillCustomerData,omitempty"` // Associated account bill data of type Bill customer data
	BillCompanyData  *BillCompanyData  `json:"BillCompanyData,omitempty"`  // Associated account bill data of type Bill company data
}

type BillOwnerData struct {
	Discriminator    string            `json:"Discriminator"` // Determines type of value.
	Value            json.RawMessage   // Structure of object depends on Bill owner data discriminator. Can be either of type Bill customer data or Bill company data.
	BillCustomerData *BillCustomerData `json:"BillCustomerData,omitempty"` // Owner data specific to a Customer
	BillCompanyData  *BillCompanyData  `json:"BillCompanyData,omitempty"`  // Owner data specific to a Company
}

func (b *BillOwnerData) UnmarshalJSON(data []byte) error {
	tmp := struct {
		Discriminator string          `json:"Discriminator"` // Determines type of value.
		Value         json.RawMessage // Structure of object depends on Bill owner data discriminator. Can be either of type Bill customer data or Bill company data.
	}{}

	err := json.Unmarshal(data, &tmp)
	if err != nil {
		return err
	}

	b.Discriminator = tmp.Discriminator
	b.Value = tmp.Value

	if tmp.Discriminator == "BillCustomerData" {
		err = json.Unmarshal(tmp.Value, &b.BillCustomerData)
		if err != nil {
			return err
		}
	} else if tmp.Discriminator == "BillCompanyData" {
		err = json.Unmarshal(tmp.Value, &b.BillCompanyData)
		if err != nil {
			return err
		}
	} else {
		// do nothing
	}

	return nil
}

type BillCustomerData struct {
	ID               string            `json:"Id"`
	Address          BillAddress       `json:"Address"`
	LegalIdentifiers map[string]string `json:"LegalIdentifiers"`
	BillingCode      string            `json:"BillingCode"`
	LastName         string            `json:"LastName"`
	FirstName        string            `json:"FirstName"`
	SecondLastName   string            `json:"SecondLastName"`
	TitlePrefix      string            `json:"TitlePrefix"`
}

type BillCompanyData struct {
	ID                      string      `json:"Id"`                      // ID of the Company.
	Address                 BillAddress `json:"Address"`                 // Address of the company.
	LegalIdentifiers        Dictionary  `json:"LegalIdentifiers"`        // The set of LegalIdentifiers for the company.
	BillingCode             string      `json:"BillingCode"`             // A unique code for Mews to list on invoices it sends to the company.
	Name                    string      `json:"Name"`                    // Name of the company.
	FiscalIdentifier        string      `json:"FiscalIdentifier"`        // Fiscal identifier of the company.
	AdditionalTaxIdentifier string      `json:"AdditionalTaxIdentifier"` // Additional tax identifier of the company.
}

type BillEnterpriseData struct {
	AdditionalTaxIdentifier string `json:"AdditionalTaxIdentifier"` // Enterprise additional tax identifier.
	CompanyName             string `json:"CompanyName"`             // Enterprise company name.
	BankAccount             string `json:"BankAccount"`             // Enterprise bank account.
	BankName                string `json:"BankName"`                // Enterprise bank name.
	IBAN                    string `json:"Iban"`                    // Enterprise IBAN (International Bank Account Number).
	BIC                     string `json:"Bic"`                     // Enterprise BIC (Bank Identifier Code).
}

type BillAddress struct {
	Line1           string `json:"Line1"`           // First line of the address.
	Line2           string `json:"Line2"`           // Second line of the address.
	City            string `json:"City"`            // City of the address.
	PostalCode      string `json:"PostalCode"`      // Postal code of the address.
	SubdivisionCode string `json:"SubdivisionCode"` // ISO 3166-2 code of the administrative division.
	CountryCode     string `json:"CountryCode"`     // ISO 3166-1 code of the country.
}

type Dictionary map[string]interface{}

type BillOptions struct {
	DisplayCustomer bool `json:"DisplayCustomer"` // Display customer information on a bill.
	DisplayTaxation bool `json:"DisplayTaxation"` // Display taxation detail on a bill.
	TrackReceivable bool `json:"TrackReceivable"` // Tracking of payments is enabled for bill, only applicable for Invoice.
	DisplayCID      bool `json:"DisplayCid"`      // Display CID number on bill, only applicable for Invoice.
}

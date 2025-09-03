package customers

import (
	"time"

	"github.com/omniboost/go-mews/configuration"
	base "github.com/omniboost/go-mews/json"
	"github.com/omniboost/go-mews/omitempty"
	"github.com/omniboost/go-mews/services"
)

const (
	endpointAll = "customers/getAll"
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

func (s *Service) NewAllRequest() *AllRequest {
	return &AllRequest{
		Extent: CustomersExtent{
			Customers: true,
			Addresses: true,
		},
	}
}

type AllRequest struct {
	base.BaseRequest

	CreatedUTC     configuration.TimeInterval `json:"CreatedUtc,omitempty"`     // Interval in which Customer was created.
	UpdatedUTC     configuration.TimeInterval `json:"UpdatedUtc,omitempty"`     // Interval in which Customer was updated.
	Extent         CustomersExtent            `json:"Extent,omitempty"`         // Extent of data to be returned.
	DeletedUTC     configuration.TimeInterval `json:"DeletedUtc,omitempty"`     // Interval in which Customer was deleted. ActivityStates value Deleted should be provided with this filter to get expected results.
	ActivityStates services.ActivityStates    `json:"ActivityStates,omitempty"` // Whether return only active, only deleted or both records.
	CustomerIDs    []string                   `json:"CustomerIds,omitempty"`    // Unique identifiers of Customers. Required if no other filter is provided.
	CompanyIDs     []string                   `json:"CompanyIds,omitempty"`     // Unique identifier of the Company the customer is associated with.
	Emails         []string                   `json:"Emails,omitempty"`         // Emails of the Customers.
	FirstNames     []string                   `json:"FirstNames,omitempty"`     // First names of the Customers.
	LastNames      []string                   `json:"LastNames,omitempty"`      // Last names of the Customers.
	LoyaltyCodes   []string                   `json:"LoyaltyCodes,omitempty"`   // Loyalty codes of the Customers.
	Limitation     base.Limitation            `json:"Limitation,omitempty"`     // Limitation on the quantity of data returned.
}

func (r AllRequest) MarshalJSON() ([]byte, error) {
	return omitempty.MarshalJSON(r)
}

type AllResponse struct {
	Customers Customers `json:"customers"`
	Cursor    string    `json:"Cursor"`
}

type CustomersExtent struct {
	Customers bool `json:"Customers"` // Whether the response should contain information about customers.
	Addresses bool `json:"Addresses"` // Whether the response should contain addresses of customers.
}

type Customers []Customer

type Customer struct {
	ID                      string                         `json:"Id"`                      // Unique identifier of the customer.
	ChainID                 string                         `json:"ChainId"`                 // Unique identifier of the chain.
	Number                  string                         `json:"Number"`                  // Number of the customer.
	Title                   Title                          `json:"Title"`                   // Title prefix of the customer.
	Sex                     Sex                            `json:"Sex"`                     // Sex of the customer.
	FirstName               string                         `json:"FirstName"`               // First name of the customer.
	LastName                string                         `json:"LastName"`                // Last name of the customer.
	SecondLastName          string                         `json:"SecondLastName"`          // Second last name of the customer.
	NationalityCode         string                         `json:"NationalityCode"`         // ISO 3166-1 alpha-2 country code (two letter country code) of the nationality.
	PreferredLanguageCode   string                         `json:"PreferredLanguageCode"`   // Language and culture code of the customer's preferred language, according to their profile. For example: en-GB, fr-CA.
	LanguageCode            string                         `json:"LanguageCode"`            // Language and culture code of the customers preferred language. E.g. en-US or fr-FR.
	BirthDate               string                         `json:"BirthDate"`               // Date of birth in ISO 8601 format.
	BirthPlace              string                         `json:"BirthPlace"`              // Place of birth.
	Occupation              string                         `json:"Occupation"`              // Occupation of the customer.
	Email                   string                         `json:"Email"`                   // Email address of the customer.
	HasOTAEmail             bool                           `json:"HasOtaEmail"`             // Whether the customer's email address is a temporary email address from an OTA. For more details, see the product documentation.
	Phone                   string                         `json:"Phone"`                   // Phone number of the customer (possibly mobile).
	TaxIdentificationNumber string                         `json:"TaxIdentificationNumber"` // tax id customer
	LoyaltyCode             string                         `json:"LoyaltyCode"`             // Loyalty code of the customer.
	AccountingCode          string                         `json:"AccountingCode"`          // Accounting code of the customer.
	BillingCode             string                         `json:"BillingCode"`             // Billing code of the customer.
	Notes                   string                         `json:"Notes"`                   // Internal notes about the customer.
	CarRegistrationNumber   string                         `json:"CarRegistrationNumber"`   // Registration number of the customer's car.
	DietaryRequirements     string                         `json:"DietaryRequirements"`     // Customer's dietary requirements, e.g. Vegan, Halal.
	CreatedUTC              time.Time                      `json:"CreatedUtc"`              // Creation date and time of the customer in UTC timezone in ISO 8601 format.
	UpdatedUTC              time.Time                      `json:"UpdatedUtc"`              // Last update date and time of the customer in UTC timezone in ISO 8601 format.
	AddressID               string                         `json:"AddressId"`               // Unique identifier of the Address of the customer.
	Classifications         []Classification               `json:"Classifications"`         // Classifications of the customer.
	Options                 Options                        `json:"Options"`                 // Options of the customer.
	Address                 configuration.Address          `json:"Address"`                 // Address of the customer.
	ItalianDestinationCode  string                         `json:"ItalianDestinationCode"`  // Value of Italian destination code.
	ItalianFiscalCode       string                         `json:"ItalianFiscalCode"`       // Value of Italian fiscal code.
	CompanyID               string                         `json:"CompanyId"`               // Unique identifier of Company the customer is associated with.
	MergeTargetID           string                         `json:"MergeTargetId"`           // Unique identifier of the account (Customer) to which this customer is linked.
	IsActive                bool                           `json:"IsActive"`                // Whether the customer record is still active.
	PreferredSpaceFeatures  ResourceFeatureClassifications // A list of room preferences, such as view type, bed type, and amenities.
	CreatorProfileID        string                         `json:"CreatorProfileId"` // Unique identifier of the user who created the customer.
	UpdaterProfileID        string                         `json:"UpdaterProfileId"` // Unique identifier of the user who last updated the customer.
}

type Title string

const (
	TitleMister Title = "Mister"
	TitleMiss   Title = "Miss"
	TitleMisses Title = "Missed"
)

type Sex string

const (
	SexMale   Title = "Male"
	SexFemale Title = "Female"
)

type Gender string

const (
	GenderMale   Gender = "Male"
	GenderFemale Gender = "Female"
)

type Classification string

type Options []string

type ResourceFeatureClassifications []ResourceFeatureClassification

type ResourceFeatureClassification string

const (
	SeaView            ResourceFeatureClassification = "SeaView"
	RiverView          ResourceFeatureClassification = "RiverView"
	OceanView          ResourceFeatureClassification = "OceanView"
	TwinBeds           ResourceFeatureClassification = "TwinBeds"
	DoubleBed          ResourceFeatureClassification = "DoubleBed"
	RollawayBed        ResourceFeatureClassification = "RollawayBed"
	UpperBed           ResourceFeatureClassification = "UpperBed"
	LowerBed           ResourceFeatureClassification = "LowerBed"
	Balcony            ResourceFeatureClassification = "Balcony"
	AccessibleBathroom ResourceFeatureClassification = "AccessibleBathroom"
	AccessibleRoom     ResourceFeatureClassification = "AccessibleRoom"
	ElevatorAccess     ResourceFeatureClassification = "ElevatorAccess"
	HighFloor          ResourceFeatureClassification = "HighFloor"
	Kitchenette        ResourceFeatureClassification = "Kitchenette"
	AirConditioning    ResourceFeatureClassification = "AirConditioning"
	PrivateJacuzzi     ResourceFeatureClassification = "PrivateJacuzzi"
	PrivateSauna       ResourceFeatureClassification = "PrivateSauna"
	EnsuiteRoom        ResourceFeatureClassification = "EnsuiteRoom"
	PrivateBathroom    ResourceFeatureClassification = "PrivateBathroom"
	SharedBathroom     ResourceFeatureClassification = "SharedBathroom"
)

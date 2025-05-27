package identitydocuments

import (
	"github.com/omniboost/go-mews/json"
)

const (
	endpointAll = "identityDocuments/getAll"
)

// List all outlets
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
	IdentityDocuments IdentityDocuments `json:"IdentityDocuments"`
	Cursor            string            `json:"Cursor"`
}

func (s *APIService) NewAllRequest() *AllRequest {
	return &AllRequest{}
}

type AllRequest struct {
	json.BaseRequest
	CustomerIDs []string        `json:"CustomerIds"` // Unique identifiers of Customer. max 100.
	Limitation  json.Limitation `json:"Limitation,omitempty"`
}

type IdentityDocuments []IdentityDocument

type IdentityDocument struct {
	ID                            string     `json:"Id"`                            // Unique identifier of the document.
	CustomerID                    string     `json:"CustomerId"`                    // Identifier of the Customer.
	Type                          string     `json:"Type"`                          // Type of the document.
	Number                        string     `json:"Number"`                        // Number of the document (e.g. passport number).
	ExpirationDate                *json.Date `json:"ExpirationDate"`                // Expiration date in ISO 8601 format.
	IssuanceDate                  *json.Date `json:"IssuanceDate"`                  // Date of issuance in ISO 8601 format.
	IssuingCountryCode            *string    `json:"IssuingCountryCode"`            // ISO 3166-1 code of the Country.
	IssuingCity                   *string    `json:"IssuingCity"`                   // City where the document was issued.
	IdentityDocumentSupportNumber *string    `json:"IdentityDocumentSupportNumber"` // Identity document support number. Only required for Spanish identity cards in Spanish hotels.
}

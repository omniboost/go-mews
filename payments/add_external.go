package payments

import (
	"github.com/omniboost/go-mews/accountingitems"
	base "github.com/omniboost/go-mews/json"
	"github.com/omniboost/go-mews/omitempty"
)

const (
	endpointAddExternal = "payments/addExternal"
)

// Adds a new external payment to a bill of the specified customer. An external
// payment represents a payment that is tracked outside of the system. Note this
// operation supports Portfolio Access Tokens.
func (s *Service) AddExternal(requestBody *AddExternalRequest) (*AddExternalResponse, error) {
	// @TODO: create wrapper?
	if err := s.Client.CheckTokens(); err != nil {
		return nil, err
	}

	apiURL, err := s.Client.GetApiURL(endpointAddExternal)
	if err != nil {
		return nil, err
	}

	responseBody := &AddExternalResponse{}
	httpReq, err := s.Client.NewRequest(apiURL, requestBody)
	if err != nil {
		return nil, err
	}

	_, err = s.Client.Do(httpReq, responseBody)
	return responseBody, err
}

type AddExternalResponse struct {
	ExternalPaymentID string `json:"ExternalPaymentId"` // Unique identifier of the Payment item.
}

func (s *Service) NewAddExternalRequest() *AddExternalRequest {
	return &AddExternalRequest{}
}

type AddExternalRequest struct {
	base.BaseRequest

	EnterpriseID         string                 `json:"EnterpriseId"`                   // Unique identifier of the Enterprise. Required when using a Portfolio Access Token, ignored otherwise.
	AccountID            string                 `json:"AccountId"`                      // Unique identifier of the Customer or Company. Company billing may not be enabled for your integration.
	BillID               string                 `json:"BillId,omitempty"`               // Unique identifier of an open bill of the customer where to assign the payment.
	ReservationID        string                 `json:"ReservationId,omitempty"`        // Unique identifier of the reservation the payment belongs to.
	Amount               accountingitems.Amount `json:"Amount"`                         // Amount of the external card payment.
	ExternalIdentifier   string                 `json:"ExternalIdentifier,omitempty"`   // Identifier of the payment from external system.
	Type                 ExternalPaymentType    `json:"Type,omitempty"`                 // Type of the external payment. *Except for the enterprises based in the French Legal Environment. Unspecified is considered as fraud.
	AccountingCategoryID string                 `json:"AccountingCategoryId,omitempty"` // Unique identifier of an Accounting category to be assigned to the external payment.
	Notes                string                 `json:"Notes,omitempty"`                // Additional payment notes.
}

func (r AddExternalRequest) MarshalJSON() ([]byte, error) {
	return omitempty.MarshalJSON(r)
}

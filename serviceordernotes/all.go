package serviceordernotes

import (
	"github.com/omniboost/go-mews/configuration"
	base "github.com/omniboost/go-mews/json"
)

const (
	endpointAll = "serviceordernotes/getAll"

	General        ServiceOrderNoteType = "General"
	ChannelManager ServiceOrderNoteType = "ChannelManager"
	SpecialRequest ServiceOrderNoteType = "SpecialRequest"
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
	ServiceOrderNotes ServiceOrderNotes `json:"ServiceOrderNotes"` // Services offered by the enterprise.
	Cursor            string            `json:"Cursor"`
}

type ServiceOrderNotes []ServiceOrderNote

type ServiceOrderNote struct {
	ID      string               `json:"Id"`      // Unique identifier of the service order note.
	OrderID string               `json:"OrderId"` // Unique identifier of the Service order to which the Service Order Note belongs.
	Text    string               `json:"Text"`    // Content of the service order note.
	Type    ServiceOrderNoteType `json:"Type"`    // A discriminator specifying the type of service order note, e.g. general or channel manager.
}

type ServiceOrderNoteType string

func (s *APIService) NewAllRequest() *AllRequest {
	return &AllRequest{}
}

type AllRequest struct {
	base.BaseRequest
	EnterpriseIDs       []string                    `json:"EnterpriseIds,omitempty"`       // Unique identifiers of the Enterprises. If not specified, the operation returns data for all enterprises within scope of the Access Token.
	ServiceOrderIDs     []string                    `json:"ServiceOrderIds,omitempty"`     // Unique identifiers of Service order. Reservation IDs or Order IDs can be used as service order identifiers.
	ServiceOrderNoteIds []string                    `json:"ServiceOrderNoteIds,omitempty"` // Unique identifiers of Service order note. Use this property if you want to fetch specific service order notes.
	UpdatedUTC          *configuration.TimeInterval `json:"UpdatedUtc,omitempty"`          // Timestamp in UTC timezone ISO 8601 format when the service order note was updated.
	Types               []ServiceOrderNoteType      `json:"Types,omitempty"`               // Type of the service order note. Defaults to ["General", "ChannelManager"].

	Limitation base.Limitation `json:"Limitation,omitempty"`
}

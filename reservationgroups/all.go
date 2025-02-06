package reservationgroups

import (
	"github.com/omniboost/go-mews/configuration"
	base "github.com/omniboost/go-mews/json"
	"github.com/omniboost/go-mews/omitempty"
)

const (
	endpointAll = "reservationgroups/getAll"
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

func (s *APIService) NewAllRequest() *AllRequest {
	return &AllRequest{}
}

type AllRequest struct {
	base.BaseRequest

	// Limitation on the quantity of reservation groups returned.
	Limitation base.Limitation `json:"Limitation,omitempty"`

	// Unique identifiers of the Enterprises. If not specified, the operation
	// returns data for all enterprises within scope of the Access Token.
	EnterpriseIDs []string `json:"EnterpriseIds,omitempty"`
	// Unique identifiers of the Reservation Group. Required if no other filter
	// is provided.
	ReservationGroupIDs []string `json:"ReservationGroupIds,omitempty"`
	// Interval in which the Reservation Group was updated. Required if no other
	// filter is provided.
	UpdatedUTC configuration.TimeInterval `json:"UpdatedUtc,omitempty"`
}

func (r AllRequest) MarshalJSON() ([]byte, error) {
	return omitempty.MarshalJSON(r)
}

type AllResponse struct {
	ReservationGroups ReservationGroups `json:"ReservationGroups"`
	Cursor            string            `json:"Cursor"`
}

type ReservationGroups []ReservationGroup

type ReservationGroup struct {
	// Unique identifier of the reservation group.
	ID string `json:"Id"`
	// Unique identifier of the Enterprise the reservation group belongs to.
	EnterpriseID string `json:"EnterpriseId"`
	// Name of the reservation group, might be empty or same for multiple groups.
	Name string `json:"Name"`
	// Name of the corresponding channel manager.
	ChannelManager string `json:"ChannelManager"`
	// Identifier of the channel manager.
	ChannelManagerGroupNumber string `json:"ChannelManagerGroupNumber"`
}

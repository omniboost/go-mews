package companionships

import (
	"github.com/omniboost/go-mews/configuration"
	"github.com/omniboost/go-mews/customers"
	base "github.com/omniboost/go-mews/json"
	"github.com/omniboost/go-mews/omitempty"
	"github.com/omniboost/go-mews/reservationgroups"
	"github.com/omniboost/go-mews/reservations"
)

const (
	endpointAll = "companionships/getAll"
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
	return &AllRequest{}
}

type AllRequest struct {
	base.BaseRequest

	// Limitation on the quantity of reservation groups returned.
	Limitation base.Limitation `json:"Limitation,omitempty"`

	// Unique identifiers of the Enterprises. If not specified, the operation
	// returns data for all enterprises within scope of the Access Token.
	EnterpriseIDs []string `json:"EnterpriseIds,omitempty"`
	// Unique identifiers of Companionship.
	CompanionshipIDs []string `json:"CompanionshipIds,omitempty"`
	// Unique identifiers of Customers.
	CustomerID string `json:"CustomerId,omitempty"`
	// Unique identifiers of reservations.
	ReservationIDs []string `json:"ReservationIds,omitempty"`
	// Unique identifiers of Reservation groups.
	ReservationGroupIDs []string `json:"ReservationGroupIds,omitempty"`
	// Interval in which the Companionship was updated.
	UpdatedUTC configuration.TimeInterval `json:"UpdatedUtc,omitempty"`
	// Extent of data to be returned. E.g. it is possible to specify that
	// together with the companionships, customers, reservations, and
	// reservation groups should be also returned.
	Extent CompanionshipsExtent `json:"Extent,omitempty"`
}

func (r AllRequest) MarshalJSON() ([]byte, error) {
	return omitempty.MarshalJSON(r)
}

type AllResponse struct {
	Cursor string `json:"Cursor"`

	Companionships    companionships                      `json:"Companionships"`
	Customers         customers.Customers                 `json:"Customers"`
	Reservations      reservations.Reservations           `json:"Reservations"`
	ReservationGroups reservationgroups.ReservationGroups `json:"ReservationGroups"`
}

type companionships []Companionship

type Companionship struct {
	// Unique identifier of Companionship.
	ID string `json:"Id"`
	// Unique identifier of Customer.
	CustomerID string `json:"CustomerId"`
	// Unique identifier of reservation.
	ReservationID string `json:"ReservationId"`
	// Unique identifier of Reservation group.
	ReservationGroupID string `json:"ReservationGroupId"`
}

type CompanionshipsExtent struct {
	// Whether the response should contain customers.
	Customers bool `json:"Customers"`
	// Whether the response should contain reservations.
	Reservations bool `json:"Reservations"`
	// Whether the response should contain reservation groups.
	ReservationGroups bool `json:"ReservationGroups"`
}

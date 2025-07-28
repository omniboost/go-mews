package services

import (
	"encoding/json"

	base "github.com/omniboost/go-mews/json"
)

const (
	endpointAll = "services/getAll"

	Reservable ServiceType = "Reservable"
	Orderable  ServiceType = "Orderable"
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
	Services Services `json:"Services"` // Services offered by the enterprise.
	Cursor   string   `json:"Cursor"`
}

type Services []Service

type Service struct {
	ID         string         `json:"Id"`         // Unique identifier of the service.
	IsActive   bool           `json:"IsActive"`   // Whether the service is still active.
	Name       string         `json:"Name"`       // Name of the service.
	StartTime  string         `json:"StartTime"`  // Default start time of the service orders in ISO 8601 duration format.
	EndTime    string         `json:"EndTime"`    // Default end time of the service orders in ISO 8601 duration format.
	Promotions Promotions     `json:"Promotions"` // Promotions of the service.
	Type       ServiceType    `json:"Type"`       // Type of the service
	Options    ServiceOptions `json:"Options"`    // Options of the service
	Data       ServiceData    `json:"Data"`       // Additional information about the specific service
}

func (s *APIService) NewAllRequest() *AllRequest {
	return &AllRequest{}
}

type AllRequest struct {
	base.BaseRequest
	ServiceIDs []string        `json:"ServiceIds,omitempty"` // Unique identifiers of the requested Services.
	Limitation base.Limitation `json:"Limitation,omitempty"`
}

type Promotions struct {
	BeforeCheckIn  bool `json:"BeforeCheckIn"`  // Whether it can be promoted before check-in.
	AfterCheckIn   bool `json:"AfterCheckIn"`   // Whether it can be promoted after check-in.
	DuringStay     bool `json:"DuringStay"`     // Whether it can be promoted during stay.
	BeforeCheckOut bool `json:"BeforeCheckOut"` // Whether it can be promoted before check-out.
	AfterCheckOut  bool `json:"AfterCheckOut"`  // Whether it can be promoted after check-out.
	DuringCheckOut bool `json:"DuringCheckOut"` // Whether it can be promoted during check-out.
}

type ServiceType string

const (
	ServiceReservable ServiceType = "Reservable"
	ServiceOrderable  ServiceType = "Orderable"
)

type ActivityStates []ActivityState

type ActivityState string

const (
	ActivityStateActive  ActivityState = "Active"
	ActivityStateDeleted ActivityState = "Deleted"
)

type ServiceOptions struct {
	BillAsPackage bool `json:"BillAsPackage"` // Products should be displayed as a single package instead of individual items.
}

type ServiceData struct {
	Discriminator   ServiceDataDiscriminator `json:"Discriminator"` // Determines type of value
	Value           json.RawMessage          `json:"Value"`         // Structure of object depends on Service data discriminator.
	BookableValue   BookableServiceData      `json:"-"`
	AdditionalValue AdditionalServiceData    `json:"-"`
}

func (d *ServiceData) UnmarshalJSON(data []byte) error {
	type alias ServiceData
	a := alias(*d)
	err := json.Unmarshal(data, &a)
	if err != nil {
		return err
	}

	err = json.Unmarshal(a.Value, &a.BookableValue)
	if err != nil {
		return err
	}

	err = json.Unmarshal(a.Value, &a.AdditionalValue)
	if err != nil {
		return err
	}

	*d = ServiceData(a)
	return nil
}

type ServiceDataDiscriminator string

type BookableServiceData struct {
	StartOffset          base.Duration `json:"StartOffset"`          // Offset from the start of the time unit which defines the default start of the service; expressed in ISO 8601 duration format.
	EndOffset            base.Duration `json:"EndOffset"`            // Offset from the end of the time unit which defines the default end of the service; expressed in ISO 8601 duration format.
	OccupancyStartOffset base.Duration `json:"OccupancyStartOffset"` // Offset from the end of the time unit which defines the default end of the service; expressed in ISO 8601 duration format.
	OccupancyEndOffset   base.Duration `json:"OccupancyEndOffset"`   // Offset from the end of the time unit which defines the occupancy end of the service; expressed in ISO 8601 duration format. 'Occupancy end' is used for availability and reporting purposes, it implies the time at which the booked resource is no longer considered occupied.
	TimeUnitPeriod       base.TimeUnit `json:"TimeUnitPeriod"`       // The length of time or period represented by a time unit, for which the service can be booked.
}

type AdditionalServiceData struct {
	Promotions Promotions `json:"Promotions"` // Promotions of the service.
}

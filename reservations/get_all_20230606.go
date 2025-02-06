package reservations

import (
	"time"

	"github.com/omniboost/go-mews/configuration"
	base "github.com/omniboost/go-mews/json"
	"github.com/omniboost/go-mews/omitempty"
)

const (
	endpointGetAll = "reservations/getAll/2023-06-06"
)

// List all products
func (s *APIService) GetAll20230606(requestBody *GetAll20230606Request) (*AllResponse20230606, error) {
	// @TODO: create wrapper?
	if err := s.Client.CheckTokens(); err != nil {
		return nil, err
	}

	apiURL, err := s.Client.GetApiURL(endpointGetAll)
	if err != nil {
		return nil, err
	}

	responseBody := &AllResponse20230606{}
	httpReq, err := s.Client.NewRequest(apiURL, requestBody)
	if err != nil {
		return nil, err
	}
	// s.Client.Debug = true
	_, err = s.Client.Do(httpReq, responseBody)
	return responseBody, err
}

type AllResponse20230606 struct {
	Reservations Reservations20230606
	Cursor       string `json:"Cursor"`
}

func (s *APIService) NewGetAll20230606Request() *GetAll20230606Request {
	return &GetAll20230606Request{}
}

type GetAll20230606Request struct {
	base.BaseRequest
	Limitation base.Limitation `json:"Limitation,omitempty"`

	EnterpriseIDs       []string                   `json:"EnterpriseIds,omitempty"`       // Unique identifiers of the Enterprises.
	ReservationIDs      []string                   `json:"ReservationIds,omitempty"`      // Unique identifiers of the Reservations.
	ServiceIDs          []string                   `json:"ServiceIds,omitempty"`                    // Unique identifiers of the Services. If not provided, all bookable services are used.
	AccountIDs          []string                   `json:"AccountIds,omitempty"`          // Unique identifiers of accounts (currently only Customers, in the future also Companies) the reservation is associated with.
	ReservationGroupIDs []string                   `json:"ReservationGroupIds,omitempty"` // Unique identifiers of Reservation groups.
	AssignedResourceIds []string                   `json:"AssignedResourceIds,omitempty"`
	States              []ReservationStates        `json:"States,omitempty"`
	CreatedUTC          configuration.TimeInterval `json:"CreatedUtc,omitempty"`
	UpdatedUTC          configuration.TimeInterval `json:"UpdatedUtc,omitempty"`
	ScheduledStartUTC   configuration.TimeInterval `json:"ScheduledStartUtc,omitempty"`
	ScheduledEndUTC     configuration.TimeInterval `json:"ScheduledEndUtc,omitempty"`
	ActualStartUTC      configuration.TimeInterval `json:"ActualStartUtc,omitempty"`
	CollidingUTC        configuration.TimeInterval `json:"CollidingUtc,omitempty"`
}

func (r GetAll20230606Request) MarshalJSON() ([]byte, error) {
	return omitempty.MarshalJSON(r)
}

type Reservations20230606 []Reservation20230606

type Reservation20230606 struct {
	ID                          string              `json:"ID,omitempty"`
	EnterpriseID                string              `json:"EnterpriseID,omitempty"`
	ServiceID                   string              `json:"ServiceID,omitempty"`
	AccountID                   string              `json:"AccountID,omitempty"`
	AccountType                 string              `json:"AccountType,omitempty"`
	CreatorProfileID            string              `json:"CreatorProfileID,omitempty"`
	UpdaterProfileID            string              `json:"UpdaterProfileID,omitempty"`
	BookerID                    string              `json:"BookerId,omitempty"`
	StartUTC                    time.Time           `json:"StartUtc,omitempty"`
	EndUTC                      time.Time           `json:"EndUtc,omitempty"`
	Number                      string              `json:"Number,omitempty"`
	State                       ReservationStates   `json:"State"`
	Origin                      string              `json:"Origin,omitempty"`
	OriginDetails               string              `json:"OriginDetails,omitempty"`
	CommanderOrigin             CommanderOrigin     `json:"CommanderOrigin,omitempty"`
	CreatedUTC                  time.Time           `json:"CreatedUtc,omitempty"`
	UpdatedUTC                  time.Time           `json:"UpdatedUtc,omitempty"`
	ReleasedUTC                 time.Time           `json:"ReleasedUtc,omitempty"`
	CancelledUTC                time.Time           `json:"CancelledUtc,omitempty"`
	VoucherID                   string              `json:"VoucherId,omitempty"`
	BusinessSegmentID           string              `json:"BusinessSegmentId,omitempty"`
	RateID                      string              `json:"RateId,omitempty"`
	CreditCardID                string              `json:"CreditCardId,omitempty"`
	GroupID                     string              `json:"GroupId,omitempty"`
	RequestedResourceCategoryID string              `json:"RequestedResourceCategoryId,omitempty"`
	AssignedResourceID          string              `json:"AssignedResourceId,omitempty"`
	AvailabilityBlockID         string              `json:"AvailabilityBlockId,omitempty"`
	CompanyID                   string              `json:"CompanyId,omitempty"`
	TravelAgencyID              string              `json:"TravelAgencyId,omitempty"`
	AssignedResourceLocked      bool                `json:"AssignedSpaceLocked,omitempty"`
	ChannelNumber               string              `json:"ChannelNumber,omitempty"`
	ChannelManagerNumber        string              `json:"ChannelManagerNumber,omitempty"`
	CancellationReason          string              `json:"CancellationReason,omitempty"` // ??
	Purpose                     string              `json:"Purpose,omitempty"`            // ??
	Options                     ServiceOrderOptions `json:"Options,omitempty"`            // ??
	PersonCounts                []PersonCounts      `json:"PersonCounts,omitempty"`       // ??
}

type CommanderOrigin string

const (
	CommanderOriginInPerson   CommanderOrigin = "InPerson"
	CommanderOriginChannel    CommanderOrigin = "Channel"
	CommanderOriginPhone      CommanderOrigin = "Phone"
	CommanderOriginEmail      CommanderOrigin = "Email"
	CommanderOriginWebsite    CommanderOrigin = "Website"
	CommanderOriginMessage    CommanderOrigin = "Message"
	CommanderOriginCallCenter CommanderOrigin = "CallCenter"
)

type ReservationStates string

const (
	ReservationStatesEnquired  ReservationStates = "Enquired"
	ReservationStatesRequested ReservationStates = "Requested"
	ReservationStatesOptional  ReservationStates = "Optional"
	ReservationStatesConfirmed ReservationStates = "Confirmed"
	ReservationStatesStarted   ReservationStates = "Started"
	ReservationStatesProcessed ReservationStates = "Processed"
	ReservationStatesCanceled  ReservationStates = "Canceled"
)

type PersonCounts struct {
	AgeCategoryID string `json:"AgeCategoryId,omitempty"`
	Count         int    `json:"Count,omitempty"`
}

type ServiceOrderOptions struct {
	OwnerCheckedIn         bool `json:"OwnerCheckedIn,omitempty"`
	AllCompanionsCheckedIn bool `json:"AllCompanionsCheckedIn,omitempty"`
	AnyCompanionCheckedIn  bool `json:"AnyCompanionCheckedIn,omitempty"`
	ConnectorCheckIn       bool `json:"ConnectorCheckIn,omitempty"`
}

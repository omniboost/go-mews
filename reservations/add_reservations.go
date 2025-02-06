package reservations

import (
	"time"

	"github.com/omniboost/go-mews/json"
	"github.com/omniboost/go-mews/omitempty"
	"github.com/omniboost/go-mews/orderitems"
)

const (
	endpointAdd = "reservations/add"
)

// Add customer
func (s *APIService) Add(requestBody *AddRequest) (*AddResponse, error) {
	// @TODO: create wrapper?
	if err := s.Client.CheckTokens(); err != nil {
		return nil, err
	}

	apiURL, err := s.Client.GetApiURL(endpointAdd)
	if err != nil {
		return nil, err
	}

	responseBody := &AddResponse{}
	httpReq, err := s.Client.NewRequest(apiURL, requestBody)
	if err != nil {
		return nil, err
	}

	_, err = s.Client.Do(httpReq, responseBody)
	return responseBody, err
}

func (s *APIService) NewAddRequest() *AddRequest {
	return &AddRequest{}
}

type AddRequest struct {
	json.BaseRequest

	EnterpriseId           string                 `json:"EnterpriseId"`                     // Unique identifier of the enterprise. Required when using Portfolio Access Tokens, ignored otherwise.
	ServiceID              string                 `json:"ServiceId"`                        // Unique identifier of the Service to be reserved.
	GroupID                string                 `json:"GroupId,omitempty"`                // Unique identifier of the Reservation group where the reservations are added. If not specified, a new group is created.
	GroupName              string                 `json:"GroupName,omitempty"`              // Name of the Reservation group which the reservations are added to. If GroupId is specified, this field is ignored. If not specified, the group name is automatically created.
	SendConfirmationEmail  bool                   `json:"SendConfirmationEmail,omitempty"`  // Whether the confirmation email is sent. Default value is true.
	CheckRateApplicability bool                   `json:"CheckRateApplicability,omitempty"` // Indicates whether the system will check and prevent a booking being made using a restricted rate, e.g. a private rate. The default is true, i.e. the system will normally check for this unless the property is set to false.
	CheckOverbooking       bool                   `json:"CheckOverbooking,omitempty"`       // Indicates whether the system will check and prevent a booking being made in the case of an overbooking, i.e. where there is an insufficient number of resources available to meet the request*1. The default is true, i.e. the system will normally check for this unless the property is set to false.
	Reservations           AddRequestReservations `json:"Reservations"`                     // Parameters of the new reservations.
}

func (r AddRequest) MarshalJSON() ([]byte, error) {
	return omitempty.MarshalJSON(r)
}

type AddRequestReservations []AddRequestReservation
type AddRequestReservation struct {
	Identifier          string            `json:"Identifier,omitempty"`     // Identifier of the reservation within the transaction.
	State               string            `json:"State,omitempty"`          // State of the newly created reservation (either Optional, Enquired or Confirmed). If not specified, Confirmed is used.
	StartUtc            time.Time         `json:"StartUtc"`                 // Reservation start in UTC timezone in ISO 8601 format.
	EndUtc              time.Time         `json:"EndUtc"`                   // Reservation end in UTC timezone in ISO 8601 format.
	ReleasedUtc         *time.Time        `json:"ReleasedUtc,omitempty"`    // Release date and time of an unconfirmed reservation in UTC timezone in ISO 8601 format.
	CustomerID          string            `json:"CustomerId"`               // Unique identifier of the Customer who owns the reservation.
	BookerID            string            `json:"BookerId,omitempty"`       // Unique identifier of the Customer on whose behalf the reservation was made.
	RequestedCategoryID string            `json:"RequestedCategoryId"`      // Identifier of the requested Resource category.
	RateID              string            `json:"RateId"`                   // Identifier of the reservation Rate.
	TravelAgencyID      string            `json:"TravelAgencyId,omitempty"` // Identifier of the Company that mediated the reservation.
	CompanyID           string            `json:"CompanyId,omitempty"`      // Identifier of the Company on behalf of which the reservation was made.
	Notes               string            `json:"Notes,omitempty"`          // Additional notes.
	TimeUnitAmount      orderitems.Amount `json:"TimeUnitAmount,omitempty"` // Amount of each night of the reservation.
	PersonCounts        []struct {
		AgeCategoryID string `json:"AgeCategoryId"` // Unique identifier of the Age category.
		Count         int    `json:"Count"`         // Number of people of a given age category. Only positive value is accepted.
	} `json:"PersonCounts,omitempty"` // Number of people per age category the reservation was booked for. At least one category with valid count must be provided.
	TimeUnitPrices []struct {
		Index  int               `json:"Index"`  // Index of the unit. Indexing starts with 0. E.g the first night of the reservation has index 0.
		Amount orderitems.Amount `json:"Amount"` // Amount of the unit.
	} `json:"TimeUnitPrices,omitempty"` // Prices for time units of the reservation. E.g. prices for the first or second night.
	ProductOrders []struct {
		ProductID string    `json:"ProductId"` // Unique identifier of the Product whose prices should be returned.
		StartUtc  time.Time `json:"StartUtc"`  // Start of the time interval, expressed as the timestamp for the start of the first time unit, in UTC timezone ISO 8601 format. See Time units.
		EndUtc    time.Time `json:"EndUtc"`    // End of the time interval, expressed as the timestamp for the start of the last time unit, in UTC timezone ISO 8601 format. See Time units. The maximum size of time interval depends on the service's time unit: 100 hours if hours, 100 days if days, or 24 months if months.
	} `json:"ProductOrders,omitempty"` // Parameters of the products ordered together with the reservation.
	CreditCardID        string `json:"CreditCardId,omitempty"`        // Identifier of Credit card belonging to Customer who owns the reservation.
	AvailabilityBlockID string `json:"AvailabilityBlockId,omitempty"` // Unique identifier of the Availability block the reservation is assigned to.
	VoucherCode         string `json:"VoucherCode,omitempty"`         // Voucher code value providing access to specified private Rate applied to this reservation.
	ChannelNumber       string `json:"ChannelNumber,omitempty"`       // Channel number of the reservation.
	BusinessSegmentID   string `json:"BusinessSegmentId,omitempty"`   // Identifier of the reservation BusinessSegment.
	Origin              string `json:"Origin,omitempty"`              // Origin of the reservation.
}

func (r AddRequestReservation) MarshalJSON() ([]byte, error) {
	return omitempty.MarshalJSON(r)
}

type AddResponse struct {
	Reservations []struct {
		Identifier  string      `json:"Identifier"`  // Identifier of the reservation within the transaction.
		Reservation Reservation `json:"Reservation"` // The added reservations.
	} `json:"Reservations"`
}

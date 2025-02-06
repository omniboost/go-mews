package reservations

import (
	"github.com/omniboost/go-mews/json"
	"github.com/omniboost/go-mews/omitempty"
	"github.com/omniboost/go-mews/orderitems"
)

const (
	endpointUpdate = "reservations/update"
)

// Update customer
func (s *APIService) Update(requestBody *UpdateRequest) (*UpdateResponse, error) {
	// @TODO: create wrapper?
	if err := s.Client.CheckTokens(); err != nil {
		return nil, err
	}

	apiURL, err := s.Client.GetApiURL(endpointUpdate)
	if err != nil {
		return nil, err
	}

	responseBody := &UpdateResponse{}
	httpReq, err := s.Client.NewRequest(apiURL, requestBody)
	if err != nil {
		return nil, err
	}

	_, err = s.Client.Do(httpReq, responseBody)
	return responseBody, err
}

func (s *APIService) NewUpdateRequest() *UpdateRequest {
	return &UpdateRequest{}
}

type UpdateRequest struct {
	json.BaseRequest

	Reason                 string             `json:"Reason"`                 // Reason for updating the reservation. Required when updating the price of the reservation.
	CheckRateApplicability bool               `json:"CheckRateApplicability"` // Indicates whether the system will check and prevent a booking being made using a restricted rate, e.g. a private rate. The default is true, i.e. the system will normally check for this unless the property is set to false.
	CheckOverbooking       bool               `json:"CheckOverbooking"`       // Indicates whether the system will check and prevent a booking being made in the case of an overbooking, i.e. where there is an insufficient number of resources available to meet the request*1. The default is true, i.e. the system will normally check for this unless the property is set to false.
	Reprice                bool               `json:"Reprice"`                // Whether the price should be updated to latest value for date/rate/category combination set in Mews. If not specified, the reservation price is updated.
	ApplyCancellationFee   bool               `json:"ApplyCancellationFee"`   // Whether the cancellation fees should be applied according to rate cancellation policies. If not specified, the cancellation fees are applied.
	ReservationUpdates     ReservationUpdates `json:"ReservationUpdates"`     // Array of properties to be updated in each reservation specified.
}

func (r UpdateRequest) MarshalJSON() ([]byte, error) {
	return omitempty.MarshalJSON(r)
}

type ReservationUpdates []ReservationUpdate

type ReservationUpdate struct {
	ReservationID          string                         `json:"ReservationId"`                    // Unique identifier of the reservation.
	StartUTC               *StringUpdateValue             `json:"StartUtc,omitempty"`               // Reservation start in UTC timezone in ISO 8601 format. (or null if the start time should not be updated).
	EndUTC                 *StringUpdateValue             `json:"EndUtc,omitempty"`                 // Reservation end in UTC timezone in ISO 8601 format. (or null if the end time should not be updated).
	AssignedResourceID     *StringUpdateValue             `json:"AssignedResourceId,omitempty"`     // Identifier of the assigned Resource.
	AssignedResourceLocked *BoolUpdateValue               `json:"AssignedResourceLocked,omitempty"` // Whether the reservation should be locked to the assigned Resource. Unlocking and assigning reservation to new Resource can be done in one call.
	ChannelNumber          *StringUpdateValue             `json:"ChannelNumber,omitempty"`          // Number of the reservation within the Channel (i.e. OTA, GDS, CRS, etc) in case the reservation group originates there (e.g. Booking.com confirmation number) (or null if the channel number should not be updated).
	RequestedCategoryID    *StringUpdateValue             `json:"RequestedCategoryId,omitempty"`    // Identifier of the requested Resource category (or null if resource category should not be updated).
	TravelAgencyID         *StringUpdateValue             `json:"TravelAgencyId,omitempty"`         // Identifier of the Company that mediated the reservation (or null if travel agency should not be updated).
	CompanyID              *StringUpdateValue             `json:"CompanyId,omitempty"`              // Identifier of the Company on behalf of which the reservation was made (or null if company should not be updated).
	BusinessSegmentID      *StringUpdateValue             `json:"BusinessSegmentId,omitempty"`      // Identifier of the reservation Business segment (or null if the business segment should not be updated).
	Purpose                *StringUpdateValue             `json:"Purpose,omitempty"`                // Purpose of the reservation (or null if the purpose should not be updated).
	RateID                 *StringUpdateValue             `json:"RateId,omitempty"`                 // Identifier of the reservation Rate (or null if the rate should not be updated).
	BookerID               *StringUpdateValue             `json:"BookerId,omitempty"`               // Identifier of the Customer on whose behalf the reservation was made. (or null if the booker should not be updated).
	TimeUnitPrices         *TimeUnitAmountUpdateValue     `json:"TimeUnitPrices,omitempty"`         // Prices for time units of the reservation. E.g. prices for the first or second night. (or null if the unit amounts should not be updated).
	PersonCounts           *PersonCountsUpdateValue       `json:"PersonCounts,omitempty"`           // Number of people per age category the reservation is for. Is supplied the person counts will be replaced. (or null if the person counts should not be updated).
	CreditCardID           *StringUpdateValue             `json:"CreditCardId,omitempty"`           // Identifier of Credit card belonging to Customer who owns the reservation. (or null if the credit card should not be updated).
	AvailabilityBlockID    *StringUpdateValue             `json:"AvailabilityBlockId,omitempty"`    // Unique identifier of the Availability block the reservation is assigned to.
	Options                *ReservationsOptionsParameters `json:"Options,omitempty"`                // Options of the reservations.
}

type StringUpdateValue struct {
	Value *string `json:"Value"`
}

type BoolUpdateValue struct {
	Value *bool `json:"Value"`
}

type PersonCountsUpdateValue struct {
	Value PersonCounts `json:"Value"` // Value which is to be updated.
}

type ReservationsOptionsParameters struct {
	OwnerCheckedIn BoolUpdateValue `json:"OwnerCheckedIn"` // True if the owner of the reservation is checked in.
}

type TimeUnitAmountUpdateValue struct {
	Value TimeUnitAmount // Value which is to be updated.
}

type TimeUnitAmount struct {
	Index  int               `json:"Index"`  // Index of the unit. Indexing starts with 0. E.g the first night of the reservation has index 0."
	Amount orderitems.Amount `json:"Amount"` // Amount of the unit.
}

// Same structure as in Get all reservations operation.
type UpdateResponse AllResponse

package resources

import (
	"time"

	"github.com/omniboost/go-mews/configuration"
	base "github.com/omniboost/go-mews/json"
	"github.com/omniboost/go-mews/omitempty"
)

const (
	endpointResourceCategoryAssignmentsAll = "resourceCategoryAssignments/getAll"
)

// List all products
func (s *APIService) CategoryAssignmentsAll(requestBody *CategoryAssignmentsAllRequest) (*CategoryAssignmentsAllResponse, error) {
	// @TODO: create wrapper?
	if err := s.Client.CheckTokens(); err != nil {
		return nil, err
	}

	apiURL, err := s.Client.GetApiURL(endpointResourceCategoryAssignmentsAll)
	if err != nil {
		return nil, err
	}

	responseBody := &CategoryAssignmentsAllResponse{}
	httpReq, err := s.Client.NewRequest(apiURL, requestBody)
	if err != nil {
		return nil, err
	}

	_, err = s.Client.Do(httpReq, responseBody)
	return responseBody, err
}

func (s *APIService) NewCategoryAssignmentsAllRequest() *CategoryAssignmentsAllRequest {
	return &CategoryAssignmentsAllRequest{}
}

type CategoryAssignmentsAllRequest struct {
	base.BaseRequest
	Limitation base.Limitation `json:"Limitation,omitempty"`

	EnterpriseIDs                 []string                   `json:"EnterpriseIds,omitempty"`                 // Unique identifiers of the Enterprises. If not specified, the operation returns the resource category assignments for all enterprises within scope of the Access Token.
	ResourceCategoryIDs           []string                   `json:"ResourceCategoryIds,omitempty"`           // Unique identifiers of Resource categories to which the resource category assignment belong.
	ResourceCategoryAssignmentIDs []string                   `json:"ResourceCategoryAssignmentIds,omitempty"` // Unique identifiers of Resource category assignment
	UpdatedUTC                    configuration.TimeInterval `json:"UpdatedUtc,omitempty"`                    // Interval in which the resource category assignments were updated.
	ActivityStates                ActivityStates             `json:"ActivityStates,omitempty"`                // Whether to return only active, only deleted or both records.
}

func (r CategoryAssignmentsAllRequest) MarshalJSON() ([]byte, error) {
	return omitempty.MarshalJSON(r)
}

type CategoryAssignmentsAllResponse struct {
	ResourceCategoryAssignments ResourceCategoryAssignments `json:"ResourceCategoryAssignments"` // Resource category assignments.
	Cursor                      string                      `json:"Cursor"`                      // Unique identifier of the last and hence oldest resource category assignment returned. This can be used in Limitation in a subsequent request to fetch the next batch of older resource category assignments.
}

type ResourceCategoryAssignments []ResourceCategoryAssignment

type ResourceCategoryAssignment struct {
	ID         string    `json:"Id"`         // Unique identifier of the assignment.
	IsActive   bool      `json:"IsActive"`   // Whether the assignment is still active.
	CategoryID string    `json:"CategoryId"` // Unique identifier of the Resource category.
	ResourceID string    `json:"ResourceId"` // Unique identifier of the Resource assigned to the Resource category.
	CreatedUTC time.Time `json:"CreatedUtc"` // Creation date and time of the assignment in UTC timezone in ISO 8601 format.
	UpdatedUTC time.Time `json:"UpdatedUtc"` // Last update date and time of the assignment in UTC timezone in ISO 8601 format.
}

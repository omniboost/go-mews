package resources

import (
	"time"

	"github.com/omniboost/go-mews/configuration"
	base "github.com/omniboost/go-mews/json"
	"github.com/omniboost/go-mews/omitempty"
)

const (
	endpointResourceFeatureAssignmentsAll = "resourceFeatureAssignments/getAll"
)

// List all products
func (s *APIService) FeatureAssignmentsAll(requestBody *FeatureAssignmentsAllRequest) (*FeatureAssignmentsAllResponse, error) {
	// @TODO: create wrapper?
	if err := s.Client.CheckTokens(); err != nil {
		return nil, err
	}

	apiURL, err := s.Client.GetApiURL(endpointResourceFeatureAssignmentsAll)
	if err != nil {
		return nil, err
	}

	responseBody := &FeatureAssignmentsAllResponse{}
	httpReq, err := s.Client.NewRequest(apiURL, requestBody)
	if err != nil {
		return nil, err
	}

	_, err = s.Client.Do(httpReq, responseBody)
	return responseBody, err
}

func (s *APIService) NewFeatureAssignmentsAllRequest() *FeatureAssignmentsAllRequest {
	return &FeatureAssignmentsAllRequest{}
}

type FeatureAssignmentsAllRequest struct {
	base.BaseRequest
	Limitation base.Limitation `json:"Limitation,omitempty"`

	EnterpriseIDs                []string                   `json:"EnterpriseIds,omitempty"`                // Unique identifiers of the Enterprises. If not specified, the operation returns the resource feature assignments for all enterprises within scope of the Access Token.
	ResourceFeatureIDs           []string                   `json:"ResourceFeatureIds,omitempty"`           // Unique identifiers of Resource features to which the resource feature assignments belong.
	ResourceFeatureAssignmentIDs []string                   `json:"ResourceFeatureAssignmentIds,omitempty"` // Unique identifiers of Resource features assignment
	UpdatedUTC                   configuration.TimeInterval `json:"UpdatedUtc,omitempty"`                   // Interval in which the resource features assignments were updated.
	ActivityStates               ActivityStates             `json:"ActivityStates,omitempty"`               // Whether to return only active, only deleted or both records.
}

func (r FeatureAssignmentsAllRequest) MarshalJSON() ([]byte, error) {
	return omitempty.MarshalJSON(r)
}

type FeatureAssignmentsAllResponse struct {
	ResourceFeatureAssignments ResourceFeatureAssignments `json:"ResourceFeatureAssignments"` // Resource features assignments.
	Cursor                     string                     `json:"Cursor"`                     // Unique identifier of the last and hence oldest resource feature assignment returned. This can be used in Limitation in a subsequent request to fetch the next batch of older resource feature assignments.
}

type ResourceFeatureAssignments []ResourceFeatureAssignment

type ResourceFeatureAssignment struct {
	ID         string    `json:"Id"`         // Unique identifier of the assignment.
	IsActive   bool      `json:"IsActive"`   // Whether the assignment is still active.
	ResourceID string    `json:"ResourceId"` // Unique identifier of the Resource.
	FeatureID  string    `json:"FeatureId"`  // Unique identifier of the Resource feature assigned to the Resource.
	CreatedUTC time.Time `json:"CreatedUtc"` // Creation date and time of the assignment in UTC timezone in ISO 8601 format.
	UpdatedUTC time.Time `json:"UpdatedUtc"` // Last update date and time of the assignment in UTC timezone in ISO 8601 format.
}

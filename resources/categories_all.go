package resources

import (
	"github.com/omniboost/go-mews/configuration"
	"github.com/omniboost/go-mews/json"
	"github.com/omniboost/go-mews/omitempty"
)

const (
	endpointResourceCategoriesAll = "resourceCategories/getAll"
)

// List all products
func (s *APIService) CategoriesAll(requestBody *CategoriesAllRequest) (*CategoriesAllResponse, error) {
	// @TODO: create wrapper?
	if err := s.Client.CheckTokens(); err != nil {
		return nil, err
	}

	apiURL, err := s.Client.GetApiURL(endpointResourceCategoriesAll)
	if err != nil {
		return nil, err
	}

	responseBody := &CategoriesAllResponse{}
	httpReq, err := s.Client.NewRequest(apiURL, requestBody)
	if err != nil {
		return nil, err
	}

	_, err = s.Client.Do(httpReq, responseBody)
	return responseBody, err
}

func (s *APIService) NewCategoriesAllRequest() *CategoriesAllRequest {
	return &CategoriesAllRequest{}
}

type CategoriesAllRequest struct {
	json.BaseRequest

	ResourceCategoryIDs []string                   `json:"ResourceCategoryIds,omitempty"` // Unique identifiers of Resource categories.
	ServiceIDs          []string                   `json:"ServiceIds,omitempty"`          // Unique identifiers of Services to which the resource categories belong.
	UpdatedUTC          configuration.TimeInterval `json:"UpdatedUtc,omitempty"`          // Interval in which the resource categories were updated.
	ActivityStates      ActivityStates             `json:"ActivityStates,omitempty"`      // Whether to return only active, only deleted or both records.
}

func (r CategoriesAllRequest) MarshalJSON() ([]byte, error) {
	return omitempty.MarshalJSON(r)
}

type ActivityStates []ActivityState

type ActivityState string

type CategoriesAllResponse struct {
	ResourceCategories ResourceCategories `json:"ResourceCategories"`
	Cursor             string             `json:"Cursor"` // Unique identifier of the last and hence oldest resource category returned. This can be used in Limitation in a subsequent request to fetch the next batch of older resource categories.
}

type ResourceCategories []ResourceCategory

type ResourceCategory struct {
	ID                 string                      `json:"Id"`           // Unique identifier of the category.
	EnterpriseID       string                      `json:"EnterpriseId"` // Unique identifier of the Enterprise.
	ServiceID          string                      `json:"ServiceId"`
	IsActive           bool                        `json:"IsActive"`           // Whether the category is still active.
	Type               ResourceCategoryType        `json:"Type"`               // Type of the category.
	Names              configuration.LocalizedText `json:"Names"`              // All translations of the name.
	ShortNames         configuration.LocalizedText `json:"ShortNames"`         // All translations of the short name.
	Descriptions       configuration.LocalizedText `json:"Descriptions"`       // All translations of the description.
	Ordering           int                         `json:"Ordering"`           // Ordering of the category, lower number corresponds to lower category (note that uniqueness nor continuous sequence is guaranteed).
	Capacity           int                         `json:"Capacity"`           // Capacity that can be served (e.g. bed count).
	ExtraCapacity      int                         `json:"ExtraCapacity"`      // Extra capacity that can be served (e.g. extra bed count).
	ExternalIdentifier string                      `json:"ExternalIdentifier"` // Identifier of the resource category from external system.
}

type ResourceCategoryType string

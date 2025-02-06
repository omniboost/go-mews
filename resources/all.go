package resources

import (
	"time"

	"github.com/omniboost/go-mews/configuration"
	"github.com/omniboost/go-mews/json"
	"github.com/omniboost/go-mews/omitempty"
)

const (
	endpointAll = "resources/getAll"
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
	json.BaseRequest

	ResourceIDs []string                   `json:"ResourceIds,omitempty"` // Unique identifiers of the requested Resources.
	CreatedUTC  configuration.TimeInterval `json:"CreatedUtc,omitempty"`  // Interval in which the Resources were created.
	UpdatedUTC  configuration.TimeInterval `json:"UpdatedUtc,omitempty"`  // Interval in which the Resources were updated.
	Extent      ResourceExtent             `json:"Extent,omitempty"`      // Extent of data to be returned.
}

func (r AllRequest) MarshalJSON() ([]byte, error) {
	return omitempty.MarshalJSON(r)
}

type AllResponse struct {
	Resources                   Resources                   `json:"Resources"`
}

type Resources []Resource

type Resource struct {
	ID               string        `json:"Id"`               // Unique identifier of the resource.
	IsActive         bool          `json:"IsActive"`         // Whether the resource is still active.
	Name             string        `json:"Name"`             // Name of the resource (e.g. room number).
	ParentResourceID string        `json:"ParentResourceId"` // Identifier of the parent Resource (e.g. room of a bed).
	State            ResourceState `json:"State"`            // State of the resource.
	Data             ResourceData  `json:"Data"`             // Additional data of the resource.
	CreatedUTC       time.Time     `json:"CreatedUtc"`       // Creation date and time of the resource in UTC timezone in ISO 8601 format.
	UpdatedUTC       time.Time     `json:"UpdatedUtc"`       // Last update date and time of the resource in UTC timezone in ISO 8601 format.
}

type ResourceExtent struct {
	Resources                        bool `json:"Resources"`                        // Whether the response should contain resources.
	ResourceCategories               bool `json:"ResourceCategories"`               // Whether the response should contain categories.
	ResourceCategoryAssignments      bool `json:"ResourceCategoryAssignments"`      // Whether the response should contain assignments of the resources to categories.
	ResourceCategoryImageAssignments bool `json:"ResourceCategoryImageAssignments"` // Whether the response should contain assignments of the images to categories.
	ResourceFeatures                 bool `json:"ResourceFeatures"`                 // Whether the response should contain resource features.
	ResourceFeatureAssignments       bool `json:"ResourceFeatureAssignments"`       // Whether the response should contain assignments of the resources to features.
	Inactive                         bool `json:"Inactive"`                         // Whether the response should contain inactive entities.
}

type ResourceState string

type ResourceData struct {
	Discriminator ResourceDataDiscriminator `json:"Discriminator"` // If resource is space, object or person.
	Value         interface{}               `json:"Value"`         // Based on Resource data discriminator, e.g. Space resource data
}

type ResourceDataDiscriminator string

package resources

import (
	"time"

	"github.com/omniboost/go-mews/configuration"
	"github.com/omniboost/go-mews/json"
	base "github.com/omniboost/go-mews/json"
	"github.com/omniboost/go-mews/omitempty"
)

const (
	endpointBlocksAll = "resourceBlocks/getAll"
)

// List all products
func (s *APIService) BlocksAll(requestBody *BlocksAllRequest) (*BlocksAllResponse, error) {
	// @TODO: create wrapper?
	if err := s.Client.CheckTokens(); err != nil {
		return nil, err
	}

	apiURL, err := s.Client.GetApiURL(endpointBlocksAll)
	if err != nil {
		return nil, err
	}

	responseBody := &BlocksAllResponse{}
	httpReq, err := s.Client.NewRequest(apiURL, requestBody)
	if err != nil {
		return nil, err
	}

	_, err = s.Client.Do(httpReq, responseBody)
	return responseBody, err
}

func (s *APIService) NewBlocksAllRequest() *BlocksAllRequest {
	return &BlocksAllRequest{}
}

type BlocksAllRequest struct {
	json.BaseRequest
	Limitation base.Limitation `json:"Limitation,omitempty"`

	ResourceBlockIDs    []string                   `json:"ResourceBlockIds,omitempty"`    // Unique identifiers of the requested Resource blocks.
	AssignedResourceIDs []string                   `json:"AssignedResourceIds,omitempty"` // Unique identifiers of the requested Assigned Resources.
	CollidingUTC        configuration.TimeInterval `json:"CollidingUtc,omitempty"`        // Interval in which the Resource block is active.
	CreatedUTC          configuration.TimeInterval `json:"CreatedUtc,omitempty"`          // Interval in which the Resource block was created.
	UpdatedUTC          configuration.TimeInterval `json:"UpdatedUtc,omitempty"`          // Interval in which the Resource block was updated.
	Extent              ResourceBlockExtent        `json:"Extent,omitempty"`              // Extent of data to be returned.
}

func (r BlocksAllRequest) MarshalJSON() ([]byte, error) {
	return omitempty.MarshalJSON(r)
}

type BlocksAllResponse struct {
	ResourceBlocks ResourceBlocks `json:"ResourceBlocks"`
	Cursor         string         `json:"Cursor"`
}

type ResourceBlocks []ResourceBlock

type ResourceBlock struct {
	ID                 string            `json:"Id"`                 // Unique identifier of the block.
	AssignedResourceID string            `json:"AssignedResourceId"` // Unique identifier of the assigned Resource.
	IsActive           bool              `json:"IsActive"`           // Whether the block is still active.
	Type               ResourceBlockType `json:"Type"`               // Type of the resource block.
	StartUTC           time.Time         `json:"StartUtc"`           // Start of the block in UTC timezone in ISO 8601 format.
	EndUTC             time.Time         `json:"EndUtc"`             // End of the block in UTC timezone in ISO 8601 format.
	CreatedUTC         time.Time         `json:"CreatedUtc"`         // Creation date and time of the block in UTC timezone in ISO 8601 format.
	UpdatedUTC         time.Time         `json:"UpdatedUtc"`         // Last update date and time of the block in UTC timezone in ISO 8601 format.
}

type ResourceBlockExtent struct {
	Inactive bool `json:"Inactive"` // Whether the response should contain inactive entities.
}

type ResourceBlockType string

const (
	ResourceBlockTypeOutOfOrder  ResourceBlockType = "OutOfOrder"
	ResourceBlockTypeInternalUse ResourceBlockType = "InternalUse"
)

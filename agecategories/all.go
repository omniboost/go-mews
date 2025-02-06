package agecategories

import (
	"time"

	"github.com/omniboost/go-mews/configuration"
	base "github.com/omniboost/go-mews/json"
	"github.com/omniboost/go-mews/omitempty"
)

const (
	endpointAll = "ageCategories/getAll"
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

	EnterpriseIDs  []string                   `json:"EnterpriseIDs,omitempty"`
	ServiceIDs     []string                   `json:"ServiceIds,omitempty"`
	AgeCategoryIDs []string                   `json:"AgeCategoryIds,omitempty"`
	UpdatedUTC     configuration.TimeInterval `json:"UpdatedUtc,omitempty"`
	ActivityStates ActivityStates             `json:"ActivityStates,omitempty"`
	Limitation     base.Limitation            `json:"Limitation,omitempty"`
}

func (r AllRequest) MarshalJSON() ([]byte, error) {
	return omitempty.MarshalJSON(r)
}

type AllResponse struct {
	AgeCategories AgeCategories `json:"AgeCategories"`
	Cursor        string        `json:"Cursor"`
}

type ActivityStates []ActivityState

type ActivityState string

type AgeCategories []AgeCategory

type AgeCategory struct {
	ID                 string                    `json:"Id"`
	ServiceID          string                    `json:"ServiceId"`
	MinimalAge         int                       `json:"MinimalAge"`
	MaximalAge         int                       `json:"MaximalAge"`
	Names              map[string]string         `json:"Names"`
	ShortNames         map[string]string         `json:"ShortNames"`
	CreatedUTC         time.Time                 `json:"CreatedUtc"`
	UpdatedUTC         time.Time                 `json:"UpdatedUtc"`
	Classification     AgeCategoryClassification `json:"Classification"`
	IsActive           bool                      `json:"IsActive"`
	ExternalIdentifier string                    `json:"ExternalIdentifier"`
}

type AgeCategoryClassification string

var (
	AgeCategoryClassificationAdult AgeCategoryClassification = "Adult"
	AgeCategoryClassificationChild AgeCategoryClassification = "Child"
)

package products

import "github.com/omniboost/go-mews/json"

type APIService struct {
	Client *json.Client
}

func NewAPIService() *APIService {
	return &APIService{}
}

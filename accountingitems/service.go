package accountingitems

import "github.com/omniboost/go-mews/json"

type APIService struct {
	Client *json.Client
}

func NewService() *APIService {
	return &APIService{}
}

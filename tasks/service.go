package tasks

import "github.com/omniboost/go-mews/json"

type Service struct {
	Client *json.Client
}

func NewService() *Service {
	return &Service{}
}

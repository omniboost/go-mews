package commands

import (
	"errors"
	"fmt"

	"github.com/omniboost/go-mews/json"
)

type Service struct {
	Client *json.Client
}

func NewService() *Service {
	return &Service{}
}

func StateFromString(s string) (CommandState, error) {
	switch s {
	case "Pending":
		return CommandStatePending, nil
	case "Received":
		return CommandStateReceived, nil
	case "Processing":
		return CommandStateProcessing, nil
	case "Processed":
		return CommandStateProcessed, nil
	case "Cancelled":
		return CommandStateCancelled, nil
	case "Error":
		return CommandStateError, nil
	default:
		return "", errors.New(fmt.Sprintf("Can't convert %s to CommandState", s))
	}

}

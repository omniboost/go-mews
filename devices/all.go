package devices

import (
	"github.com/omniboost/go-mews/json"
	"github.com/omniboost/go-mews/omitempty"
)

const (
	endpointAll = "devices/getAll"
)

// List all commands
func (s *Service) All(requestBody *AllRequest) (*AllResponse, error) {
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
	json.BaseRequest
	Limitation json.Limitation `json:"Limitation,omitempty"`
}

func (r AllRequest) MarshalJSON() ([]byte, error) {
	return omitempty.MarshalJSON(r)
}

type AllResponse struct {
	Devices Devices `json:"Devices"`
	Cursor  string  `json:"Cursor"`
}

type Devices []Device

type Device struct {
	ID   string     `json:"Id"`   // Unique identifier of the Device to which the command is send
	Name string     `json:"Name"` // Name of the Device to which the command is send
	Type DeviceType `json:"Type"` //Type of Device
}

type DeviceType string

const (
	DevicePrinter         DeviceType = "Printer"
	DevicePaymentTerminal DeviceType = "PaymentTerminal"
	DevicePassportScanner DeviceType = "PassportScanner"
	DeviceFiscalMachine   DeviceType = "FiscalMachine"
	DeviceKeyCutter       DeviceType = "KeyCutter"
	DeviceVisiKeyCutter   DeviceType = "VisiOnlineKeyCutter"
)

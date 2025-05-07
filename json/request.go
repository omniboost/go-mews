package json

import (
	"context"
	"github.com/cydev/zero"
)

type BaseRequest struct {
	AccessToken  string `json:"AccessToken"`
	ClientToken  string `json:"ClientToken,omitempty"`
	LanguageCode string `json:"LanguageCode,omitempty"`
	CultureCode  string `json:"CultureCode,omitempty"`
	Client       string `json:"Client,omitempty"`
	context      context.Context
}

func (req *BaseRequest) SetContext(ctx context.Context) {
	req.context = ctx
}

func (req *BaseRequest) GetContext() context.Context {
	if req.context == nil {
		return context.Background()
	}
	return req.context
}

func (req *BaseRequest) SetAccessToken(token string) {
	req.AccessToken = token
}

func (req *BaseRequest) SetClientToken(token string) {
	req.ClientToken = token
}

func (req *BaseRequest) SetLanguageCode(code string) {
	req.LanguageCode = code
}

func (req *BaseRequest) SetCultureCode(code string) {
	req.CultureCode = code
}

type Limitation struct {
	Cursor string `json:"Cursor,omitempty"`
	Count  int    `json:"Count,omitempty"`
}

func (l Limitation) IsEmpty() bool {
	return zero.IsZero(l)
}

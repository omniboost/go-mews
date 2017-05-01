package mews

import (
	"net/http"
	"net/url"

	"github.com/tim-online/go-mews/accountingcategories"
	"github.com/tim-online/go-mews/accountingitems"
	"github.com/tim-online/go-mews/json"
)

const (
	libraryVersion = "0.0.1"
	userAgent      = "go-mews/" + libraryVersion
)

var (
	BaseURL = &url.URL{
		Scheme: "https",
		Host:   "www.mews.li",
		Path:   "/api/connector/v1/",
	}
	BaseURLDemo = &url.URL{
		Scheme: "https",
		Host:   "demo.mews.li",
		Path:   "/api/connector/v1/",
	}
)

// NewClient returns a new Nmbrs API client
func NewClient(httpClient *http.Client, token string) *Client {
	if httpClient == nil {
		httpClient = http.DefaultClient
	}

	jsonClient := json.NewClient(httpClient, token)
	jsonClient.UserAgent = userAgent
	jsonClient.Token = token
	jsonClient.Debug = false

	c := &Client{
		client:  jsonClient,
	}

	c.SetBaseURL(BaseURL)

	// Services
	c.AccountingItems = accountingitems.NewService()
	c.AccountingItems.Client = c.client
	c.AccountingCategories = accountingcategories.NewService()
	c.AccountingCategories.Client = c.client

	return c
}

// Client manages communication with Nmbrs API
type Client struct {
	// HTTP client used to communicate with the API.
	client *json.Client

	// Services used for communicating with the API
	AccountingItems      *accountingitems.Service
	AccountingCategories *accountingcategories.Service
}

func (c *Client) SetDebug(debug bool) {
	c.client.Debug = debug
}

func (c *Client) SetBaseURL(baseURL *url.URL) {
	c.client.BaseURL = baseURL
}

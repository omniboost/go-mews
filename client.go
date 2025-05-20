package mews

import (
	"context"
	"net/http"
	"net/url"

	"github.com/omniboost/go-mews/accountingcategories"
	"github.com/omniboost/go-mews/accountingitems"
	"github.com/omniboost/go-mews/agecategories"
	"github.com/omniboost/go-mews/bills"
	"github.com/omniboost/go-mews/businesssegments"
	"github.com/omniboost/go-mews/cashiertransactions"
	"github.com/omniboost/go-mews/commands"
	"github.com/omniboost/go-mews/companies"
	"github.com/omniboost/go-mews/companionships"
	"github.com/omniboost/go-mews/configuration"
	"github.com/omniboost/go-mews/counters"
	"github.com/omniboost/go-mews/countries"
	"github.com/omniboost/go-mews/creditcards"
	"github.com/omniboost/go-mews/customers"
	"github.com/omniboost/go-mews/enterprises"
	"github.com/omniboost/go-mews/finance"
	"github.com/omniboost/go-mews/json"
	"github.com/omniboost/go-mews/ledgerbalances"
	"github.com/omniboost/go-mews/ledgerentries"
	"github.com/omniboost/go-mews/orderitems"
	"github.com/omniboost/go-mews/outletitems"
	"github.com/omniboost/go-mews/outlets"
	"github.com/omniboost/go-mews/payments"
	"github.com/omniboost/go-mews/products"
	"github.com/omniboost/go-mews/productserviceorders"
	"github.com/omniboost/go-mews/rates"
	"github.com/omniboost/go-mews/reservationgroups"
	"github.com/omniboost/go-mews/reservations"
	"github.com/omniboost/go-mews/resources"
	"github.com/omniboost/go-mews/services"
	"github.com/omniboost/go-mews/tasks"
)

const (
	libraryVersion = "0.0.1"
	userAgent      = "go-mews/" + libraryVersion
)

var (
	BaseURL = &url.URL{
		Scheme: "https",
		Host:   "api.mews.com",
		Path:   "/api/connector/v1/",
	}
	BaseURLDemo = &url.URL{
		Scheme: "https",
		Host:   "api.mews-demo.com",
		Path:   "/api/connector/v1/",
	}
)

// NewClient returns a new MEWS API client
func NewClient(httpClient *http.Client, accessToken string, clientToken string) *Client {
	if httpClient == nil {
		httpClient = http.DefaultClient
	}

	jsonClient := json.NewClient(httpClient, accessToken, clientToken)
	jsonClient.UserAgent = userAgent
	jsonClient.AccessToken = accessToken
	jsonClient.ClientToken = clientToken
	jsonClient.Debug = false

	c := &Client{
		client: jsonClient,
	}

	c.SetBaseURL(BaseURL)

	// Services
	c.AccountingItems = accountingitems.NewService()
	c.AccountingItems.Client = c.client
	c.AgeCategories = agecategories.NewService()
	c.AgeCategories.Client = c.client
	c.Payments = payments.NewService()
	c.Payments.Client = c.client
	c.OrderItems = orderitems.NewService()
	c.OrderItems.Client = c.client
	c.OutletItems = outletitems.NewService()
	c.OutletItems.Client = c.client
	c.AccountingCategories = accountingcategories.NewService()
	c.AccountingCategories.Client = c.client
	c.Companies = companies.NewService()
	c.Companies.Client = c.client
	c.Companionships = companionships.NewService()
	c.Companionships.Client = c.client
	c.Counters = counters.NewService()
	c.Counters.Client = c.client
	c.Countries = countries.NewService()
	c.Countries.Client = c.client
	c.Customers = customers.NewService()
	c.Customers.Client = c.client
	c.Outlets = outlets.NewAPIService()
	c.Outlets.Client = c.client
	c.Enterprises = enterprises.NewAPIService()
	c.Enterprises.Client = c.client
	c.Products = products.NewAPIService()
	c.Products.Client = c.client
	c.Reservations = reservations.NewAPIService()
	c.Reservations.Client = c.client
	c.ReservationGroups = reservationgroups.NewAPIService()
	c.ReservationGroups.Client = c.client
	c.Resources = resources.NewAPIService()
	c.Resources.Client = c.client
	c.ProductServiceOrders = productserviceorders.NewAPIService()
	c.ProductServiceOrders.Client = c.client
	c.Services = services.NewAPIService()
	c.Services.Client = c.client
	c.Rates = rates.NewAPIService()
	c.Rates.Client = c.client
	c.Bills = bills.NewService()
	c.Bills.Client = c.client
	c.Commands = commands.NewService()
	c.Commands.Client = c.client
	c.Configuration = configuration.NewService()
	c.Configuration.Client = c.client
	c.BusinessSegments = businesssegments.NewService()
	c.BusinessSegments.Client = c.client
	c.Tasks = tasks.NewService()
	c.Tasks.Client = c.client
	c.Finance = finance.NewService()
	c.Finance.Client = c.client
	c.CashierTransactions = cashiertransactions.NewService()
	c.CashierTransactions.Client = c.client
	c.CreditCards = creditcards.NewService()
	c.CreditCards.Client = c.client
	c.LedgerBalances = ledgerbalances.NewService()
	c.LedgerBalances.Client = c.client
	c.LedgerEntries = ledgerentries.NewService()
	c.LedgerEntries.Client = c.client

	return c
}

// Client manages communication with MEWS API
type Client struct {
	// HTTP client used to communicate with the API.
	client *json.Client

	// Services used for communicating with the API
	AccountingItems      *accountingitems.APIService
	AgeCategories        *agecategories.Service
	Payments             *payments.Service
	OrderItems           *orderitems.Service
	OutletItems          *outletitems.Service
	AccountingCategories *accountingcategories.Service
	Companies            *companies.Service
	Companionships       *companionships.Service
	Counters             *counters.Service
	Countries            *countries.Service
	Customers            *customers.Service
	Outlets              *outlets.APIService
	Enterprises          *enterprises.APIService
	Products             *products.APIService
	Reservations         *reservations.APIService
	ReservationGroups    *reservationgroups.APIService
	Resources            *resources.APIService
	ProductServiceOrders *productserviceorders.APIService
	Services             *services.APIService
	Rates                *rates.APIService
	Bills                *bills.Service
	Commands             *commands.Service
	Configuration        *configuration.Service
	BusinessSegments     *businesssegments.Service
	Tasks                *tasks.Service
	Finance              *finance.Service
	CashierTransactions  *cashiertransactions.Service
	CreditCards          *creditcards.Service
	LedgerBalances       *ledgerbalances.Service
	LedgerEntries        *ledgerentries.Service
}

func (c *Client) SetDebug(debug bool) {
	c.client.Debug = debug
}

func (c *Client) SetBaseURL(baseURL *url.URL) {
	c.client.BaseURL = baseURL
}

func (c *Client) SetDisallowUnknownFields(disallowUnknownFields bool) {
	c.client.DisallowUnknownFields = disallowUnknownFields
}

func (c *Client) SetLanguageCode(code string) {
	c.client.SetLanguageCode(code)
}

func (c *Client) SetCultureCode(code string) {
	c.client.SetCultureCode(code)
}

func (c *Client) GetWebsocket(ctx context.Context) *Websocket {
	ws := NewWebsocket(c.client.Client, c.client.AccessToken, c.client.ClientToken)
	url := &url.URL{
		Scheme: WebsocketURL.Scheme,
		Host:   WebsocketURL.Host,
		Path:   WebsocketURL.Path,
	}
	ws.SetBaseURL(url)
	ws.SetDebug(c.client.Debug)
	return ws
}

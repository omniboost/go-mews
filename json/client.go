package json

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"log"
	"net/http"
	"net/http/httputil"
	"net/url"
	"time"
)

var (
	defaultUserAgent = "go"
	mediaType        = "application/json"
	charset          = "utf-8"
	defaultTimeout   = 10 * time.Second

	ErrNoAccessToken = errors.New("No access token specified")
	ErrNoClientToken = errors.New("No client token specified")
)

type Client struct {
	// HTTP client used to communicate with the DO API.
	Client *http.Client

	// Base URL for API requests
	BaseURL *url.URL

	// Debugging flag
	Debug bool

	// Disallow unknown json fields
	DisallowUnknownFields bool

	// User agent for client
	UserAgent string

	AccessToken string
	ClientToken string

	languageCode string
	cultureCode  string

	// Optional function called after every successful request made to the DO APIs
	onRequestCompleted RequestCompletionCallback
}

// RequestCompletionCallback defines the type of the request callback function
type RequestCompletionCallback func(*http.Request, *http.Response)

func NewClient(httpClient *http.Client, accessToken string, clientToken string) *Client {
	c := &Client{
		Client:      nil,
		UserAgent:   defaultUserAgent,
		AccessToken: accessToken,
		ClientToken: clientToken,
	}

	if httpClient == nil {
		c.Client = http.DefaultClient
		c.Client.Timeout = defaultTimeout
	} else {
		c.Client = httpClient
	}

	return c
}

func (c *Client) GetApiURL(path string) (*url.URL, error) {
	apiURL, err := url.Parse(c.BaseURL.String())
	if err != nil {
		return nil, err
	}

	apiURL.Path = apiURL.Path + path
	return apiURL, nil
}

// Do sends an API request and returns the API response. The API response is XML decoded and stored in the value
// pointed to by v, or returned as an error if an API error has occurred. If v implements the io.Writer interface,
// the raw response will be written to v, without attempting to decode it.
func (c *Client) Do(req *http.Request, response interface{}) (*http.Response, error) {
	if c.Debug == true {
		dump, _ := httputil.DumpRequestOut(req, true)
		log.Println(string(dump))
	}

	httpResp, err := c.Client.Do(req)
	if err != nil {
		return nil, err
	}
	if c.onRequestCompleted != nil {
		c.onRequestCompleted(req, httpResp)
	}

	// close body io.Reader
	defer func() {
		if rerr := httpResp.Body.Close(); err == nil {
			err = rerr
		}
	}()

	if c.Debug == true {
		dump, _ := httputil.DumpResponse(httpResp, true)
		log.Println(string(dump))
	}

	// check if the response isn't an error
	err = CheckResponse(httpResp)
	if err != nil {
		return httpResp, err
	}

	// check the provided interface parameter
	if response == nil {
		return httpResp, err
	}

	// interface implements io.Writer: write Body to it
	if w, ok := response.(io.Writer); ok {
		_, err := io.Copy(w, httpResp.Body)
		return httpResp, err
	}

	// try to decode body into interface parameter
	dec := json.NewDecoder(httpResp.Body)
	if c.DisallowUnknownFields {
		dec.DisallowUnknownFields()
	}
	err = dec.Decode(response)
	if err != nil {
		return nil, err
	}
	return httpResp, err
}

// NewRequest creates an API request. A relative URL can be provided in urlStr, which will be resolved to the
// BaseURL of the Client. Relative URLS should always be specified without a preceding slash. If specified, the
// value pointed to by body is XML encoded and included in as the request body.
func (c *Client) NewRequest(apiURL *url.URL, requestBody interface{}) (*http.Request, error) {
	buf := new(bytes.Buffer)
	if requestBody != nil {
		if s, ok := requestBody.(RequestBody); ok {
			s.SetAccessToken(c.AccessToken)
			s.SetClientToken(c.ClientToken)
			if c.languageCode != "" {
				s.SetLanguageCode(c.languageCode)
			}
			if c.cultureCode != "" {
				s.SetCultureCode(c.cultureCode)
			}
		}

		err := json.NewEncoder(buf).Encode(requestBody)
		if err != nil {
			return nil, err
		}
	}

	httpReq, err := http.NewRequest(http.MethodPost, apiURL.String(), buf)
	if err != nil {
		return nil, err
	}

	httpReq.Header.Add("Content-Type", fmt.Sprintf("%s; charset=%s", mediaType, charset))
	httpReq.Header.Add("Accept", mediaType)
	httpReq.Header.Add("User-Agent", c.UserAgent)
	return httpReq, nil
}

// OnRequestCompleted sets the DO API request completion callback
func (c *Client) OnRequestCompleted(rc RequestCompletionCallback) {
	c.onRequestCompleted = rc
}

func (c *Client) CheckTokens() error {
	if c.AccessToken == "" {
		return ErrNoAccessToken
	}

	if c.ClientToken == "" {
		return ErrNoClientToken
	}

	return nil
}

func (c *Client) SetLanguageCode(code string) {
	c.languageCode = code
}

func (c *Client) SetCultureCode(code string) {
	c.cultureCode = code
}

type RequestBody interface {
	SetAccessToken(string)
	SetClientToken(string)
	SetLanguageCode(string)
	SetCultureCode(string)
}

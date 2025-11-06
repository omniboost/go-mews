package json

import (
	"bytes"
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"log"
	"net/http"
	"net/http/httputil"
	"net/url"
	"time"

	httperr "github.com/omniboost/go-httperr"
)

var (
	defaultUserAgent = "go"
	mediaType        = "application/json"
	charset          = "utf-8"

	ErrNoAccessToken = errors.New("No access token specified")
	ErrNoClientToken = errors.New("No client token specified")
	ctxRetryAttempt  = ContextKey("retry-attempt")
)

type ContextKey string

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

	Timeout        time.Duration
	RetryOnTimeout bool
	MaxRetries     int

	// 429 - Too many requests handling
	retryAfter *time.Time
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

func cloneRequest(req *http.Request, ctx context.Context) (*http.Request, error) {
	newReq := req.Clone(ctx)
	if req.Body != nil && req.Body != http.NoBody {
		// If the request has a body, we need to read it and set it again
		body, err := io.ReadAll(req.Body)
		if err != nil {
			return nil, err
		}
		req.Body = io.NopCloser(bytes.NewReader(body))
		newReq.Body = io.NopCloser(bytes.NewReader(body))
	}
	return newReq, nil
}

// Do sends an API request and returns the API response. The API response is XML decoded and stored in the value
// pointed to by v, or returned as an error if an API error has occurred. If v implements the io.Writer interface,
// the raw response will be written to v, without attempting to decode it.
func (c *Client) Do(req *http.Request, response interface{}) (*http.Response, error) {
	if c.Debug == true {
		dump, _ := httputil.DumpRequestOut(req, true)
		log.Println(string(dump))
	}

	// Wait until "Retry-After" time has passed
	c.sleepUntilRetryAfter()

	allowRetry := c.RetryOnTimeout
	originalContext := req.Context()
	retryAttempt, ok := originalContext.Value(ctxRetryAttempt).(int)
	if !ok {
		retryAttempt = 0
	}
	// if the request context doesn't have a deadline set, and we have a default deadline, set a timeout
	if _, ok := originalContext.Deadline(); !ok && c.Timeout > 0 {
		ctx, cancel := context.WithTimeout(originalContext, c.Timeout)
		defer cancel()
		req = req.WithContext(ctx)
		// only allow retry if we haven't reached the maximum number of retries
		allowRetry = retryAttempt < c.MaxRetries
	} else {
		// if the request context already has a deadline set,
		// retry on timeout is useless, since the deadline will already be in the past
		allowRetry = false
	}

	var originalReq *http.Request
	if allowRetry {
		var err error
		originalReq, err = cloneRequest(req, context.WithValue(originalContext, ctxRetryAttempt, retryAttempt+1))
		if err != nil {
			return nil, fmt.Errorf("failed to clone request: %w", err)
		}
	}

	httpResp, err := c.Client.Do(req)
	if err != nil {
		if errors.Is(err, context.DeadlineExceeded) && allowRetry {
			// if the request timed out, retry it
			if c.Debug == true {
				log.Println("Request timed out, retrying...")
			}
			time.Sleep(500 * time.Millisecond)
			return c.Do(originalReq, response)
		}

		// wrap error in http error so we can handle it properly
		return nil, &httperr.Error{StatusCode: httpResp.StatusCode, Err: err}
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

	// Handle '429 - Too many requests' responses
	if httpResp.StatusCode == http.StatusTooManyRequests {
		c.SetRetryAfterByResponse(httpResp)
		c.sleepUntilRetryAfter()
		c.SetRetryAfter(nil)
		return c.Do(originalReq, response)
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
	ctx := context.Background()
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
			ctx = s.GetContext()
		}

		err := json.NewEncoder(buf).Encode(requestBody)
		if err != nil {
			return nil, err
		}
	}

	httpReq, err := http.NewRequestWithContext(ctx, http.MethodPost, apiURL.String(), buf)
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

func (c *Client) SetRetryAfter(retryAfter *time.Time) {
	// Set the "Retry-After" time
	c.retryAfter = retryAfter
}

func (c *Client) SetRetryAfterByResponse(req *http.Response) error {
	// Get the "Retry-After" header
	retryAfter := req.Header.Get("Retry-After")

	// When the "Retry-After" header is not set, continue
	if retryAfter == "" {
		return nil
	}

	// Parse the "Retry-After" header to a duration in seconds
	retryAfterTime, err := time.Parse(time.RFC1123, retryAfter)
	if err != nil {
		return err
	}

	// Set the "Retry-After" time
	c.SetRetryAfter(&retryAfterTime)

	return nil
}

func (c *Client) sleepUntilRetryAfter() {
	// When the "Retry-After" time is not set, continue
	if c.retryAfter == nil || c.retryAfter.IsZero() {
		return
	}

	// Calculate the duration to sleep
	now := time.Now()
	diff := c.retryAfter.Sub(now)
	if diff <= 0 {
		return
	}

	// Sleep for the duration of "Retry-After"
	time.Sleep(diff)
}

type RequestBody interface {
	SetAccessToken(string)
	SetClientToken(string)
	SetLanguageCode(string)
	SetCultureCode(string)
	GetContext() context.Context
}

package etherscan

import (
	"context"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"

	"golang.org/x/time/rate"
)

const baseURL = "https://api.etherscan.io/api"
const baseURLPolygon = "https://api.polygonscan.com/api"

// 5 calls each second per IP
var rateLimiter = rate.NewLimiter(rate.Every(5), 1)

// Client struct
type Client struct {
	httpClient *http.Client
	apiKey     string
}

// NewClient create new client object
func NewClient(httpClient *http.Client, apiKey string) *Client {
	if httpClient == nil {
		httpClient = http.DefaultClient
	}

	return &Client{
		httpClient: httpClient,
		apiKey:     apiKey,
	}
}

func (c *Client) doReq(req *http.Request) ([]byte, error) {
	if err := rateLimiter.Wait(context.Background()); err != nil {
		return nil, err
	}

	resp, err := c.httpClient.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("%s", body)
	}

	var envelope Envelope
	if err := jsonex.Unmarshal(body, &envelope); err != nil {
		return nil, err
	}

	if envelope.Status != 1 && envelope.Message != "No transactions found" {
		return nil, errors.New(envelope.Message)
	}

	return envelope.Result, nil
}

// MakeReq HTTP request helper
func (c *Client) MakeReq(url string) ([]byte, error) {
	req, err := http.NewRequest(http.MethodGet, url, nil)
	if err != nil {
		return nil, err
	}

	resp, err := c.doReq(req)
	if err != nil {
		return nil, err
	}

	return resp, err
}

// NewURL new url
func (c *Client) NewURL(module, action string, params Params) string {
	return fmt.Sprintf("%s?module=%s&action=%s&%s&apikey=%s", baseURL, module, action, params.Encode(), c.apiKey)
}

// NewURL new url for polygon
func (c *Client) NewURLPolygon(module, action string, params Params) string {
	return fmt.Sprintf("%s?module=%s&action=%s&%s&apikey=%s", baseURLPolygon, module, action, params.Encode(), c.apiKey)
}

// Params struct
type Params map[string]interface{}

// Encode encode to query string
func (p Params) Encode() string {
	values := url.Values{}

	for k, v := range p {
		values.Set(k, fmt.Sprint(v))
	}

	return values.Encode()
}

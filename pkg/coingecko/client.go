package coingecko

import (
	"context"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
	"strconv"
	"time"

	"golang.org/x/time/rate"
)

const baseURL = "https://api.coingecko.com/api/v3"

// 100 calls each minute per IP
var rateLimiter = rate.NewLimiter(rate.Every(99*time.Second/60), 1)

// Client struct
type Client struct {
	httpClient *http.Client
}

// NewClient create new client object
func NewClient(httpClient *http.Client) *Client {
	if httpClient == nil {
		httpClient = http.DefaultClient
	}

	return &Client{
		httpClient: httpClient,
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

	return body, nil
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

// CoinsIDMarketChartRange /coins/{id}/market_chart?vs_currency={usd, eur, jpy, etc.}&from={eq. 1392577232}&to={eg. 1422577232}
func (c *Client) CoinsIDMarketChartRange(id string, vsCurrency string, from uint64, to uint64) (*CoinsIDMarketChart, error) {
	params := url.Values{}
	params.Add("vs_currency", vsCurrency)
	params.Add("from", strconv.FormatUint(from, 10))
	params.Add("to", strconv.FormatUint(to, 10))

	url := fmt.Sprintf("%s/coins/%s/market_chart/range?%s", baseURL, id, params.Encode())
	resp, err := c.MakeReq(url)
	if err != nil {
		log.Printf("failed to get coingecko id: %s", id)
		return nil, err
	}

	marketChart := CoinsIDMarketChart{}
	if err := jsonex.Unmarshal(resp, &marketChart); err != nil {
		return nil, err
	}

	return &marketChart, nil
}

// CoinsContractMarketChartRange /coins/{id}/contract/{contract_address}/market_chart/range?vs_currency={usd, eur, jpy, etc.}&from={eq. 1392577232}&to={eg. 1422577232}
func (c *Client) CoinsContractMarketChartRange(contractAddress string, vsCurrency string, from uint64, to uint64) (*CoinsIDMarketChart, error) {
	params := url.Values{}
	params.Add("vs_currency", vsCurrency)
	params.Add("from", strconv.FormatUint(from, 10))
	params.Add("to", strconv.FormatUint(to, 10))

	url := fmt.Sprintf("%s/coins/ethereum/contract/%s/market_chart/range?%s", baseURL, contractAddress, params.Encode())
	resp, err := c.MakeReq(url)
	if err != nil {
		return nil, err
	}

	marketChart := CoinsIDMarketChart{}
	if err := jsonex.Unmarshal(resp, &marketChart); err != nil {
		return nil, err
	}

	return &marketChart, nil
}

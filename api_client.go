package openaiorgs

import (
	"encoding/json"
	"fmt"

	"github.com/go-resty/resty/v2"
)

const DefaultBaseURL = "https://api.openai.com/v1"

type HTTPRequester interface {
	R() *resty.Request
}

type Client struct {
	httpClient HTTPRequester
	BaseURL    string
}

type ListResponse[T any] struct {
	Object  string `json:"object"`
	Data    []T    `json:"data"`
	FirstID string `json:"first_id"`
	LastID  string `json:"last_id"`
	HasMore bool   `json:"has_more"`
}

func NewClient(baseURL, token string) *Client {
	client := resty.New()
	if baseURL == "" {
		baseURL = DefaultBaseURL
	}
	client.SetBaseURL(baseURL)
	client.SetAuthToken(token)
	client.SetHeader("Content-Type", "application/json")

	return &Client{
		httpClient: client,
		BaseURL:    baseURL,
	}
}

func GetSingle[T any](client HTTPRequester, endpoint string) (*T, error) {
	resp, err := client.R().
		ExpectContentType("application/json").
		Get(endpoint)
	if err != nil {
		return nil, fmt.Errorf("error making GET request: %v", err)
	}

	if resp.IsError() {
		return nil, fmt.Errorf("API request failed with status code %d: %s", resp.StatusCode(), string(resp.Body()))
	}

	var result T
	err = json.Unmarshal(resp.Body(), &result)
	if err != nil {
		return nil, fmt.Errorf("error unmarshaling response: %v", err)
	}

	return &result, nil
}

func Get[T any](client HTTPRequester, endpoint string, queryParams map[string]string) (*ListResponse[T], error) {
	resp, err := client.R().
		SetQueryParams(queryParams).
		ExpectContentType("application/json").
		Get(endpoint)
	if err != nil {
		return nil, fmt.Errorf("error making GET request: %v", err)
	}

	if resp.IsError() {
		return nil, fmt.Errorf("API request failed with status code %d: %s", resp.StatusCode(), string(resp.Body()))
	}

	var listResp ListResponse[T]
	err = json.Unmarshal(resp.Body(), &listResp)
	if err != nil {
		return nil, fmt.Errorf("error unmarshaling response: %v", err)
	}

	return &listResp, nil
}

func Post[T any](client HTTPRequester, endpoint string, body interface{}) (*T, error) {
	resp, err := client.R().
		SetBody(body).
		ExpectContentType("application/json").
		Post(endpoint)
	if err != nil {
		return nil, fmt.Errorf("error making POST request: %v", err)
	}

	if resp.IsError() {
		return nil, fmt.Errorf("API request failed with status code %d: %s", resp.StatusCode(), string(resp.Body()))
	}

	var result T
	err = json.Unmarshal(resp.Body(), &result)
	if err != nil {
		return nil, fmt.Errorf("error unmarshaling response: %v", err)
	}

	return &result, nil
}

func Delete[T any](client HTTPRequester, endpoint string) error {
	_, err := client.R().
		Delete(endpoint)
	if err != nil {
		return fmt.Errorf("error making DELETE request: %v", err)
	}

	return nil
}

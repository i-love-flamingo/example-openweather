package openweather

import (
	"context"
	"errors"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"net/url"
	"strings"

	"flamingo.me/flamingo/v3/framework/flamingo"
)

type (
	// APIClient for openweather
	APIClient struct {
		baseURL    string
		apiKey     string
		httpClient *http.Client
		logger     flamingo.Logger
	}
)

// Inject dependencies
func (c *APIClient) Inject(
	httpClient *http.Client,
	logger flamingo.Logger,
	cfg *struct {
		BaseURL string `inject:"config:openweather.apiurl"`
		APIKey  string `inject:"config:openweather.apikey"`
	},
) *APIClient {
	c.httpClient = httpClient
	c.logger = logger
	if cfg != nil {
		c.baseURL = cfg.BaseURL
		c.apiKey = cfg.APIKey
	}

	return c
}

func (c *APIClient) request(ctx context.Context, method, path string, body io.Reader) (*http.Response, error) {
	path = strings.TrimRight(c.baseURL, "/") + "/" + strings.TrimLeft(path, "/")

	c.logger.Debugf("openweather.apiClient.request: method: %v path: %v", method, path)

	u, err := url.Parse(path)
	if err != nil {
		return nil, err
	}

	query := u.Query()
	query.Add("appid", c.apiKey)
	query.Add("units", "metric")

	u.RawQuery = query.Encode()

	c.logger.Info("Requesting", u.String())

	request, err := http.NewRequest(method, u.String(), body)
	if err != nil {
		return nil, err
	}
	request.Header.Set("Content-Type", "application/json")

	response, err := c.httpClient.Do(request.WithContext(ctx))

	if err != nil {
		return nil, err
	}

	return response, nil
}

// errorFromResponse  helper
func (c *APIClient) errorFromResponse(response *http.Response) error {
	c.logger.Error(fmt.Sprintf("openweather.apiClient Unexpected Status. %v  Body: %v", response.StatusCode, response.Body))
	b, _ := ioutil.ReadAll(response.Body)
	return errors.New(string(b))
}

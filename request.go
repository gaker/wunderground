package wunderground

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"net/url"
	"time"
)

const BaseURL = "https://api.weather.com/v2/pws"

//observations/current?stationId=KMAHANOV10&format=json&units=e&apiKey=yourApiKey

func NewRequest(cfg *Config) *Request {
	duration, _ := time.ParseDuration("30s")

	params := url.Values{}
	params.Set("format", "json")
	params.Set("stationId", cfg.StationId)
	if cfg.Unit == "" {
		cfg.Unit = English
	}
	params.Set("units", string(cfg.Unit))
	params.Set("apiKey", cfg.ApiKey)

	return &Request{
		defaultParams: params,
		client: &http.Client{
			Timeout: duration,
		},
	}
}

type Request struct {
	client        *http.Client
	defaultParams url.Values
}

func (r *Request) Params() url.Values {
	return r.defaultParams
}

func (r *Request) parseUrl(uri string) string {
	return fmt.Sprintf("%s%s", BaseURL, uri)
}

// Get performs a GET request to a given url
// unmarshals the response into the given ret
// struct
func (r *Request) Get(
	ctx context.Context,
	uri string,
	ret any) error {

	req, _ := http.NewRequestWithContext(
		ctx,
		http.MethodGet,
		r.parseUrl(uri),
		nil,
	)

	req.URL.RawQuery = r.defaultParams.Encode()

	resp, err := r.client.Do(req)
	if err != nil {
		return err
	}

	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		e := &RespErr{}

		if err := json.NewDecoder(resp.Body).Decode(e); err != nil {
			return errors.New("unable to decode error response")
		}
		return e
	}

	if err := json.NewDecoder(resp.Body).Decode(ret); err != nil {
		return err
	}

	return nil
}

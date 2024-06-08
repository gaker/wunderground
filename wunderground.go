package wunderground

import (
	"context"
)

type unit string

const (
	Metric  unit = "m"
	English unit = "e"
	UKUnits unit = "h"
)

type Config struct {
	ApiKey    string
	StationId string
	Unit      unit
}

func New(cfg *Config) *Wunderground {
	return &Wunderground{
		config: cfg,
		req:    NewRequest(cfg),
	}
}

type Wunderground struct {
	config *Config
	req    *Request
}

func (w *Wunderground) Current(ctx context.Context) (*Observation, error) {
	uri := "/observations/current"
	ret := &ObservationsResponse{}

	if err := w.req.Get(
		ctx,
		uri,
		ret,
	); err != nil {
		return nil, err
	}

	return ret.Observations[0], nil
}

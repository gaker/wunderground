package wunderground

import (
	"encoding/json"
	"time"
)

type APIErr struct {
	Error struct {
		Code    string `json:"code"`
		Message string `json:"message"`
	} `json:"error"`
}

type RespErr struct {
	Metadata struct {
		TransactionId string `json:"transaction_id"`
	} `json:"metadata"`
	Success bool     `json:"success"`
	Errors  []APIErr `json:"errors"`
}

func (re *RespErr) Error() string {
	e, err := json.Marshal(re)
	if err != nil {
		return "unable to marshal error response"
	}
	return string(e)
}

type Data struct {
	Temp        *float64 `json:"temp,omitempty"`
	HeatIndex   *float64 `json:"heatIndex,omitempty"`
	DewPoint    *float64 `json:"dewpt,omitempty"`
	WindChill   *float64 `json:"windChill,omitempty"`
	WindSpeed   *float64 `json:"windSpeed,omitempty"`
	WindGust    *float64 `json:"windGust,omitempty"`
	Pressure    *float64 `json:"pressure,omitempty"`
	PrecipRate  *float64 `json:"precipRate,omitempty"`
	PrecipTotal *float64 `json:"precipTotal,omitempty"`
	Elev        *float64 `json:"elev,omitempty"`
}

// https://docs.google.com/document/d/1KGb8bTVYRsNgljnNH67AMhckY8AQT2FVwZ9urj8SWBs/edit
type Observation struct {
	_ struct{}

	StationId         string     `json:"stationID"`
	ObsTimeUtc        *time.Time `json:"obsTimeUtc"`
	ObsTimeLocal      *string    `json:"obsTimeLocal"`
	Neighborhood      string     `json:"neighborhood"`
	SoftwareType      *string    `json:"softwareType"`
	Country           *string    `json:"country"`
	SolarRadiation    *float64   `json:"solarRadiation"`
	Lon               *float64   `json:"lon"`
	RealtimeFrequency *float64   `json:"realtimeFrequency"`
	Epoch             int        `json:"epoch"`
	Lat               *float64   `json:"lat"`
	UV                *float64   `json:"uv"`
	WindDir           *float64   `json:"winddir"`
	Humidity          *float64   `json:"humidity"`
	QCStatus          int        `json:"qcStatus"`
	Imperial          *Data      `json:"imperial,omitempty"`
	Metric            *Data      `json:"metric,omitempty"`
	UK                *Data      `json:"uk_hybrid,omitempty"`
}

type ObservationsResponse struct {
	Observations []*Observation `json:"observations"`
}

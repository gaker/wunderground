package wunderground_test

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/jarcoal/httpmock"
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"

	"github.com/gaker/wunderground"
)

var _ = Describe("Wunderground", func() {

	var mockErr = func(uri string) {
		resp := &wunderground.RespErr{}
		Expect(json.Unmarshal(
			[]byte(`{
				"metadata":{"transaction_id":"##############"},
				"success":false,
				"errors":[{
					"error":{"code":"CDN-0001","message":"Invalid apiKey."}
				}
			]}`),
			resp)).To(BeNil())
		responder, err := httpmock.NewJsonResponder(
			http.StatusBadRequest,
			resp,
		)
		Expect(err).To(BeNil())
		httpmock.RegisterResponder(
			http.MethodGet,
			fmt.Sprintf("%s%s", wunderground.BaseURL, uri),
			responder,
		)

	}

	var w *wunderground.Wunderground

	var _ = BeforeEach(func() {
		w = wunderground.New(&wunderground.Config{
			ApiKey:    "abc123",
			StationId: "xxx212",
		})
	})

	Context("Current()", func() {

		var imperial = `"imperial":{
			"temp": 97,
			"heatIndex": 94,
			"dewpt": 51,
			"windChill": 97,
			"windSpeed": 4,
			"windGust": 7,
			"pressure": 26.41,
			"precipRate": 0.00,
			"precipTotal": 0.00,
			"elev": 3373
		}`

		const resp = `{
			"observations":[
				{
					"stationID": "SOMEStATION123",
					"obsTimeUtc": "2024-06-08T20:24:00Z",
					"obsTimeLocal": "2024-06-08 15:24:00",
					"neighborhood": "Somewhere",
					"softwareType":null,
					"country": "US",
					"solarRadiation": 958.0,
					"realtimeFrequency": null,
					"epoch": 1717878240,
					"lat": 27.4,
					"lon": -101.111,
					"uv": 8.0,
					"winddir": 119,
					"humidity": 22,
					"qcStatus": 1,
					%s
				}
			]
		}`

		var getResp = func(which string) []byte {
			out := fmt.Sprintf(resp, which)
			return []byte(out)
		}

		var mockImperialSuccess = func() {
			response := wunderground.ObservationsResponse{}
			Expect(
				json.Unmarshal(getResp(imperial), &response),
			).To(BeNil())

			responder, err := httpmock.NewJsonResponder(
				http.StatusOK, response,
			)
			Expect(err).To(BeNil())

			httpmock.RegisterResponder(
				http.MethodGet,
				fmt.Sprintf("%s/observations/current", wunderground.BaseURL),
				responder,
			)
		}

		It("should return an observation", func() {
			mockImperialSuccess()
			obs, err := w.Current(context.Background())

			Expect(err).To(BeNil())
			Expect(obs).ToNot(BeNil())
			Expect(obs.StationId).To(Equal("SOMEStATION123"))
			Expect(obs.Imperial).ToNot(BeNil())
			Expect(obs.UK).To(BeNil())
			Expect(obs.Metric).To(BeNil())
		})

		It("should handle an error", func() {
			mockErr("/observations/current")
			obs, err := w.Current(context.Background())
			Expect(obs).To(BeNil())

			Expect(err).ToNot(BeNil())
		})

	})
})

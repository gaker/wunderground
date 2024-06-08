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

var _ = Describe("Request", func() {
	var cfg = &wunderground.Config{
		ApiKey:    "abc123",
		StationId: "xxx212",
	}

	It("should create a new client with default params", func() {
		r := wunderground.NewRequest(cfg)
		Expect(r.Params().Encode()).To(
			Equal("apiKey=abc123&format=json&stationId=xxx212&units=e"))
	})

	It("should return an error if the statusCode is not 200", func() {
		r := wunderground.NewRequest(cfg)

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
			fmt.Sprintf("%s/foo/bar", wunderground.BaseURL),
			responder,
		)

		err = r.Get(context.Background(), "/foo/bar", nil)
		Expect(err).ToNot(BeNil())
		Expect(err.Error()).To(MatchJSON(`{"metadata":{"transaction_id":"##############"},"success":false,"errors":[{"error":{"code":"CDN-0001","message":"Invalid apiKey."}}]}`))
	})

	It("should return a valid response", func() {
		r := wunderground.NewRequest(cfg)

		type x struct {
			Hello string  `json:"hello"`
			A     int     `json:"a"`
			B     *string `json:"b"`
		}
		ret := &x{}

		Expect(json.Unmarshal(
			[]byte(`{
				"hello": "world",
				"a": 1,
				"b": null
			}`),
			ret)).To(BeNil())
		responder, err := httpmock.NewJsonResponder(
			http.StatusOK,
			ret,
		)

		Expect(err).To(BeNil())
		httpmock.RegisterResponder(
			http.MethodGet,
			fmt.Sprintf("%s/foo/bar", wunderground.BaseURL),
			responder,
		)

		resp := x{}

		Expect(
			r.Get(context.Background(), "/foo/bar", &resp),
		).To(BeNil())

		Expect(resp).To(Equal(*ret))
	})

})

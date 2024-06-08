package main

import (
	"context"

	"github.com/davecgh/go-spew/spew"
	"github.com/gaker/wunderground"
)

func main() {
	api := wunderground.New(&wunderground.Config{
		ApiKey:    "your-api-key-here",
		StationId: "your-station-here",
		Unit:      wunderground.English,
	})

	obs, err := api.Current(context.Background())
	if err != nil {
		panic(err.Error())
	}
	spew.Dump(obs)

}

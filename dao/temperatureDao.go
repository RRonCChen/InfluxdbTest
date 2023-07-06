package dao

import (
	"fmt"

	"github.com/RRonCChen/influxdbTest/db"
	"github.com/RRonCChen/influxdbTest/errors"
)

type TemperatureDao struct{}

func (dao *TemperatureDao) GetAvgTempertureByIdAndLocation(deviceId string, location string) (float64, error) {
	query := fmt.Sprintf(
		`from(bucket: "RonTest")
        |> range(start: -1d)
        |> filter(fn: (r) => r["_measurement"] == "temperature")
        |> filter(fn: (r) => r["_field"] == "temperature")
        |> filter(fn: (r) => r["deviceId"] == "%s")
		|> filter(fn: (r) => r["location"] == "%s")
        |> mean()`, deviceId, location)

	queryTableResult := db.InitInfluxConn().Query(query)
	queryTableResult.Next()
	if queryTableResult.Record() != nil {
		return queryTableResult.Record().Value().(float64), nil
	} else {
		return 0, errors.NotFoundError
	}
}

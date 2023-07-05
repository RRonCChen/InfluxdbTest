package service

import (
	"time"

	"github.com/RRonCChen/influxdbTest/db"
	"github.com/RRonCChen/influxdbTest/model"
)

func InsertTemperature(temperatureReocrd model.TemperatureRecord) {
	record := model.Record{
		Measurement: "temperature",
		Tags:        getTags(temperatureReocrd),
		Fields:      getFields(temperatureReocrd),
		Time:        time.Now(),
	}
	conn := db.InitInfluxConn()
	conn.Insert(record)
}

func getFields(temperatureReocrd model.TemperatureRecord) map[string]any {
	fields := make(map[string]any)
	fields["temperature"] = temperatureReocrd.Temperature
	return fields
}

func getTags(temperatureReocrd model.TemperatureRecord) map[string]string {
	tags := make(map[string]string)
	tags["location"] = temperatureReocrd.Location
	tags["deviceId"] = temperatureReocrd.DeviceId
	return tags
}

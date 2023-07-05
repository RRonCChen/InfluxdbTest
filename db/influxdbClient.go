package db

import (
	"context"
	"log"

	"github.com/RRonCChen/influxdbTest/model"
	influxdb2 "github.com/influxdata/influxdb-client-go/v2"
)

const (
	token  = "MLKxt6MvY7-hO4ifdmsHe1nyW0mvFx8uqVdpU2aJDfKUTnPuMYMvkXmIEhrTXX7VQBIBFOd9Vh7XRgb58KK1JA=="
	url    = "http://localhost:8086"
	bucket = "RonTest"
	org    = "RonTest"
)

type InfluxdbConn struct {
	client influxdb2.Client
}

func InitInfluxConn() *InfluxdbConn {
	return &InfluxdbConn{influxdb2.NewClient(url, token)}
}

func (conn *InfluxdbConn) Insert(record model.Record) {
	writeApi := conn.client.WriteAPIBlocking(org, bucket)

	point := influxdb2.NewPoint(record.Measurement, record.Tags, record.Fields, record.Time)

	err := writeApi.WritePoint(context.Background(), point)
	if err != nil {
		log.Fatal("Influxdb insert fail: ", err)
	}
	defer conn.client.Close()
}

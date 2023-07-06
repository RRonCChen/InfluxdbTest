package db

import (
	"context"
	"log"

	"github.com/RRonCChen/influxdbTest/config"
	"github.com/RRonCChen/influxdbTest/model"
	influxdb2 "github.com/influxdata/influxdb-client-go/v2"
	"github.com/influxdata/influxdb-client-go/v2/api"
)

type InfluxdbConn struct {
	client influxdb2.Client
}

func InitInfluxConn() *InfluxdbConn {
	conf := config.LoadInfluxdbConfig()
	return &InfluxdbConn{influxdb2.NewClient(conf.Url, conf.Token)}
}

func (conn *InfluxdbConn) Insert(record model.Record) {
	conf := config.LoadInfluxdbConfig()
	writeApi := conn.client.WriteAPIBlocking(conf.Org, conf.Bucket)

	point := influxdb2.NewPoint(record.Measurement, record.Tags, record.Fields, record.Time)

	if err := writeApi.WritePoint(context.Background(), point); err != nil {
		log.Fatal("Influxdb insert fail: ", err)
	}
	defer conn.client.Close()
}

func (conn *InfluxdbConn) Query(query string) *api.QueryTableResult {
	conf := config.LoadInfluxdbConfig()
	queryAPI := conn.client.QueryAPI(conf.Org)

	log.Printf("query : %s \n", query)

	result, err := queryAPI.Query(context.Background(), query)
	if err == nil {
		if result.Err() != nil {
			log.Printf("query error: %s\n", result.Err().Error())
			panic(result.Err())
		}
	} else {
		panic(err)
	}
	return result
}

package main

import (
	"github.com/RRonCChen/influxdbTest/middleware"
	"github.com/RRonCChen/influxdbTest/router"
	"github.com/gin-gonic/gin"
)

// @title        Ron influxdb swagger demo
// @version      1.0
// @description  influxdb practice
// @host         localhost:8080
func main() {
	gin := gin.Default()

	middleware.SetMiddleware(gin)
	router.SetRoute(gin)

	gin.Run(":8080")
}

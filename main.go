package main

import (
	"influxdbTest/router"

	"github.com/gin-gonic/gin"
)

func main() {
	gin := gin.Default()
	router.SetRoute(gin)
	gin.Run(":8080")
}

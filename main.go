package main

import (
	_ "github.com/RRonCChen/influxdbTest/docs"
	"github.com/RRonCChen/influxdbTest/router"
	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

// @title        Ron influxdb swagger demo
// @version      1.0
// @description  influxdb practice
// @host         localhost:8080
func main() {
	gin := gin.Default()
	router.SetRoute(gin)

	url := ginSwagger.URL("http://localhost:8080/swagger/doc.json")
	gin.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler, url))

	gin.Run(":8080")
}

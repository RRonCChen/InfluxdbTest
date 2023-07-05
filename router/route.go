package router

import (
	"github.com/RRonCChen/influxdbTest/controller"
	"github.com/gin-gonic/gin"
)

func SetRoute(engine *gin.Engine) {
	controller := controller.NewController()
	influx := engine.Group("/influx")
	{
		v1 := influx.Group("/v1")
		{
			hello := v1.Group("/hello")
			{
				hello.GET("", controller.SayHello)
			}
			temperature := v1.Group("/temperature")
			{
				temperature.POST("", controller.InsertTemperature)
			}
		}
	}
}

package router

import (
	"github.com/RRonCChen/influxdbTest/controller"
	_ "github.com/RRonCChen/influxdbTest/docs"
	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
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
				temperature.GET("/avg/:location/:deviceId", controller.GetAvgTempertureByIdAndLocation)
			}
		}
	}

	docJsonUrl := ginSwagger.URL("http://localhost:8080/swagger/doc.json")
	engine.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler, docJsonUrl))
}

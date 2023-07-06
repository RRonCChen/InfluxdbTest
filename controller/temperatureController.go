package controller

import (
	"log"
	"net/http"

	"github.com/RRonCChen/influxdbTest/model"
	"github.com/RRonCChen/influxdbTest/service"
	"github.com/gin-gonic/gin"
)

// @Tags      Temperature
// @Success   200
// @Accept    json
// @Param     temperatureRecord body model.TemperatureRecord true "location deviceId temperature"
// @Router    /influx/v1/temperature [post]
func (c *Controller) InsertTemperature(context *gin.Context) {
	body := model.TemperatureRecord{}
	if err := context.BindJSON(&body); err != nil {
		setResponse(context, http.StatusBadRequest, "input parameters error")
		return
	}
	log.Printf("input body : %v", body)
	service.InsertTemperature(body)
	setResponse(context, http.StatusOK, "insert successd")
}

// @Tags      Temperature
// @Success   200
// @Accept    json
// @Param     deviceId path string true "deviceId"
// @Param     location path string true "location"
// @Router    /influx/v1/temperature/avg/{location}/{deviceId} [get]
func (c *Controller) GetAvgTempertureByIdAndLocation(context *gin.Context) {
	deviceId := context.Param("deviceId")
	location := context.Param("location")
	if avg, err := service.GetAvgTempertureByIdAndLocation(deviceId, location); err != nil {
		context.Abort()
		context.Error(err)
		return
	} else {
		setResponse(context, http.StatusOK, avg)
	}
}

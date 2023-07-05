package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// @Tags      Hello
// @Success   200
// @Router    /influx/v1/hello [get]
func (c *Controller) SayHello(context *gin.Context) {
	setResponse(context, http.StatusOK, "hello influxdb")
}

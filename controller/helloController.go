package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func (c *Controller) SayHello(context *gin.Context) {
	setResponse(context, http.StatusOK, "hello influxdb")
}

package controller

import (
	"github.com/gin-gonic/gin"
)

type Controller struct{}

func NewController() *Controller {
	return &Controller{}
}

type Response struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
}

func setResponse(context *gin.Context, code int, message string) {
	context.JSON(code, Response{Code: code, Message: message})
}

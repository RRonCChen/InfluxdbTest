package middleware

import (
	"net/http"

	"github.com/RRonCChen/influxdbTest/errors"
	"github.com/gin-gonic/gin"
)

func ErrorHandler() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Next()
		for _, e := range c.Errors {
			err := e.Err
			if customError, ok := err.(*errors.CustomError); ok {
				c.JSON(customError.Code, gin.H{
					"message": customError.Message,
				})
			} else {
				c.JSON(http.StatusInternalServerError, gin.H{
					"msg":  "Some error occurred",
					"data": err.Error(),
				})
			}
		}
	}
}

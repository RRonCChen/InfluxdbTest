package middleware

import (
	"github.com/gin-gonic/gin"
)

func SetMiddleware(engine *gin.Engine) {
	engine.Use(ErrorHandler())
}

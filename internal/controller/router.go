package controller

import (
	"github.com/gin-gonic/gin"
)

func SetupRoutes(r *gin.Engine) {
	r.POST("/login", loginHandler)
	r.POST("/registration", registrationHandler)
	r.GET("/ws", gin.WrapF(webSocketHandler))
}

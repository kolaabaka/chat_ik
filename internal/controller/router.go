package controller

import (
	"github.com/gin-gonic/gin"
)

func SetupRoutes(r *gin.Engine, userController *UserController) {
	r.POST("/login", userController.loginHandler)
	r.POST("/registration", userController.registrationHandler)
	r.GET("/ws", gin.WrapF(webSocketHandler))
}

package routes

import (
	controller "golangsidang/controllers"
	"golangsidang/middleware"

	"github.com/gin-gonic/gin"
)

func UserRoutes(incomingRoutes *gin.Engine) { // membuat routes auth
	incomingRoutes.Use(middleware.Authenticate())               // menggunakan middleware authenticate
	incomingRoutes.GET("/usersss", controller.GetAllAbssenis()) // membuat routes user untuk mengani user
	incomingRoutes.GET("/user/:usersGetId", controller.GetUserAll())
}

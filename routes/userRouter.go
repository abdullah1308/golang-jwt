package routes

import (
	
	controller "github.com/abdullah1308/golang-jwt/controllers"
	"github.com/abdullah1308/golang-jwt/middleware"
	"github.com/gin-gonic/gin"
)

func UserRoutes(incomingRoutes *gin.Engine) {
	// The routes must only be used after authentication
	incomingRoutes.Use(middleware.Authenticate())
	incomingRoutes.GET("/users", controller.GetUsers())
	incomingRoutes.GET("/users/:user_id", controller.GetUser())
}

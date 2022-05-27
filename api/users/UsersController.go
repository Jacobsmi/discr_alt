package users

import (
	"github.com/gin-gonic/gin"
)

func Routes(router *gin.Engine) {

	userRoutes := router.Group("/user")
	{
		userRoutes.GET("/", getUser)
		userRoutes.POST("/", createUser)
	}
}

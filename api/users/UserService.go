package users

import (
	"fmt"

	"github.com/gin-gonic/gin"
)

func getUser(c *gin.Context) {
	fmt.Println("Hello, world!")
	c.JSON(200, gin.H{
		"status": "testing",
	})
}

func createUser(c *gin.Context) {
	c.JSON(200, gin.H{
		"status": "POSTING",
	})
}

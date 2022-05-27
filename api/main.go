package main

import (
	users "discr_api/v1/users"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()
	users.Routes(router)
	router.Run()
}

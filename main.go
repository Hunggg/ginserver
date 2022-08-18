package main

import (
	"hello/server"

	"github.com/gin-gonic/gin"
)


func main() {
	env := server.Env{}
	env.GetConfig()
	host := env.Host
	
	engine := gin.Default()
	
	engine.GET("/", server.GetEmployee)
	
	engine.Run(":" + host)
}
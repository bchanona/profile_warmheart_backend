package main

import (
	supervisor "github.com/bchanona/profile_warmheart_backend/Supervisor/infrastructure/Dependences"
	routeSupervisor "github.com/bchanona/profile_warmheart_backend/Supervisor/infrastructure/Routes"
	"github.com/gin-gonic/gin"
)

func main() {
	supervisor.Init()
	defer supervisor.CloseDB()

	r := gin.Default()
	r.Use(func(c *gin.Context) {
		c.Header("Access-Control-Allow-Origin", "*")
		c.Header("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE, OPTIONS")
		c.Header("Access-Control-Allow-Headers", "Content-Type, Authorization")

		if c.Request.Method == "OPTIONS" {
			c.AbortWithStatus(204)
			return
		}

		c.Next()
	})
	routeSupervisor.Routes(r)
	r.Run(":8081")
}

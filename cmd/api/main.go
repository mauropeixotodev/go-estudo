package main

import "github.com/gin-gonic/gin"

func main() {
	router := gin.Default()
	router.GET("/heatlh", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"success": true,
		})
	})
	router.Run(":8080") // listens on 0.0.0.0:8080 by default
}

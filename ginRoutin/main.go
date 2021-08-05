package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func main() {
	r := gin.Default()

	// services
	r.GET("/api/services", func(context *gin.Context) {
		context.JSON(http.StatusOK, gin.H{
			"data":   []string{"webapp", "gui"},
			"total":  2,
			"limit":  0,
			"offset": 0,
			"errors": nil,
		})
	})

	//

	r.Run(":3002")
}

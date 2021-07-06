package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func main() {
	r := gin.Default()

	// match all request
	r.Any("/test", func(context *gin.Context) {
		context.JSON(http.StatusOK, gin.H{
			"message": "test",
		})
	})

	/*// no route
	r.NoRoute(func(context *gin.Context) {
		context.HTML(http.StatusNotFound, "./views/404.html", nil)
	})*/
	r.Run()
}
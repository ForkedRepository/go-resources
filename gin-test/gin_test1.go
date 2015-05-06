// test_gin
// API reference http://godoc.org/github.com/gin-gonic/gin
package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func main() {
	// Returns a Engine instance with the Logger and Recovery already attached.
	router := gin.Default()

	// Engine contains a RouterGroup
	// http://godoc.org/github.com/gin-gonic/gin#RouterGroup
	router.GET("/", func(c *gin.Context) {
		c.String(http.StatusOK, "hello world")
	})
	router.GET("/ping", func(c *gin.Context) {
		c.String(http.StatusOK, "pong")
	})
	router.POST("/submit", func(c *gin.Context) {
		c.String(http.StatusUnauthorized, "not authorized")
	})
	router.PUT("/error", func(c *gin.Context) {
		c.String(http.StatusInternalServerError, "and error happened :(")
	})
	router.GET("/json", func(c *gin.Context) {
		m := map[string]interface{}{"try": "OK"}
		c.JSON(http.StatusOK, m)
	})
	router.Run(":8080")
}

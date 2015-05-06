// test_gin
package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func main() {
	router := gin.Default()
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

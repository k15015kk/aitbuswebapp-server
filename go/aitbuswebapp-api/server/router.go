package server

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func Router() (*gin.Engine, error) {
	// Ginの定義
	engine := gin.Default()

	// ECHOでルートのGETアクセスしたときに返す
	engine.GET("/", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "Hello world",
		})
	})

	return engine, nil
}

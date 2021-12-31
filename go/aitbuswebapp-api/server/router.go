package server

import (
	"aitbuswebapp-api/controller"
	"net/http"

	"github.com/gin-gonic/gin"
)

func Router() (*gin.Engine, error) {
	// Ginの定義
	engine := gin.Default()

	api := engine.Group("api/v1")

	// ECHOでルートのGETアクセスしたときに返す
	api.GET("/", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "This is main",
		})
	})

	api.GET("/timetable", controller.GetTimeTableByDate)
	api.GET("/timetable/:diagram", controller.GetTimeTableByDiagram)
	api.GET("/calendar/:year/:month", controller.GetCalendar)

	return engine, nil
}

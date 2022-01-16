package server

import (
	"aitbuswebapp-api/controller"
	"net/http"
	"time"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func Router() (*gin.Engine, error) {
	// Ginの定義
	engine := gin.Default()

	engine.Use(cors.New(cors.Config{
		// 許可したいHTTPメソッドの一覧
		AllowMethods: []string{
			"POST",
			"GET",
			"OPTIONS",
			"PUT",
			"DELETE",
		},
		// 許可したいHTTPリクエストヘッダの一覧
		AllowHeaders: []string{
			"Access-Control-Allow-Headers",
			"Content-Type",
			"Content-Length",
			"Accept-Encoding",
			"X-CSRF-Token",
			"Authorization",
		},
		// 許可したいアクセス元の一覧
		AllowOrigins: []string{
			"http://localhost:3000",
		},
		MaxAge: 24 * time.Hour,
	}))

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

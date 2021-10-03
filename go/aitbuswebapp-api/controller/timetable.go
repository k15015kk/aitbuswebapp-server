package controller

import (
	"aitbuswebapp-api/middleware"
	"aitbuswebapp-api/models"
	"fmt"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

func GetTimeTableByDate(c *gin.Context) {
	// 日付の取得
	dateId := c.Query("date")

	var datetime time.Time

	if len(dateId) == 0 {
		// date指定がない場合は，現在時刻の取得
		datetime = middleware.NowTime()
	} else {
		datetime = middleware.StringToTime(dateId)
	}

	fmt.Println(datetime)

	// ダイヤグラムの検索
	dia, diaErr := models.CalendarFindByDate(datetime)

	if diaErr != nil {
		fmt.Println("error")
		c.JSON(http.StatusBadRequest, gin.H{
			"type":    "error",
			"code":    "400",
			"message": "suspension",
		})
	} else {
		fmt.Println(dia)
		// 時刻表を取得
		c.JSON(http.StatusOK, gin.H{
			"dia": dia.ServiceId,
		})
	}

}

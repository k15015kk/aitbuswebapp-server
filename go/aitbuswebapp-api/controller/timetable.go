package controller

import (
	"aitbuswebapp-api/middleware"
	"aitbuswebapp-api/models"
	"fmt"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

type ArrivalAndDepartureTime struct {
	Arrival   string
	Departure string
}

func BadRequest(c *gin.Context, message string) {
	c.JSON(http.StatusBadRequest, gin.H{
		"type":    "error",
		"code":    "400",
		"message": message,
	})
}

func GetTimeTableByDate(c *gin.Context) {
	// 日付の取得
	dateId := c.Query("date")

	// 方向の指定
	directionId := c.Query("direction")

	// 方向が1と2でなければ，無効な方向IDとして返す
	if directionId != "1" && directionId != "2" {
		fmt.Println("error")
		BadRequest(c, "DirectionID is not valid.")
		return
	}

	var datetime time.Time

	if len(dateId) == 0 {
		// date指定がない場合は，現在時刻の取得
		datetime = middleware.NowTime()
	} else {
		datetime = middleware.StringToTime(dateId)
	}

	// ダイヤグラムの検索
	dia, diaErr := models.CalendarFindByDate(datetime)

	// 運行していない場合は，運行していない日であると返す．
	if diaErr != nil {
		fmt.Println("error")
		BadRequest(c, "No service today.")
		return
	}

	nowTime := middleware.NowTime()
	nowTimeString := nowTime.Format("2006-01-02T15:04:05+09:00")
	targetDate := datetime.Format("2006-01-02")

	// 時刻表を返す
	c.JSON(http.StatusOK, gin.H{
		"nowDateTime": nowTimeString,
		"targateDate": targetDate,
		"diagram":     dia.ServiceId,
		"directionId": directionId,
	})

}

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

func GetTimeTableByDate(c *gin.Context) {
	// 日付の取得
	date := c.Query("date")

	// 方向の指定
	direction := c.Query("direction")

	// 方向が1と2でなければ，無効な方向IDとして返す
	if direction != "0" && direction != "1" {
		fmt.Println("error")
		BadRequest(c, "DirectionID is not valid.")
		return
	}

	var datetime time.Time

	if len(date) == 0 {
		// date指定がない場合は，現在時刻の取得
		datetime = middleware.NowTime()
	} else {
		datetime = middleware.StringToTime(date)
	}

	// ダイヤグラムの検索
	dia, diaErr := models.FindCalendarByDate(datetime)

	// 運行していない場合は，運行していない日であると返す．
	if diaErr != nil {
		fmt.Println("error")
		BadRequest(c, "No service today.")
		return
	}

	nowTime := middleware.NowTime()
	nowTimeString := nowTime.Format("2006-01-02T15:04:05+09:00")
	targetDate := datetime.Format("2006-01-02")

	// その日のダイヤの特定方向のtripsを取得する
	trips, tripErr := models.FindTripIdByServiceAndDirection(dia.ServiceId, direction)

	// エラーが起きた場合は，その旨を返す．
	if tripErr != nil {
		fmt.Println("error")
		BadRequest(c, "Could not get trip.")
		return
	}

	fmt.Printf("%#v\n", trips)

	// 時刻表を取得する
	stoptimes, tripErr := models.FindStoptimesByTripIds(trips, "0")

	fmt.Printf("%#v\n", stoptimes)

	// 時刻表を返す
	c.JSON(http.StatusOK, gin.H{
		"nowDateTime": nowTimeString,
		"targateDate": targetDate,
		"diagram":     dia.ServiceId,
		"directionId": direction,
		"schedule":    stoptimes,
	})
}

func GetTimeTableByDiagram(c *gin.Context) {
	// ダイヤグラムの取得
	diagram := c.Param("diagram")

	// 方向の指定
	direction := c.Query("direction")

	// 方向が1と2でなければ，無効な方向IDとして返す
	if direction != "0" && direction != "1" {
		fmt.Println("error")
		BadRequest(c, "DirectionID is not valid.")
		return
	}

	// ダイヤグラムがA,B,Cでない場合は無効なダイヤグラムとして返す
	if diagram != "A" && diagram != "B" && diagram != "C" {
		fmt.Println("error")
		BadRequest(c, "Diagram is not valid.")
		return
	}

	// その日のダイヤの特定方向のtripsを取得する
	trips, tripErr := models.FindTripIdByServiceAndDirection(diagram, direction)

	// エラーが起きた場合は，その旨を返す．
	if tripErr != nil {
		fmt.Println("error")
		BadRequest(c, "Could not fetch trip.")
		return
	}

	fmt.Printf("%#v\n", trips)

	// 時刻表を取得する
	stoptimes, tripErr := models.FindStoptimesByTripIds(trips, "0")

	fmt.Printf("%#v\n", stoptimes)

	// 時刻表を返す
	c.JSON(http.StatusOK, gin.H{
		"diagram":     diagram,
		"directionId": direction,
		"schedule":    stoptimes,
	})

}

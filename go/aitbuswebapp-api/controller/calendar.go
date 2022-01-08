package controller

import (
	"aitbuswebapp-api/models"
	"fmt"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func GetCalendar(c *gin.Context) {
	// 年と月の取得
	yearQuery := c.Param("year")
	monthQuery := c.Param("month")

	year, err := strconv.Atoi(yearQuery)

	if err != nil {
		fmt.Println("error")
		BadRequest(c, "Please confirm Year.")
		return
	}

	month, err := strconv.Atoi(monthQuery)

	if err != nil {
		fmt.Println("error")
		BadRequest(c, "Please confirm Month.")
		return
	}

	calendar, calendarErr := models.FindDiagramByDate(year, month)

	// エラーが起きた場合は，その旨を返す．
	if calendarErr != nil {
		fmt.Println("error")
		BadRequest(c, "Could not fetch calendar.")
		return
	}

	fmt.Printf("%#v\n", calendar)

	// スケジュールを返す
	c.JSON(http.StatusOK, gin.H{
		"year":     year,
		"month":    month,
		"calendar": calendar,
	})
}

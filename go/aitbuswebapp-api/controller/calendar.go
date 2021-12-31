package controller

import (
	"aitbuswebapp-api/models"
	"fmt"
	"strconv"

	"github.com/gin-gonic/gin"
)

func GetCalendar(c *gin.Context) {
	// 年と月の取得
	yearQuery := c.Param("year")
	monthQuery := c.Param("month")

	year, err := strconv.Atoi(yearQuery)
	month, err := strconv.Atoi(monthQuery)

	data, err := models.FindDiagramByDate(year, month)

	// エラーが起きた場合は，その旨を返す．
	if err != nil {
		fmt.Printf("%#v\n", err)
		fmt.Println("error")
		return
	}

	fmt.Printf("%#v\n", data)

}

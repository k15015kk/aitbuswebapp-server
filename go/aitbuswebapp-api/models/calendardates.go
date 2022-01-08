package models

import (
	"aitbuswebapp-api/database"
	"fmt"
	"strconv"
	"time"
)

type CalendarDates struct {
	ServiceId     string
	Date          time.Time
	ExceptionType int
}

func FindCalendarByDate(date time.Time) (CalendarDates, error) {
	var cd CalendarDates
	db := database.GetDB()
	err := db.Where("date = ?", date.Format("2006-01-02")).First(&cd).Error

	return cd, err
}

func FindDiagramByDate(year int, month int) ([]map[string]string, error) {
	var (
		nextYear  int
		nextMonth int
		date      time.Time
		serviceId string
		cds       []CalendarDates
		layout    = "02"
		calendar  []map[string]string
	)

	if month == 12 {
		nextYear = year + 1
		nextMonth = 1
	} else {
		nextYear = year
		nextMonth = month + 1
	}

	db := database.GetDB()

	rows, fetchErr := db.Select("service_id, date").Where("date >= ? AND date < ?", strconv.Itoa(year)+"-"+strconv.Itoa(month)+"-01", strconv.Itoa(nextYear)+"-"+strconv.Itoa(nextMonth)+"-01").Find(&cds).Rows()

	if fetchErr != nil {
		fmt.Println("error")
		fmt.Println("SQL Fetch has failed (calenderdates)")
		return calendar, fetchErr
	}

	for rows.Next() {
		scanErr := rows.Scan(&serviceId, &date)

		if scanErr != nil {
			fmt.Println("error")
			fmt.Println("SQL Scan has failed")
			return calendar, scanErr
		}

		var calendarMap = make(map[string]string, 2)

		calendarMap["service_id"] = serviceId
		calendarMap["date"] = date.Format(layout)

		calendar = append(calendar, calendarMap)
	}

	return calendar, nil
}

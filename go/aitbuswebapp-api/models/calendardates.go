package models

import (
	"aitbuswebapp-api/database"
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

func FindDiagramByDate(year int, month int) ([]CalendarDates, error) {
	var (
		nextYear  int
		nextMonth int
	)

	if month == 12 {
		nextYear = year + 1
		nextMonth = 1
	} else {
		nextYear = year
		nextMonth = month + 1
	}

	var cds []CalendarDates
	db := database.GetDB()
	err := db.Where("date >= ? AND date < ?", strconv.Itoa(year)+"-"+strconv.Itoa(month)+"-01", strconv.Itoa(nextYear)+"-"+strconv.Itoa(nextMonth)+"-01").Find(&cds).Error

	return cds, err
}

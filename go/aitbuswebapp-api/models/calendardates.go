package models

import (
	"aitbuswebapp-api/database"
	"time"
)

type CalendarDates struct {
	ServiceId     string
	Date          time.Time
	ExceptionType int
}

func CalendarFindByDate(date time.Time) (CalendarDates, error) {
	var cd CalendarDates
	db := database.GetDB()
	err := db.Where("date = ?", date.Format("2006-01-02")).First(&cd).Error

	return cd, err
}

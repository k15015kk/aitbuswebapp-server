package models

import (
	"aitbuswebapp-api/database"
)

type CalendarDates struct {
	ServiceId     string
	Date          int
	ExceptionType int
}

func (cd *CalendarDates) FindByDate(date string) (err error) {
	db := database.GetDB()
	return db.Where("date = ?", date).First(cd).Error
}

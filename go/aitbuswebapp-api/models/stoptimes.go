package models

import (
	"aitbuswebapp-api/database"
	"time"
)

type StopTime struct {
	TripId        string
	ArrivalTime   time.Time
	DepartureTime time.Time
	StopId        string
	StopSequence  int
	PickupType    int
	DropOffType   int
}

func StoptimesFindByDepartureTime(deaprtureTime string, stop_sequence string) ([]StopTime, error) {
	var sts []StopTime
	db := database.GetDB()

	err := db.Limit(3).Where("departure_time >= ? AND stop_sequence = ?", deaprtureTime, stop_sequence).Find(&sts).Error

	return sts, err
}

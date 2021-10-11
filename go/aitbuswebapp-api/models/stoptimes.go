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

func FindStoptimesByDepartureTime(deaprture string, sequence string) ([]StopTime, error) {
	var stoptimes []StopTime
	db := database.GetDB()

	err := db.Limit(3).Where("departure_time >= ? AND stop_sequence = ?", deaprture, sequence).Find(&stoptimes).Error

	return stoptimes, err
}

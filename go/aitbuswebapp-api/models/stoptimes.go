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

func FindByDepartureTime(deaprtureTime string) ([]StopTime, error) {
	var sts []StopTime
	db := database.GetDB()

	err := db.Limit(3).Where("departure_time = ?", deaprtureTime).Find(&sts).Error

	return sts, err
}

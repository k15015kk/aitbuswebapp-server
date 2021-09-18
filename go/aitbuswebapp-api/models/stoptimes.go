package models

import (
	"aitbuswebapp-api/database"
)

type StopTimes struct {
	TripId        string
	ArrivalTime   string
	DepartureTime string
	StopId        string
	StopSequence  int
	PickupType    int
	DropOffType   int
}

func (st *StopTimes) FindByTripId(tripId string) (err error) {
	db := database.GetDB()
	return db.Where("trip_id = ?", tripId).First(st).Error
}

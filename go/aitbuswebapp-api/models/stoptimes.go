package models

import (
	"aitbuswebapp-api/database"
	"fmt"
)

// TODO: ここに問題がありそう
type StopTimes struct {
	trip_id        string
	arrival_time   string
	departure_time string
	stop_id        string
	stop_sequence  int
	pickup_type    int
	drop_off_type  int
}

type StopTimesArray []*StopTimes

func (st *StopTimes) FindByTripId(tripId string) (err error) {
	db := database.GetDB()
	test := db.Where("trip_id = ?", tripId).First(st).Value
	fmt.Printf("%#v\n", test)
	return db.Where("trip_id = ?", tripId).First(st).Error
}

package models

import (
	"aitbuswebapp-api/database"
)

type Trip struct {
	RouteId              string
	ServiceId            string
	TripId               string
	TripHeadsign         string
	DirectionId          int
	WheelChairAccessible int
	BikeAllowed          int
}

func FindTripIdByServiceAndDirection(service string, direction string) ([]Trip, error) {
	var trips []Trip
	db := database.GetDB()

	err := db.Where("service_id = ? AND direction_id = ?", service, direction).Find(&trips).Error

	return trips, err
}

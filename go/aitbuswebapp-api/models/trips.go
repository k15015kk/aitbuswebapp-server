package models

import (
	"aitbuswebapp-api/database"
	"fmt"
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

func FindTripIdByServiceAndDirection(service string, direction string) ([]string, error) {
	var (
		tripId      string
		tripIdArray []string
		trips       []Trip
	)

	db := database.GetDB()

	rows, fetchErr := db.Select("trip_id").Where("service_id = ? AND direction_id = ?", service, direction).Find(&trips).Rows()

	defer rows.Close()

	if fetchErr != nil {
		fmt.Println("fetchError")
		fmt.Println("SQL Fetch has failed")
		return tripIdArray, fetchErr
	}

	for rows.Next() {
		scanErr := rows.Scan(&tripId)

		if scanErr != nil {
			fmt.Println("error")
			fmt.Println("SQL Scan has failed")
			return tripIdArray, scanErr
		}

		tripIdArray = append(tripIdArray, tripId)
	}

	return tripIdArray, nil
}

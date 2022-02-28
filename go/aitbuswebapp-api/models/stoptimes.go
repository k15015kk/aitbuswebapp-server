package models

import (
	"aitbuswebapp-api/database"
	"fmt"
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

func FindStoptimesByTripIds(trips []string, sequence string) ([]map[string]string, error) {
	var (
		stoptimes      []StopTime
		departureTime  time.Time
		arrivalTime    time.Time
		departureTimes []map[string]string
		layout         = "15:04"
	)
	db := database.GetDB()

	rows, fetchErr := db.Select("departure_time, arrival_time").Where("trip_id IN (?) AND stop_sequence = ?", trips, sequence).Find(&stoptimes).Order("departure_time").Rows()

	defer rows.Close()

	if fetchErr != nil {
		fmt.Println("error")
		fmt.Println("SQL Fetch has failed")
		return departureTimes, fetchErr
	}

	for rows.Next() {
		scanErr := rows.Scan(&departureTime, &arrivalTime)

		if scanErr != nil {
			fmt.Println("error")
			fmt.Println("SQL Scan has failed")
			return departureTimes, scanErr
		}

		var scheduleMap = make(map[string]string, 2)

		scheduleMap["departure_time"] = departureTime.Format(layout)
		scheduleMap["arrival_time"] = arrivalTime.Format(layout)

		departureTimes = append(departureTimes, scheduleMap)
	}

	return departureTimes, nil
}

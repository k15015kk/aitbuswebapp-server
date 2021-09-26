package main

import (
	"aitbuswebapp-api/database"
	"aitbuswebapp-api/middleware"
	"aitbuswebapp-api/models"
	"aitbuswebapp-api/server"
	"fmt"
	"log"

	_ "github.com/lib/pq"
)

func main() {
	// 現在時刻の取得
	now := middleware.NowTime()
	nowString := middleware.TimeToString(now)
	fmt.Println(nowString)

	// データベースの初期化
	database.Init()
	fmt.Println(database.GetDB())

	defer database.Close()

	testResult, err := models.StoptimesFindByDepartureTime(nowString, "0")

	if err != nil {
		log.Fatal(err)
	}

	for _, v := range testResult {
		fmt.Printf("%#v\n", v.TripId)
	}

	// サーバを立ち上げる
	if err := server.Init(); err != nil {
		log.Fatal(err)
	}
}

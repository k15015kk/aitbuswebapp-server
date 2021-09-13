package main

import (
	"aitbuswebapp-api/database"
	mwtime "aitbuswebapp-api/middleware"
	"aitbuswebapp-api/models"
	"aitbuswebapp-api/server"
	"fmt"
	"log"

	_ "github.com/lib/pq"
)

func main() {
	// 現在時刻の取得
	now := mwtime.NowTime()
	fmt.Println(now)

	// データベースの初期化
	database.Init()
	fmt.Println(database.GetDB())

	defer database.Close()

	stoptimeTest := new(models.StopTimes)
	dbErr := stoptimeTest.FindByTripId("A1000")

	if dbErr != nil {
		log.Fatal(dbErr)
	}

	fmt.Println("stoptime")
	fmt.Printf("%#v\n", stoptimeTest)

	// サーバを立ち上げる
	if err := server.Init(); err != nil {
		log.Fatal(err)
	}
}

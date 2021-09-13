package main

import (
	"aitbuswebapp-api/database"
	mwtime "aitbuswebapp-api/middleware"
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

	// サーバを立ち上げる
	if err := server.Init(); err != nil {
		log.Fatal(err)
	}
}

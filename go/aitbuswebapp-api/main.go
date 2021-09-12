package main

import (
	"aitbuswebapp-api/middleware"
	"aitbuswebapp-api/server"
	"database/sql"
	"fmt"
	_ "github.com/lib/pq"
	"log"
)

func main() {

	// 現在時刻の取得
	now := mwtime.NowTime()
	fmt.Println(now)

	// データベースの初期化
	var Db *sql.DB
	Db, err := sql.Open(
		"postgres",
		"host=postgres user=kyotonagoya  password=password dbname=shuttlebus_gtfs sslmode=disable")

	// 接続エラー時のハンドリング
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(Db)

	// サーバを立ち上げる
	if err := server.Init(); err != nil {
		log.Fatal(err)
	}
}

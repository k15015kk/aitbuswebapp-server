package database

import (
	"log"

	"github.com/jinzhu/gorm"
)

var d *gorm.DB

func Init() {
	var err error

	// データベースの定義
	db := "postgres"
	dbname := "shuttlebus_gtfs"
	host := "postgres"
	user := "kyotonagoya"
	password := "password"
	sslmode := "disable"

	// 接続文章の作成
	connect := "host=" + host + " user=" + user + " password=" + password + " dbname=" + dbname + " sslmode=" + sslmode

	d, err = gorm.Open(db, connect)

	if err != nil {
		log.Fatal(err)
	}
}

func GetDB() *gorm.DB {
	return d
}

func Close() {
	d.Close()
}

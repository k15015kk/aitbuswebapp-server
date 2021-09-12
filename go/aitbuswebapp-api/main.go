package main

import (
	"aitbuswebapp-api/test"
	"database/sql"
	"fmt"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	_ "github.com/lib/pq"
)

func main() {

	// データベースの初期化
	var Db *sql.DB
	Db, err := sql.Open(
		"postgres",
		"host=postgres user=kyotonagoya  password=password dbname=shuttlebus_gtfs sslmode=disable")

	// 接続エラー時のハンドリング
	if err != nil {
		log.Fatal(err)
	}

	fmt.Print(Db)
	fmt.Print("\n")

	test.TestPrint()

	// Ginの定義
	engine := gin.Default()

	// ECHOでルートのGETアクセスしたときに返す
	engine.GET("/", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "Hello world",
		})
	})
	engine.Run(":8080")
}

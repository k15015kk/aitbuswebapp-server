package middleware

import (
	"strconv"
	"time"
)

var nowTime time.Time

const layout = "00:00"

func NowTime() time.Time {
	// 現在時刻を取得してUTCに変換
	nowTime = time.Now()
	nowUTC := nowTime.UTC()

	// JSTを定義
	jst := time.FixedZone("Asia/Tokyo", 9*60*60)

	// UTCからJSTに変換
	nowJST := nowUTC.In(jst)
	return nowJST
}

func TimeToString(t time.Time) string {

	// 時間と分を文字列に変換
	// Hour(), Minute()はfloat64で帰ってくるため，
	// 一旦 int に変換後に文字列に変換
	hourString := strconv.Itoa(int(t.Hour()))
	minuteString := strconv.Itoa(int(t.Minute()))

	return hourString + ":" + minuteString
}

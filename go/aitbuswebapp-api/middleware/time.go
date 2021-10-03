package middleware

import (
	"regexp"
	"strconv"
	"time"
)

var nowTime time.Time

const layout = "00:00"

func CheckRegexp(reg, str string) bool {
	return regexp.MustCompile(reg).Match([]byte(str))
}

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

func StringToTime(date string) time.Time {
	// 文字列が８文字じゃなければ，1970年1月1日を返す
	if CheckRegexp(date, "[0-9]{8}") {
		t, _ := time.Parse("20060102", "19700101")
		return t
	} else {
		t, _ := time.Parse("20060102", date)
		return t
	}
}

func TimeToString(t time.Time) string {
	// 時間と分を文字列に変換
	hourString := strconv.Itoa(t.Hour())
	minuteString := strconv.Itoa(t.Minute())

	return hourString + ":" + minuteString
}

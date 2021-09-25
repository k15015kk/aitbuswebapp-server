package mwtime

import (
	"time"
)

var nowTime time.Time

func NowTime() time.Time {
	nowTime = time.Now()
	return nowTime
}

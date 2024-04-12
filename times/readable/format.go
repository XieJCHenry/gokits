package readable

import (
	"fmt"
	"time"
)

// DurationToChineseReadable 将时间范围转换为 “xx 天 xx 小时 xx 分 xx 秒” 的格式。
// duration 单位为秒
func DurationToChineseReadable(duration int64, splitBySpace bool) string {
	durationTime := int64(time.Duration(duration).Seconds())
	var days, hours, minutes, seconds int64

	if durationTime < 60 {
		seconds = durationTime
	} else if durationTime < 3600 {
		minutes = durationTime / 60
		seconds = durationTime - minutes*60
	} else if durationTime < 3600*24 {
		hours = durationTime / 3600
		remainTime1 := durationTime - hours*3600
		minutes = remainTime1 / 60
		seconds = remainTime1 - minutes*60
	} else {
		days = durationTime / 3600 / 24
		remainTime1 := durationTime - days*3600*24
		hours = remainTime1 / 3600
		remainTime2 := remainTime1 - hours*3600
		minutes = remainTime2 / 60
		seconds = remainTime2 - minutes*60
	}

	var formatter = "%d天%d小时%d分%d秒"
	if splitBySpace {
		formatter = "%d 天 %d 小时 %d 分 %d 秒"
	}
	return fmt.Sprintf(formatter, days, hours, minutes, seconds)
}

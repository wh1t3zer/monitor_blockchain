package util

import "strconv"

func FormatTime(t int64) string {
	if t > 60 {
		minutes := t / 60
		return strconv.Itoa(int(minutes)) + "小时前"
	} else {
		return strconv.Itoa(int(t)) + "分钟前"
	}

}

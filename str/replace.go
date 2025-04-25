package str

import (
	"fmt"
	"regexp"
	"time"
)

type CaseInsensitiveReplacer struct {
	toReplace   *regexp.Regexp
	replaceWith string
}

func NewCaseInsensitiveReplacer(toReplace, replaceWith string) *CaseInsensitiveReplacer {
	return &CaseInsensitiveReplacer{
		toReplace:   regexp.MustCompile("(?i)" + toReplace),
		replaceWith: replaceWith,
	}
}

func (cir *CaseInsensitiveReplacer) Replace(str string) string {
	return cir.toReplace.ReplaceAllString(str, cir.replaceWith)
}

func GetAgeFromBirth(birthDay string) (age int) {
	if birthDay == "" {
		return age
	}

	by, bm, bd, err := GetYMDFromTimeStr(birthDay)
	if err != nil {
		return age
	}
	ny, nm, nd := GetYMDFromTime(time.Now())
	age = ny - by
	if nm <= bm {
		if nm == bm {
			if nd < bd {
				age = age - 1
			}
		} else {
			age = age - 1
		}
	}

	return age
}

func GetYMDFromTimeStr(timeStr string) (y, m, d int, err error) {
	// 定义时间格式
	layout := "2006-01-02 15:04:05"
	// 解析时间字符串
	t, err := time.Parse(layout, timeStr)
	if err != nil {
		layout = "2006-01-02"
		t, err = time.Parse(layout, timeStr)
		if err != nil {
			fmt.Println("时间解析出错:", err)
			return y, m, d, err
		}
	}
	y, m, d = GetYMDFromTime(t)
	return y, m, d, nil
}

func GetYMDFromTime(t time.Time) (y, m, d int) {

	// 获取年、月、日
	y = t.Year()
	m = int(t.Month())
	d = t.Day()
	return y, m, d
}

package str

import "time"

func GetTimeAfterString(sec int) string {
	return time.Now().Add(time.Second * time.Duration(sec)).Format("2006-01-02 15:04:05")
}

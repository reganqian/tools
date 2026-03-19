package yx

import "fmt"

// DoDeleteYxRoom 删除云信房间
func DoDeleteYxRoom(appKey, appSecret string, roomID string) (bool, error) {

	baseURL := fmt.Sprintf("https://rtc.yunxinapi.com/v2/api/rooms/%s", roomID)

	// 调用云信接口DELETE, 返回码是200表示成功
	code, err := SendYunxinDeleteApi(appKey, appSecret, baseURL)
	if err != nil {
		return false, err
	}
	if code != 200 {
		return false, fmt.Errorf("delete room failed, code: %d", code)
	}
	//删除成功
	return true, nil
}

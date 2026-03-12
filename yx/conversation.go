package yx

import (
	"encoding/json"
	"fmt"
)

type GetUnreadCountResponse struct {
	Code int    `json:"code"`
	Msg  string `json:"msg"`
	Data struct {
		AccountID   string `json:"account_id"`
		UnreadCount int    `json:"unread_count"`
	} `json:"data"`
}

// 获取未读总数 GET https://{endpoint}/im/v2/conversation_overviews/{account_id}
func GetUnreadCount(appKey, appSecret string, accountID string) (response GetUnreadCountResponse, err error) {
	if accountID == "" {
		return response, fmt.Errorf("accountID is empty")
	}

	requestBody := map[string]interface{}{
		"account_id": accountID,
	}

	baseUrl := fmt.Sprintf("https://rtc.yunxinapi.com/im/v2/conversation_overviews/%s", accountID)
	// 发送 GET 请求到云信 RTC 服务
	resp, err := SendYunxinGetApi(appKey, appSecret, baseUrl, requestBody)
	if err != nil {
		return response, err
	}
	// 解析响应
	// response = GetUnreadCountResponse{}
	err = json.Unmarshal([]byte(resp), &response)
	if err != nil {
		return response, err
	}

	return response, nil
}

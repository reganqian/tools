package yx

import (
	"encoding/json"
	"fmt"
)

// 调用云信接口发送通知
// 参数名称	类型	是否必选	描述
// content	String	是	广播消息内容，长度上限 4096 位字符。
// from_account_id	String	否	广播消息发送者的 IM 账号 ID。
// offline_enabled	Boolean	否	广播消息是否存离线。默认为 false（不存）。
// ttl	Integer	否	存离线的有效时间，单位为小时，默认为 168 小时，即 7 天。
// 只有 offline_enabled = true 时，该参数才有效。
// target_os	Array of strings	否	接收广播消息的目标客户端，默认为所有客户端。JSONArray 格式，例如：
// ["ios","aos","pc","web","mac"]
func SendYXNotification(appKey, appSecret string, req SendYXNotificationRequest) (response SendYXNotificationResponse, err error) {
	if req.Content == "" {
		return response, fmt.Errorf("content is empty")
	}

	params := map[string]interface{}{
		"content": req.Content,
	}
	if req.FromAccountID != "" {
		params["from_account_id"] = req.FromAccountID
	}
	if req.OfflineEnabled {
		params["offline_enabled"] = req.OfflineEnabled
	}
	if req.TTL > 0 {
		params["ttl"] = req.TTL
	}
	if len(req.TargetOS) > 0 {
		params["target_os"] = req.TargetOS
	}

	baseUrl := "https://rtc.yunxinapi.com/v2/broadcast_notification"
	// 发送 POST 请求到云信 RTC 服务
	resp, err := SendYunxinPostApi(appKey, appSecret, baseUrl, params)
	if err != nil {
		return response, err
	}
	// 解析响应
	// response = SendYXNotificationResponse{}
	err = json.Unmarshal([]byte(resp), &response)
	if err != nil {
		return response, err
	}

	return response, nil
}

type SendYXNotificationResponse struct {
	Code int    `json:"code"` // 状态码，200 表示请求成功。	是
	Msg  string `json:"msg"`  // 提示信息。请求失败时返回错误信息，请求成功时返回 "success"。	是
	Data struct {
		BroadcastID    string   `json:"broadcast_id"`    // 广播消息的 ID。	是
		FromAccountID  string   `json:"from_account_id"` // 广播消息发送方账号 ID。	否
		Content        string   `json:"content"`         // 广播消息内容。	是
		OfflineEnabled bool     `json:"offline_enabled"` // 广播消息是否存离线。	否
		TTL            int      `json:"ttl"`             // 存离线的有效时间，单位为小时，默认为 168 小时，即 7 天。只有 offline_enabled = true 时，该参数才有效。
		TargetOS       []string `json:"target_os"`       // 接收广播消息的目标客户端。	否
		CreateTime     int64    `json:"create_time"`     // 广播消息发送时间戳。	是
		ExpireTime     int64    `json:"expire_time"`     // 广播消息过期时间戳。	是
	} `json:"data"` // - data	Object	返回的 JSON 数据对象，请求失败则返回空对象。	是

}

type SendYXNotificationRequest struct {
	Content        string   `json:"content"`         // 广播消息内容，长度上限 4096 位字符。
	FromAccountID  string   `json:"from_account_id"` // 广播消息发送者的 IM 账号 ID。可不指定
	OfflineEnabled bool     `json:"offline_enabled"` // 广播消息是否存离线。默认为 false（不存）。
	TTL            int      `json:"ttl"`             // 存离线的有效时间，单位为小时，默认为 168 小时，即 7 天。只有 offline_enabled = true 时，该参数才有效。
	TargetOS       []string `json:"target_os"`       // 接收广播消息的目标客户端，默认为所有客户端。JSONArray 格式，例如： // ["ios","aos","pc","web","mac"]
}

type CustomNotificationRequest struct {
	SenderID   string `json:"sender_id"`   // 发送者 IM 账号 ID。
	Type       int    `json:"type"`        // 自定义系统通知类型。1：单聊系统通知。2：高级群系统通知。3：超大群系统通知。
	ReceiverID string `json:"receiver_id"` // 系统通知接收者 ID。type =1 时，该参数接收者账号 ID。type = 2 或 3 时，该参数为接收的群 ID（创建群组时服务器生成并返回的 ID）。
	Content    string `json:"content"`     // 自定义系统通知的内容，由开发者自行组装的 JSON 格式字符串，长度上限 4096 位字符。
	Sound      string `json:"sound"`       // 指定的客户端本地的声音文件名，长度上限 30 位字符。
}

// 发送自定义系统通知
func CustomNotification(appKey, appSecret string, req CustomNotificationRequest) (jsonStr string, err error) {
	if req.Content == "" {
		return jsonStr, fmt.Errorf("content is empty")
	}
	if req.SenderID == "" {
		return jsonStr, fmt.Errorf("sender_id is empty")
	}
	if req.Type == 0 {
		return jsonStr, fmt.Errorf("type is empty")
	}
	if req.ReceiverID == "" {
		return jsonStr, fmt.Errorf("receiver_id is empty")
	}

	params := map[string]interface{}{
		"content":     req.Content,
		"sender_id":   req.SenderID,
		"type":        req.Type,
		"receiver_id": req.ReceiverID,
		"sound":       req.Sound,
	}

	baseUrl := "https://rtc.yunxinapi.com/v2/custom_notification"
	// 发送 POST 请求到云信 RTC 服务
	resp, err := SendYunxinPostApi(appKey, appSecret, baseUrl, params)
	if err != nil {
		return jsonStr, err
	}

	return resp, nil
}

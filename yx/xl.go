package yx

import (
	"encoding/json"
	"fmt"
)

// 信令
type XlInfoResponse struct {
	Code int    `json:"code"` //	业务状态码，200 表示请求成功。	是
	Msg  string `json:"msg"`  //	提示信息。请求失败时返回错误信息，请求成功时返回 "success"。	是
	Data struct {
		ChannelId        string `json:"channel_id"`         //	信令房间 ID，服务器生成，唯一标识符。	是
		ChannelName      string `json:"channel_name"`       //	信令房间名称。	是
		CreatorAccountId string `json:"creator_account_id"` //	信令房间创建者账号 ID。	是
		ChannelExtension string `json:"channel_extension"`  //	信令房间相关自定义扩展字段。	否
		ChannelType      int    `json:"channel_type"`       //	信令房间类型，1：音频；2：视频；3：自定义。	是
		Createtime       int64  `json:"create_time"`        //	信令房间创建时间。	是
		ExpireTime       int64  `json:"expire_time"`        //	信令房间过期时间。	是
		MemberList       []struct {
			AccountId  string `json:"account_id"`  //	信令房间中的成员列表。	是
			Uid        int64  `json:"uid"`         //	成员 uid。	是
			DeviceId   string `json:"device_id"`   //	成员设备 ID。	否
			JoinTime   int64  `json:"join_time"`   //	成员进入信令房间的时间。	是
			ExpireTime int64  `json:"expire_time"` //	成员过期时间。
		} `json:"member_list"` //	信令房间中的成员列表。	是
	} `json:"data"` //	返回的 JSON 数据对象，请求失败则返回空对象。	是
}

// 获取未读总数 GET https://{endpoint}/im/v2/conversation_overviews/{account_id}
func GetXlInfo(appKey, appSecret string, channelName string) (response XlInfoResponse, err error) {
	if channelName == "" {
		return response, fmt.Errorf("channelName is empty")
	}

	requestBody := map[string]interface{}{
		"channel_id": channelName,
	}

	baseUrl := fmt.Sprintf("https://rtc.yunxinapi.com/im/v2/signalling_room")
	// 发送 GET 请求到云信 RTC 服务
	resp, err := SendYunxinGetApi(appKey, appSecret, baseUrl, requestBody)
	if err != nil {
		return response, err
	}
	// 解析响应

	err = json.Unmarshal([]byte(resp), &response)
	if err != nil {
		return response, err
	}

	return response, nil
}

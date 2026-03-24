package yx

import "fmt"

func DoDisableUser(appKey, appSecret, accountId string) (jsonStr string, err error) {

	params := map[string]interface{}{}
	params["enabled"] = true
	params["need_kick"] = true
	// params["kick_notify_extension"] = "notification"

	baseUrl := fmt.Sprintf("https://rtc.yunxinapi.com/im/v2/accounts/%s/actions/disable", accountId)
	// 发送 POST 请求到云信 自定义系统通知服务
	resp, err := SendYunxinPatchApi(appKey, appSecret, baseUrl, params)
	if err != nil {
		return jsonStr, err
	}

	return resp, nil
}

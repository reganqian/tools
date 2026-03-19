package yx

import "fmt"

//POST https://{endpoint}/im/v2/accounts/{account_id}/actions/kick 强制账号退出, deviceId 为保留的设备id, 如果deviceId 为空, 则强制退出所有设备
func DoForceLogout(appKey, appSecret, accountId, deviceId string) (jsonStr string, err error) {

	params := map[string]interface{}{}
	if deviceId == "" {
		params["type"] = 1 //所有设备都踢下线
	} else {
		params["type"] = 3
		params["device_id_list"] = []string{deviceId}
	}

	baseUrl := fmt.Sprintf("https://rtc.yunxinapi.com/im/v2/accounts/%s/actions/kick", accountId)
	// 发送 POST 请求到云信 自定义系统通知服务
	resp, err := SendYunxinPostApi(appKey, appSecret, baseUrl, params)
	if err != nil {
		return jsonStr, err
	}

	return resp, nil
}

package tx

import (
	"bytes"
	"encoding/json"
	"math/rand"
	"net/http"
)

type ForceEndCallResp struct {
	ErrorCode    int    `json:"ErrorCode"`
	ErrorInfo    string `json:"ErrorInfo"`
	ActionStatus string `json:"ActionStatus"`
	RequestId    string `json:"RequestId"`
}

// CheckAccount 查询单个 IM 账号是否存在
func ForceEndCall(
	conf TxImConf,
	CallId string,
) (*ForceEndCallResp, error) {

	// 1. 生成 UserSig
	userSig, err := GetTxSig(&conf, conf.AdminUser)
	if err != nil {
		return nil, err
	}

	// 2. 请求体（单个账号）
	reqBody := map[string]string{
		"CallId": CallId,
	}
	bodyBytes, _ := json.Marshal(reqBody)

	// 3. 拼接 URL
	url := "https://console.tim.qq.com/v4/call_engine_http_srv/end_call" +
		"?sdkappid=" + intToStr(conf.SDKAppID) +
		"&identifier=" + conf.AdminUser +
		"&usersig=" + userSig +
		"&random=" + intToStr(int(rand.Int63())) +
		"&contenttype=json"

	// 4. 发起请求
	resp, err := http.Post(url, "application/json", bytes.NewReader(bodyBytes))
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	// 5. 解析响应
	var result ForceEndCallResp
	if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
		return nil, err
	}

	return &result, nil
}

type GetTxCallInfoResp struct {
	ErrorCode    int    `json:"ErrorCode"`
	ErrorInfo    string `json:"ErrorInfo"`
	ActionStatus string `json:"ActionStatus"`
	RequestId    string `json:"RequestId"`
	Response     struct {
		CallRecord struct {
			CallId            string   `json:"CallId"`
			CallerAccount     string   `json:"Caller_Account"`
			MediaType         string   `json:"MediaType"`
			CallType          string   `json:"CallType"`
			StartTime         int64    `json:"StartTime"`
			EndTime           int64    `json:"EndTime"`
			AcceptTime        int64    `json:"AcceptTime"`
			CallResult        string   `json:"CallResult"`
			CalleeListAccount []string `json:"CalleeList_Account"`
			RoomId            string   `json:"RoomId"`
			RoomIdType        int      `json:"RoomIdType"`
		} `json:"CallRecord"`
	} `json:"Response"`
}

// CheckAccount 查询单个 IM 账号是否存在
func GetTxCallInfo(
	conf TxImConf,
	CallId string,
) (*GetTxCallInfoResp, error) {

	// 1. 生成 UserSig
	userSig, err := GetTxSig(&conf, conf.AdminUser)
	if err != nil {
		return nil, err
	}

	// 2. 请求体（单个账号）
	reqBody := map[string]string{
		"CallId": CallId,
	}
	bodyBytes, _ := json.Marshal(reqBody)

	// 3. 拼接 URL
	url := "https://console.tim.qq.com/v4/call_record_http_srv/get_record_by_callid" +
		"?sdkappid=" + intToStr(conf.SDKAppID) +
		"&identifier=" + conf.AdminUser +
		"&usersig=" + userSig +
		"&random=" + intToStr(int(rand.Int63())) +
		"&contenttype=json"

	// 4. 发起请求
	resp, err := http.Post(url, "application/json", bytes.NewReader(bodyBytes))
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	// 5. 解析响应
	var result GetTxCallInfoResp
	if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
		return nil, err
	}

	return &result, nil
}

package tx

import (
	"bytes"
	"encoding/json"
	"fmt"
	"log"
	"math/rand"
	"net/http"
	"time"
)

// ImportAccountResp 导入账号响应体结构
type ImportAccountResp struct {
	ActionStatus string `json:"ActionStatus"`
	ErrorInfo    string `json:"ErrorInfo"`
	ErrorCode    int    `json:"ErrorCode"`
}

type ImportAccountReq struct {
	UserID  string `json:"UserID"`            // 必填，用户名
	Nick    string `json:"Nick,omitempty"`    // 选填，昵称
	FaceUrl string `json:"FaceUrl,omitempty"` // 选填，头像URL
}

func RegistTxImUser(conf TxImConf, req ImportAccountReq) (*ImportAccountResp, error) {

	userSig, err := GetTxSig(&conf, conf.AdminUser)
	if err != nil {
		return nil, fmt.Errorf("生成用户Sig失败: %w", err)
	}

	// 2. 构造请求 URL
	random := rand.Int63()
	apiPath := "v4/im_open_login_svc/account_import"
	url := fmt.Sprintf("https://%s/%s?sdkappid=%d&identifier=%s&usersig=%s&random=%d&contenttype=json",
		conf.APIHost, apiPath, conf.SDKAppID, conf.AdminUser, userSig, random)

	// 3. 序列化请求体
	bodyBytes, err := json.Marshal(req)
	if err != nil {
		return nil, fmt.Errorf("序列化请求体失败: %w", err)
	}

	// 4. 发送 POST 请求
	client := &http.Client{Timeout: 10 * time.Second}
	httpReq, _ := http.NewRequest("POST", url, bytes.NewBuffer(bodyBytes))
	httpReq.Header.Set("Content-Type", "application/json")

	resp, err := client.Do(httpReq)
	if err != nil {
		return nil, fmt.Errorf("HTTP请求失败: %w", err)
	}
	defer resp.Body.Close()

	// 5. 解析响应
	var result ImportAccountResp
	if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
		return nil, fmt.Errorf("解析响应失败: %w", err)
	}

	return &result, nil
}

type ProfileItem struct {
	Tag   string      `json:"Tag"`
	Value interface{} `json:"Value"`
}

type portraitSetReq struct {
	FromAccount string        `json:"From_Account"`
	ProfileItem []ProfileItem `json:"ProfileItem"`
}

type PortraitSetResp struct {
	ActionStatus string `json:"ActionStatus"`
	ErrorCode    int    `json:"ErrorCode"`
	ErrorInfo    string `json:"ErrorInfo"`
}

// SetUserProfile 设置 IM 用户资料
// sdkAppID: 控制台分配的 SDKAppID
// secretKey: 控制台获取的 SecretKey
// adminUserID: App 管理员 UserID
// targetUserID: 需要设置资料的目标 UserID
// items: 资料字段列表，如 []profileItem{{Tag:"Tag_Profile_IM_Nick", Value:"昵称"}}
func SetUserProfile(conf TxImConf,
	targetUserID string,
	items []ProfileItem,
) (*PortraitSetResp, error) {

	// 1. 生成 UserSig
	userSig, err := GetTxSig(&conf, conf.AdminUser)
	if err != nil {
		return nil, err
	}

	// 2. 组装请求体
	reqBody := portraitSetReq{
		FromAccount: targetUserID,
		ProfileItem: items,
	}
	bodyBytes, _ := json.Marshal(reqBody)
	random := rand.Int63()
	// 3. 组装请求 URL
	url := "https://console.tim.qq.com/v4/profile/portrait_set" +
		"?sdkappid=" + intToStr(conf.SDKAppID) +
		"&identifier=" + conf.AdminUser +
		"&usersig=" + userSig +
		"&random=" + intToStr(int(random)) +
		"&contenttype=json"

	log.Println("SetUserProfile url: %v", url)
	// 4. 发送 POST 请求
	resp, err := http.Post(url, "application/json", bytes.NewReader(bodyBytes))
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	// 5. 解析响应
	var result PortraitSetResp
	if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
		return nil, err
	}

	return &result, nil
}

func intToStr(i int) string {
	return fmt.Sprintf("%d", i)
}

// AccountCheckResult 单个账号查询结果
type AccountCheckResult struct {
	UserID        string `json:"UserID"`
	ResultCode    int    `json:"ResultCode"`
	ResultInfo    string `json:"ResultInfo"`
	AccountStatus string `json:"AccountStatus"` // Imported / NotImported
}

// AccountCheckResp 接口返回结构
type AccountCheckResp struct {
	ActionStatus string               `json:"ActionStatus"`
	ErrorCode    int                  `json:"ErrorCode"`
	ErrorInfo    string               `json:"ErrorInfo"`
	ResultItem   []AccountCheckResult `json:"ResultItem"`
}

// CheckAccount 查询单个 IM 账号是否存在
func CheckAccount(
	conf TxImConf,
	userID string,
) (*AccountCheckResp, error) {

	// 1. 生成 UserSig
	userSig, err := GetTxSig(&conf, conf.AdminUser)
	if err != nil {
		return nil, err
	}

	// 2. 请求体（单个账号）
	reqBody := map[string][]map[string]string{
		"CheckItem": {
			{"UserID": userID},
		},
	}
	bodyBytes, _ := json.Marshal(reqBody)

	// 3. 拼接 URL
	url := "https://console.tim.qq.com/v4/im_open_login_svc/account_check" +
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
	var result AccountCheckResp
	if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
		return nil, err
	}

	return &result, nil
}

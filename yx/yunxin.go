package yx

import (
	"bytes"
	"crypto/sha1"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"strings"
	"time"

	"github.com/google/uuid"
	"github.com/labstack/gommon/log"
)

const (
	REQUEST_TYPE_POST   = "POST"   // 调用云信接口POST
	REQUEST_TYPE_GET    = "GET"    // 调用云信接口GET
	REQUEST_TYPE_DELETE = "DELETE" // 调用云信接口DELETE
	REQUEST_TYPE_PATCH  = "PATCH"  // 调用云信接口PATCH
)

// 调用云信接口PATCH
func SendYunxinPatchApi(appKey, appSecret, baseURL string, requestBody map[string]interface{}) (string, error) {
	fmt.Println("SendYunxinPatchApi", requestBody)
	// 序列化为 JSON
	jsonData, err := json.Marshal(requestBody)
	if err != nil {
		fmt.Println("Error marshalling JSON:", err)
		return "", err
	}
	fmt.Println("request jsondata:", string(jsonData))

	//curl -X POST -H "AppKey: go9***3mgq3" -H "Nonce: 4t***23t23t" -H "CurTime: 1443592222" -H "CheckSum: 9e9db3b***583f86" -H "Content-Type: application/x-www-form-urlencoded" -d 'accid=123456&name=zhangsan' 'https://api.yunxinapi.com/nimserver/user/create.action'
	// 生成随机 nonce
	nonce := uuid.New().String()

	// 获取当前时间戳（秒）
	curTime := fmt.Sprintf("%d", time.Now().Unix())

	// 计算 CheckSum = SHA1(AppSecret + nonce + curTime)
	checkSum := CalculateCheckSum(appSecret, nonce, curTime)

	// req.ContentLength = int64(len(formValues.Encode()))
	// 创建 HTTP 请求
	req, err := http.NewRequest(REQUEST_TYPE_PATCH, baseURL, bytes.NewBuffer(jsonData))

	if err != nil {
		fmt.Println("Error creating request:", err)
		return "", err
	}

	// 设置请求头
	req.Header.Set("AppKey", appKey)
	req.Header.Set("Nonce", nonce)
	req.Header.Set("CurTime", curTime)
	req.Header.Set("CheckSum", checkSum)
	req.Header.Set("Content-Type", "application/json;charset=utf-8") //v10版本是json， v9版本是application/x-www-form-urlencoded;charset=utf-8
	// 发起请求
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		fmt.Println("Error sending request:", err)
		return "", err
	}
	defer resp.Body.Close()

	// 读取响应体
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Println("Error reading response body:", err)
		return "", err
	}

	// 直接作为字符串输出（即原始 JSON）
	responseJSON := string(body)

	// 打印 JSON 字符串
	log.Info("baseURL:", baseURL)
	log.Info("Response Status:", resp.Status)
	log.Info("Response JSON String:", responseJSON)
	return responseJSON, nil
}

// 调用云信接口POST
func SendYunxinPostApi(appKey, appSecret, baseURL string, requestBody map[string]interface{}) (string, error) {
	fmt.Println("SendYunxinApi", requestBody)
	// 序列化为 JSON
	jsonData, err := json.Marshal(requestBody)
	if err != nil {
		fmt.Println("Error marshalling JSON:", err)
		return "", err
	}
	fmt.Println("request jsondata:", string(jsonData))

	//curl -X POST -H "AppKey: go9***3mgq3" -H "Nonce: 4t***23t23t" -H "CurTime: 1443592222" -H "CheckSum: 9e9db3b***583f86" -H "Content-Type: application/x-www-form-urlencoded" -d 'accid=123456&name=zhangsan' 'https://api.yunxinapi.com/nimserver/user/create.action'
	// 生成随机 nonce
	nonce := uuid.New().String()

	// 获取当前时间戳（秒）
	curTime := fmt.Sprintf("%d", time.Now().Unix())

	// 计算 CheckSum = SHA1(AppSecret + nonce + curTime)
	checkSum := CalculateCheckSum(appSecret, nonce, curTime)

	// req.ContentLength = int64(len(formValues.Encode()))
	// 创建 HTTP 请求
	req, err := http.NewRequest(REQUEST_TYPE_POST, baseURL, bytes.NewBuffer(jsonData))

	if err != nil {
		fmt.Println("Error creating request:", err)
		return "", err
	}

	// 设置请求头
	req.Header.Set("AppKey", appKey)
	req.Header.Set("Nonce", nonce)
	req.Header.Set("CurTime", curTime)
	req.Header.Set("CheckSum", checkSum)
	req.Header.Set("Content-Type", "application/json;charset=utf-8") //v10版本是json， v9版本是application/x-www-form-urlencoded;charset=utf-8
	// 发起请求
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		fmt.Println("Error sending request:", err)
		return "", err
	}
	defer resp.Body.Close()

	// 读取响应体
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Println("Error reading response body:", err)
		return "", err
	}

	// 直接作为字符串输出（即原始 JSON）
	responseJSON := string(body)

	// 打印 JSON 字符串
	log.Info("baseURL:", baseURL)
	log.Info("Response Status:", resp.Status)
	log.Info("Response JSON String:", responseJSON)
	return responseJSON, nil
}

// 调用云信接口POST
func SendYunxinPostApiWithForm(appKey, appSecret, baseURL string, requestBody map[string]interface{}) (string, error) {
	fmt.Println("SendYunxinApi", requestBody)
	// 序列化为 JSON
	jsonData, err := json.Marshal(requestBody)
	if err != nil {
		fmt.Println("Error marshalling JSON:", err)
		return "", err
	}
	fmt.Println("request jsondata:", string(jsonData))

	//curl -X POST -H "AppKey: go9***3mgq3" -H "Nonce: 4t***23t23t" -H "CurTime: 1443592222" -H "CheckSum: 9e9db3b***583f86" -H "Content-Type: application/x-www-form-urlencoded" -d 'accid=123456&name=zhangsan' 'https://api.yunxinapi.com/nimserver/user/create.action'
	// 生成随机 nonce
	nonce := uuid.New().String()

	// 获取当前时间戳（秒）
	curTime := fmt.Sprintf("%d", time.Now().Unix())

	// 计算 CheckSum = SHA1(AppSecret + nonce + curTime)
	checkSum := CalculateCheckSum(appSecret, nonce, curTime)

	// req.ContentLength = int64(len(formValues.Encode()))
	// 创建 HTTP 请求
	req, err := http.NewRequest(REQUEST_TYPE_POST, baseURL, bytes.NewBuffer(jsonData))

	if err != nil {
		fmt.Println("Error creating request:", err)
		return "", err
	}

	// 设置请求头
	req.Header.Set("AppKey", appKey)
	req.Header.Set("Nonce", nonce)
	req.Header.Set("CurTime", curTime)
	req.Header.Set("CheckSum", checkSum)
	// req.Header.Set("Content-Type", "application/json;charset=utf-8") //v10版本是json， v9版本是application/x-www-form-urlencoded;charset=utf-8
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded;charset=utf-8") //v10版本是json， v9版本是
	// 发起请求
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		fmt.Println("Error sending request:", err)
		return "", err
	}
	defer resp.Body.Close()

	// 读取响应体
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Println("Error reading response body:", err)
		return "", err
	}

	// 直接作为字符串输出（即原始 JSON）
	responseJSON := string(body)

	// 打印 JSON 字符串
	log.Info("baseURL:", baseURL)
	log.Info("Response Status:", resp.Status)
	log.Info("Response JSON String:", responseJSON)
	return responseJSON, nil
}

// 调用云信接口DELETE, 返回码是200表示成功
func SendYunxinDeleteApi(appKey, appSecret, baseURL string) (int, error) {

	//curl -X POST -H "AppKey: go9***3mgq3" -H "Nonce: 4t***23t23t" -H "CurTime: 1443592222" -H "CheckSum: 9e9db3b***583f86" -H "Content-Type: application/x-www-form-urlencoded" -d 'accid=123456&name=zhangsan' 'https://api.yunxinapi.com/nimserver/user/create.action'
	// 生成随机 nonce
	nonce := uuid.New().String()

	// 获取当前时间戳（秒）
	curTime := fmt.Sprintf("%d", time.Now().Unix())

	// 计算 CheckSum = SHA1(AppSecret + nonce + curTime)
	checkSum := CalculateCheckSum(appSecret, nonce, curTime)

	// req.ContentLength = int64(len(formValues.Encode()))
	// 创建 HTTP 请求
	req, err := http.NewRequest(REQUEST_TYPE_DELETE, baseURL, nil)

	if err != nil {
		fmt.Println("Error creating request:", err)
		return 0, err
	}

	// 设置请求头
	req.Header.Set("AppKey", appKey)
	req.Header.Set("Nonce", nonce)
	req.Header.Set("CurTime", curTime)
	req.Header.Set("CheckSum", checkSum)
	req.Header.Set("Content-Type", "application/json;charset=utf-8") //v10版本是json， v9版本是application/x-www-form-urlencoded;charset=utf-8
	// 发起请求
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		fmt.Println("Error sending request:", err)
		return 0, err
	}
	defer resp.Body.Close()

	// // 直接作为字符串输出（即原始 JSON）
	// body, err := ioutil.ReadAll(resp.Body)
	// if err != nil {
	// 	fmt.Println("Error reading response body:", err)
	// 	return 0, err
	// }
	// responseJSON := string(body)

	// 打印 JSON 字符串
	log.Info("baseURL:", baseURL)
	log.Info("Response Status:", resp.Status)
	// log.Info("Response JSON String:", responseJSON)
	return resp.StatusCode, nil
}

// 调用云信接口GET
func SendYunxinGetApi(appKey, appSecret, baseURL string, requestBody map[string]interface{}) (string, error) {
	//url序列化参数
	// 将 requestBody 中的参数序列化为 URL 查询字符串
	values := url.Values{}
	for k, v := range requestBody {
		switch val := v.(type) {
		case string:
			values.Add(k, val)
		case int, int64, uint, uint64:
			values.Add(k, fmt.Sprintf("%v", val))
		case float32, float64:
			values.Add(k, fmt.Sprintf("%v", val))
		case bool:
			values.Add(k, fmt.Sprintf("%v", val))
		default:
			// 对于复杂类型，尝试 JSON 序列化
			if b, err := json.Marshal(v); err == nil {
				values.Add(k, string(b))
			}
		}
	}
	queryStr := values.Encode()
	if queryStr != "" {
		baseURL += "?" + queryStr
	}

	//curl -X POST -H "AppKey: go9***3mgq3" -H "Nonce: 4t***23t23t" -H "CurTime: 1443592222" -H "CheckSum: 9e9db3b***583f86" -H "Content-Type: application/x-www-form-urlencoded" -d 'accid=123456&name=zhangsan' 'https://api.yunxinapi.com/nimserver/user/create.action'
	// 生成随机 nonce
	nonce := uuid.New().String()

	// 获取当前时间戳（秒）
	curTime := fmt.Sprintf("%d", time.Now().Unix())

	// 计算 CheckSum = SHA1(AppSecret + nonce + curTime)
	checkSum := CalculateCheckSum(appSecret, nonce, curTime)

	// req.ContentLength = int64(len(formValues.Encode()))
	// 创建 HTTP 请求
	req, err := http.NewRequest(REQUEST_TYPE_GET, baseURL, nil)

	if err != nil {
		fmt.Println("Error creating request:", err)
		return "", err
	}

	// 设置请求头
	req.Header.Set("AppKey", appKey)
	req.Header.Set("Nonce", nonce)
	req.Header.Set("CurTime", curTime)
	req.Header.Set("CheckSum", checkSum)
	req.Header.Set("Content-Type", "application/json;charset=utf-8") //v10版本是json， v9版本是application/x-www-form-urlencoded;charset=utf-8
	// 发起请求
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		fmt.Println("Error sending request:", err)
		return "", err
	}
	defer resp.Body.Close()

	// 读取响应体
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Println("Error reading response body:", err)
		return "", err
	}

	// 直接作为字符串输出（即原始 JSON）
	responseJSON := string(body)

	// 打印 JSON 字符串
	log.Info("baseURL:", baseURL)
	log.Info("Response Status:", resp.Status)
	log.Info("Response JSON String:", responseJSON)
	return responseJSON, nil
}

func CalculateCheckSum(appSecret, nonce, curTime string) string {
	// 拼接字符串
	raw := appSecret + nonce + curTime
	// 计算 SHA1 哈希
	h := sha1.New()
	h.Write([]byte(raw))
	// 取哈希值并转为小写十六进制字符串
	return strings.ToLower(hex.EncodeToString(h.Sum(nil)))
}

func SendYunXinSms(appKey, appSecret string, mobile string, templateID int64, customCode string) error {
	// 接口地址
	targetURL := "https://sms.yunxinapi.com/sms/sendcode.action"

	// 构建请求体
	params := map[string]interface{}{
		"mobile":     mobile,
		"templateid": templateID,
	}
	if customCode != "" {
		params["codeLen"] = len(customCode)
		params["paramMap"] = fmt.Sprintf(`{"code":"%s"}`, customCode)
	}

	// 发送 POST 请求
	resp, err := SendYunxinPostApiWithForm(appKey, appSecret, targetURL, params)
	if err != nil {
		fmt.Println("Error sending YunXinSms request:", err)
		return err
	}
	log.Info("baseURL:", targetURL)
	log.Info("YunXinSms response:", resp)
	return nil
}

// 1：单聊会话。2：高级群会话。3：超大群会话
const conversationType = "1"

func MakeYunxinConversationId(senderId, receiverId string) string {
	// owner_id | conversation_type | other_id
	return fmt.Sprintf("%s|%s|%s", senderId, conversationType, receiverId)
}

const (
	// 云信消息类型：文本消息
	YUNXIN_MSG_TYPE_TXT int = 0
	// 云信消息类型：图片消息
	YUNXIN_MSG_TYPE_PIC int = 1
	// 云信消息类型：语音消息
	YUNXIN_MSG_TYPE_VOICE int = 2
	// 云信消息类型：视频消息
	YUNXIN_MSG_TYPE_VIDEO int = 3
	// 云信消息类型：地理位置消息
	YUNXIN_MSG_TYPE_LOCATION int = 4
	// 云信消息类型：文件消息
	YUNXIN_MSG_TYPE_FILE int = 6
	// 云信消息类型：提示消息
	YUNXIN_MSG_TYPE_NOTICE int = 10
	// 云信消息类型：自定义消息
	YUNXIN_MSG_TYPE_CUSTOM int = 100
)

type YunxinMsg struct {
	MessageType     int         `json:"message_type"`                // 消息类型
	SubType         int         `json:"sub_type,omitempty"`          // 消息子类型
	Text            string      `json:"text,omitempty"`              // 文本/提示消息内容
	Attachment      interface{} `json:"attachment,omitempty"`        // 非文本消息属性或自定义消息内容
	MessageClientID string      `json:"message_client_id,omitempty"` // 客户端消息ID
}

// 云信sendmsg
func SendYunxinMsg(appKey, appSecret string, userId uint, targetId uint, msgData YunxinMsg) error {
	accid := fmt.Sprintf("%v", userId)
	conversationId := MakeYunxinConversationId(accid, fmt.Sprintf("%v", targetId))
	targetUrl := fmt.Sprintf("https://open.yunxinapi.com/im/v2/conversations/%s/messages", conversationId)
	//msgData YunxinMsg 转成map[string]interface{}
	msgDataMap := map[string]interface{}{
		"message_type": msgData.MessageType,
	}
	if msgData.SubType != 0 {
		msgDataMap["sub_type"] = msgData.SubType
	}
	if msgData.Text != "" {
		msgDataMap["text"] = msgData.Text
	}
	if msgData.Attachment != nil {
		msgDataMap["attachment"] = msgData.Attachment
	}
	if msgData.MessageClientID != "" {
		msgDataMap["message_client_id"] = msgData.MessageClientID
	}

	resp, err := SendYunxinPostApi(appKey, appSecret, targetUrl, map[string]interface{}{
		"accids": fmt.Sprintf(`["%s"]`, accid),
	})
	if err != nil {
		fmt.Println("Error sending request:", err)
		return err
	}
	fmt.Println(resp)
	return nil
}

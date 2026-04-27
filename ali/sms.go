package ali

import (
	"encoding/json"
	"errors"

	openapi "github.com/alibabacloud-go/darabonba-openapi/client"
	dysmsapi "github.com/alibabacloud-go/dysmsapi-20170525/v2/client"
	console "github.com/alibabacloud-go/tea-console/client"
	util "github.com/alibabacloud-go/tea-utils/service"
	"github.com/alibabacloud-go/tea/tea"
	"github.com/labstack/gommon/log"
)

// 使用AK&SK初始化账号Client
func CreateSmsClient(accessKeyId *string, accessKeySecret *string) (_result *dysmsapi.Client, _err error) {
	config := &openapi.Config{}
	config.AccessKeyId = accessKeyId
	config.AccessKeySecret = accessKeySecret
	_result = &dysmsapi.Client{}
	_result, _err = dysmsapi.NewClient(config)
	return _result, _err
}

type SmsResponse struct {
	Code      string `json:"Code"`      //返回 OK 代表请求成功。 // 其他值代表请求失败。
	Message   string `json:"Message"`   //状态码的描述。
	BizId     string `json:"BizId"`     //发送回执 ID。可根据发送回执 ID 在接口 QuerySendDetails 中查询具体的发送状态。
	RequestId string `json:"RequestId"` //请求 ID。
}

type codeParam struct {
	Code string `json:"code"`
}

func (c *codeParam) String() string {
	//转化为json字符串
	data, err := json.Marshal(c)
	if err != nil {
		return ""
	}
	return string(data)
}

func SendAliSms(accessKeyId, accessKeySecret, phoneNum, signName, templateCode, smsCode string) (_err error) {
	client, _err := CreateSmsClient(&accessKeyId, &accessKeySecret)
	if _err != nil {
		log.Error("CreateSmsClient error: %v", _err)
		return _err
	}

	codeParam := &codeParam{}
	codeParam.Code = smsCode
	// 1.发送短信
	sendReq := &dysmsapi.SendSmsRequest{
		PhoneNumbers:  tea.String(phoneNum),
		SignName:      tea.String(signName),
		TemplateCode:  tea.String(templateCode),
		TemplateParam: tea.String(codeParam.String()),
	}
	log.Info("SendSms request: ", sendReq)

	sendResp, _err := client.SendSms(sendReq)
	if _err != nil {
		log.Error("SendSms error: ", _err)
		return _err
	}

	code := sendResp.Body.Code
	if !tea.BoolValue(util.EqualString(code, tea.String("OK"))) {
		console.Log(tea.String("错误信息: " + tea.StringValue(sendResp.Body.Message)))
		_err = errors.New(tea.StringValue(sendResp.Body.Message))
		return _err
	}

	SendId := sendResp.Body.BizId
	console.Log(tea.String("发送回执 ID: " + tea.StringValue(SendId)))
	// 2. 等待 10 秒后查询结果
	_err = util.Sleep(tea.Int(10000))
	if _err != nil {
		return _err
	}
	return _err
}

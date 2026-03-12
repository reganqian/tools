package ali

import (
	"errors"

	openapi "github.com/alibabacloud-go/darabonba-openapi/client"
	dypnsapi "github.com/alibabacloud-go/dypnsapi-20170525/client"
	console "github.com/alibabacloud-go/tea-console/client"
	util "github.com/alibabacloud-go/tea-utils/v2/service"
	"github.com/alibabacloud-go/tea/tea"
)

// 使用AK&SK初始化账号Client
func CreateAliClient(accessKeyId *string, accessKeySecret *string) (_result *dypnsapi.Client, _err error) {
	config := &openapi.Config{}
	config.AccessKeyId = accessKeyId
	config.AccessKeySecret = accessKeySecret
	_result = &dypnsapi.Client{}
	_result, _err = dypnsapi.NewClient(config)
	return _result, _err
}

// 获取手机号
func AliGetMobile(accessKeyId, accessKeySecret, accessToken string) (resBody *dypnsapi.GetMobileResponseBody, _err error) {
	client, _err := CreateAliClient(tea.String(accessKeyId), tea.String(accessKeySecret))
	if _err != nil {
		return nil, _err
	}

	request := &dypnsapi.GetMobileRequest{}
	request.AccessToken = tea.String(accessToken)
	response, _err := client.GetMobile(request)
	if _err != nil {
		return nil, _err
	}

	code := response.Body.Code
	if !tea.BoolValue(util.EqualString(code, tea.String("OK"))) {
		console.Log(tea.String("错误信息:" + tea.StringValue(response.Body.Message)))
		_err = errors.New(tea.StringValue(response.Body.Message))
		return nil, _err
	}

	console.Log(tea.String("响应结果:" + tea.ToString(response.Body)))
	return response.Body, _err
}

// 手机号校验
func AliVerifyMobile(accessKeyId, accessKeySecret, accessToken, phoneNumber string) (resBody *dypnsapi.VerifyMobileResponseBody, _err error) {
	client, _err := CreateAliClient(tea.String(accessKeyId), tea.String(accessKeySecret))
	if _err != nil {
		return nil, _err
	}

	request := &dypnsapi.VerifyMobileRequest{}
	request.AccessCode = tea.String(accessToken)
	request.PhoneNumber = tea.String(phoneNumber)

	response, _err := client.VerifyMobile(request)
	if _err != nil {
		return nil, _err
	}

	code := response.Body.Code
	if !tea.BoolValue(util.EqualString(code, tea.String("OK"))) {
		console.Log(tea.String("错误信息:" + tea.StringValue(response.Body.Message)))
		_err = errors.New(tea.StringValue(response.Body.Message))
		return nil, _err
	}

	console.Log(tea.String("响应结果:" + tea.ToString(response.Body)))
	return response.Body, _err
}

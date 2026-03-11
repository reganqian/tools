// This file is auto-generated, don't edit it. Thanks.
package ali

import (
	sdk "github.com/alibabacloud-go/cloudauth-20190307/v4/client"
	openapi "github.com/alibabacloud-go/darabonba-openapi/v2/client"
	console "github.com/alibabacloud-go/tea-console/client"
	util "github.com/alibabacloud-go/tea-utils/v2/service"
	"github.com/alibabacloud-go/tea/tea"
	credential "github.com/aliyun/credentials-go/credentials"
)

type ClientInterface interface {
}

type AliyunIdCardCheckReq struct {
	UserName        string `json:"userName"`        //真实姓名
	IdentifyNum     string `json:"identifyNum"`     //证件号码
	AccessKeyId     string `json:"accessKeyId"`     //accessKeyId
	AccessKeySecret string `json:"accessKeySecret"` //accessKeySecret
}

type AliyunIdCardCheckResponse struct {
	CheckResult int `json:"checkResult"` //检验结果, 1:是,  2否
}

const (
	ALIYUN_CHECK_RESULT_YES = 1
	ALIYUN_CHECK_RESULT_NO  = 2
)

func AliyunIdCardCheck(req AliyunIdCardCheckReq) (res *AliyunIdCardCheckResponse, err error) {
	tryErr := func() (_e error) {
		defer func() {
			if r := tea.Recover(recover()); r != nil {
				_e = r
			}
		}()
		// 构建request。
		request := &sdk.Id2MetaVerifyRequest{}
		request.ParamType = tea.String("normal")
		request.UserName = &req.UserName
		request.IdentifyNum = &req.IdentifyNum
		// 自动路由服务。
		response := Id2MetaVerifyAutoRoute(request, req.AccessKeyId, req.AccessKeySecret)
		if tea.BoolValue(util.EqualNumber(tea.ToInt(response.StatusCode), tea.Int(200))) {
			res := &AliyunIdCardCheckResponse{}
			if tea.BoolValue(util.EqualString(response.Body.ResultObject.BizCode, tea.String("1"))) {
				//一致
				res.CheckResult = ALIYUN_CHECK_RESULT_YES
			} else {
				// 不一致
				res.CheckResult = ALIYUN_CHECK_RESULT_NO
			}
		}

		ret := util.ToJSONString(util.ToMap(response))
		console.Log(tea.String("最终结果（若此处为空，则所有服务点均异常，请逐步调试）：" + tea.StringValue(ret)))

		return nil
	}()

	if tryErr != nil {
		return res, tryErr
	}
	return res, err
}

// 主备服务点循环调用，获取到成功结果返回。
func Id2MetaVerifyAutoRoute(request *sdk.Id2MetaVerifyRequest, accessKeyId, accessKeySecret string) (_result *sdk.Id2MetaVerifyResponse) {
	endpoints := []*string{tea.String("cloudauth.cn-shanghai.aliyuncs.com"), tea.String("cloudauth.cn-beijing.aliyuncs.com"), tea.String("cloudauth.aliyuncs.com")}
	var lastResponse *sdk.Id2MetaVerifyResponse
	for _, endpoint := range endpoints {
		_, tryErr := func() (_r *sdk.Id2MetaVerifyResponse, _e error) {
			defer func() {
				if r := tea.Recover(recover()); r != nil {
					_e = r
				}
			}()
			// 调用服务。
			response, _err := Id2MetaVerify(endpoint, request, accessKeyId, accessKeySecret)
			if _err != nil {
				return nil, _err
			}
			// 节点调用结果
			ret := util.ToJSONString(util.ToMap(response))
			console.Log(tea.String("节点 " + tea.StringValue(endpoint) + " 结果：" + tea.StringValue(ret) + " "))
			// 有一个服务调用成功即返回。
			if !tea.BoolValue(util.IsUnset(response)) && tea.BoolValue(util.EqualNumber(tea.ToInt(response.StatusCode), tea.Int(200))) {
				if !tea.BoolValue(util.IsUnset(response.Body)) && tea.BoolValue(util.EqualString(response.Body.Code, tea.String("200"))) {
					lastResponse = response
					return
				}

			}

			return nil, nil
		}()

		if tryErr != nil {
			var error = &tea.SDKError{}
			if _t, ok := tryErr.(*tea.SDKError); ok {
				error = _t
			} else {
				error.Message = tea.String(tryErr.Error())
			}
			console.Error(tea.String("节点 " + tea.StringValue(endpoint) + " 调用异常：" + tea.StringValue(error.Message)))
		}
	}
	_result = lastResponse
	return _result
}

// Description:
//
// 获取服务Client实例，调用验证方法。
func Id2MetaVerify(endpoint *string, request *sdk.Id2MetaVerifyRequest, accessKeyId, accessKeySecret string) (_result *sdk.Id2MetaVerifyResponse, _err error) {
	// 获取SDK Client实例。
	client := CreateClient(endpoint, accessKeyId, accessKeySecret)
	// 构建RuntimeObject
	runtime := &util.RuntimeOptions{}
	runtime.ReadTimeout = tea.Int(5000)
	runtime.ConnectTimeout = tea.Int(5000)
	// 连接
	_result = &sdk.Id2MetaVerifyResponse{}
	_body, _err := client.Id2MetaVerifyWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, nil
}

// Description:
//
// 安全创建服务Client实例。
func CreateClient(endpoint *string, accessKeyId, accessKeySecret string) (_result *sdk.Client) {
	// 获取Credential工具，此工具会从环境变量中读取AccessKey。
	credentialConfig := &credential.Config{
		AccessKeyId:     tea.String(accessKeyId),
		AccessKeySecret: tea.String(accessKeySecret),
	}

	credential, _err := credential.NewCredential(credentialConfig)
	if _err != nil {
		return _result
	}

	// 创建SDK Client实例。
	apiConfig := &openapi.Config{}
	apiConfig.Credential = credential
	apiConfig.Endpoint = endpoint
	_result = &sdk.Client{}
	_result, _err = sdk.NewClient(apiConfig)
	return _result
}

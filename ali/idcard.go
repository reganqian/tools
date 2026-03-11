package ali

import (
	sdk "github.com/alibabacloud-go/cloudauth-20190307/v4/client"
	openapi "github.com/alibabacloud-go/darabonba-openapi/v2/client"
	console "github.com/alibabacloud-go/tea-console/client"
	util "github.com/alibabacloud-go/tea-utils/v2/service"
	"github.com/alibabacloud-go/tea/tea"
)

func AliIdcardCheck(regionId, accessKeyId, accessKeySecret, identifyNum, userName string) (_result *sdk.Id2MetaStandardVerifyResponse, _err error) {
	// 可用区域Id （请自行配置）
	// regionId := args[0]
	// 认证场景ID。您必须先在智能核身控制台创建认证场景，才能获得认证场景ID。
	// sceneId := tea.Int64(10000000000000)
	// 客户服务端自定义的业务唯一标识，用于后续定位排查问题使用。值最长为32位长度的字母数字组合，请确保唯一。
	// outerOrderNo := tea.String(orderNo)
	// 认证方案。唯一取值：LR_FR。
	// productCode := tea.String("LR_FR")
	// 活体检测类型。取值：LIVENESS（默认）：动作活体检测 | PHOTINUS_LIVENESS：动作活体+炫彩活体双重检测
	// model := tea.String("LIVENESS")
	// 证件类型。取值：IDENTITY_CARD，表示身份证。
	// certType := tea.String("IDENTITY_CARD")
	// 用户的真实姓名。
	// certName := tea.String("张三")
	// 用户的证件号码。
	// certNo := tea.String("330103xxxxxxxxxxxx")
	// MetaInfo环境参数，需要通过客户端SDK获取。
	// metaInfo := tea.String(`{"zimVer":"3.0.0","appVersion":"1","bioMetaInfo":"4.1.0:11501568,0","appName":"com.aliyun.antcloudauth","deviceType":"ios","osVersion":"iOS10.3.2","apdidToken":"","deviceModel":"iPhone9,1"}`)
	tryErr := func() (_e error) {
		defer func() {
			if r := tea.Recover(recover()); r != nil {
				_e = r
			}
		}()
		config := &openapi.Config{}
		// config.AccessKeyId = env.GetEnv(tea.String("ACCESS_KEY_ID"))
		// // 您的AccessKey ID
		// config.AccessKeySecret = env.GetEnv(tea.String("ACCESS_KEY_SECRET"))
		// // 您的AccessKey Secret

		config.AccessKeyId = tea.String(accessKeyId)
		config.AccessKeySecret = tea.String(accessKeySecret)

		config.RegionId = tea.String(regionId)
		// 您的可用区ID
		client, _err := sdk.NewClient(config)

		if _err != nil {
			return _err
		}

		request := &sdk.Id2MetaStandardVerifyRequest{}
		request.ParamType = tea.String("normal")
		request.UserName = &userName
		request.IdentifyNum = &identifyNum
		response, _err := client.Id2MetaStandardVerify(request)
		if _err != nil {
			return _err
		}
		_result = response
		console.Log(util.ToJSONString(util.ToMap(response)))

		return nil
	}()

	if tryErr != nil {
		var error = &tea.SDKError{}
		if _t, ok := tryErr.(*tea.SDKError); ok {
			error = _t
		} else {
			error.Message = tea.String(tryErr.Error())
		}
		console.Log(error.Message)
	}
	return _result, _err
}

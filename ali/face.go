package ali

import (
	sdk "github.com/alibabacloud-go/cloudauth-20190307/v4/client"
	console "github.com/alibabacloud-go/tea-console/client"
	util "github.com/alibabacloud-go/tea-utils/v2/service"
	"github.com/alibabacloud-go/tea/tea"
)

type DoFaceVerifyRequest struct {
	SceneId int64 `json:"sceneId"` //认证场景 ID。注意 该字段类型为 Long，在序列化/反序列化的过程中可能导致精度丢失，请注意数值不得大于 9007199254740991。
	// OuterOrderNo string `json:"outerOrderNo"` //商户请求的唯一标识。值为 32 位长度的字母数字组合。前面几位字符是商户自定义的简称，中间可以使用一段时间，后段可以使用一个随机或递增序列。
	// ProductCode  string `json:"productCode"`  //固定值。在不同的产品方案中，该参数值不同：APP 认证方案：参数固定值为 ID_PRO;活体人脸验证方案：参数固定值为 PV_FV;活体检测方案：参数固定值为 LR_FR
	// MetaInfo string `json:"metaInfo"` //Metainfo 环境参数，需要通过客户端 SDK 获取。
	// Ip                     string `json:"ip"`                     //用户 IP
	UserId                 string `json:"userId"`                 //客户业务自定义的用户 ID，请保持唯一。
	FaceContrastPictureUrl string `json:"faceContrastPictureUrl"` //OSS 照片地址，目前只支持已授权 OSS 照片地址。
	AccessKeyId            string `json:"accessKeyId"`            //accessKeyId
	AccessKeySecret        string `json:"accessKeySecret"`        //accessKeySecret
	// FaceContrastPicture    string `json:"faceContrastPicture"`    //照片 Base64 编码。
	// CertifyId              string `json:"certifyId"`              //之前实人认证通过的 CertifyId，认证时的照片作为比对照片。
	// OssBucketName          string `json:"OssBucketName"`          //
	// OssObjectName          string `json:"OssObjectName"`          //
	//说明 活体检测类型仅支持下列取值，暂不支持自定义动作或组合。
	// LIVENESS（默认）：眨眼
	// PHOTINUS_LIVENESS：眨眼+ 炫彩
	// MULTI_ACTION：眨眼 + 摇头 （眨眼和摇头顺序随机）
	// MOVE_ACTION（推荐）：远近移动 + 眨眼
	// MOVE_PHOTINUS：远近移动 + 炫彩
	// 说明
	// 默认的活体检测类型在下列版本中支持：
	// Android SDK1.2.6 及以上
	// iOS SDK1.2.4 及以上
	// Harmony SDK1.0.0 及以上;其他类型都在 Android/iOS/Harmony 的最新 SDK 版本上支持，建议集成最新版本使用。
	// Model string `json:"model"` //活体检测类型，取值：

}

// 返回实体
type DoFaceVerifyResponse struct {
	// Unique identifier for real-person authentication.
	//
	// example:
	//
	// 91707dc296d469ad38e4c5efa6a0f24b
	CertifyId *string `json:"certifyId,omitempty" xml:"CertifyId,omitempty"`
	// URL for real-person authentication in a Web browser, which will redirect according to the ReturnUrl parameter after authentication.
	//
	// 	Notice:
	//
	// - The CertifyUrl returned by the initialization interface is valid for **30 minutes and can only be used once**. Please use it within the validity period to avoid reuse.
	//
	// - This parameter requires the correct input of **MetaInfo*	- to return a CertifyUrl that matches the client. If you cannot obtain it, please check whether **MetaInfo*	- and other input parameters are correct.
	//
	// - The domain name of this URL may change with service updates. To ensure normal service availability, it is recommended not to apply access control to this domain name.
	//
	// - When redirecting in the browser, try not to use incognito mode or modify the URL, as this may result in a **signature error**.
	//
	// example:
	//
	// https://t.aliyun.com/****
	CertifyUrl *string `json:"certifyUrl,omitempty" xml:"CertifyUrl,omitempty"`
}

func DoFaceVerify(req DoFaceVerifyRequest) (result *sdk.InitFaceVerifyResponse, err error) {
	request := sdk.InitFaceVerifyRequest{}
	request.UserId = &req.UserId
	request.FaceContrastPictureUrl = &req.FaceContrastPictureUrl
	request.SceneId = &req.SceneId
	result = InitFaceVerifyAutoRoute(&request, req.AccessKeyId, req.AccessKeySecret)
	// if tea.BoolValue(util.EqualNumber(tea.ToInt(result.StatusCode), tea.Int(200))) {
	// 	res := &DoFaceVerifyResponse{}
	// 	res.CertifyId = result.Body.ResultObject.CertifyId
	// 	res.CertifyUrl = result.Body.ResultObject.CertifyUrl
	// } else {
	// 	return res, err
	// }

	return result, err
}

func InitFaceVerifyAutoRoute(request *sdk.InitFaceVerifyRequest, accessKeyId, accessKeySecret string) (_result *sdk.InitFaceVerifyResponse) {
	endpoints := []*string{tea.String("cloudauth.cn-shanghai.aliyuncs.com"), tea.String("cloudauth.cn-beijing.aliyuncs.com")}
	var lastResponse *sdk.InitFaceVerifyResponse
	for _, endpoint := range endpoints {
		_, tryErr := func() (_r *sdk.Id2MetaVerifyResponse, _e error) {
			defer func() {
				if r := tea.Recover(recover()); r != nil {
					_e = r
				}
			}()
			// 调用服务。
			response := InitFaceVerify(endpoint, request, accessKeyId, accessKeySecret)
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

// 获取服务Client实例，调用验证方法。
func InitFaceVerify(endpoint *string, request *sdk.InitFaceVerifyRequest, accessKeyId, accessKeySecret string) (_result *sdk.InitFaceVerifyResponse) {
	// 获取SDK Client实例。
	client := CreateClient(endpoint, accessKeyId, accessKeySecret)
	// 连接
	_result = &sdk.InitFaceVerifyResponse{}
	_body, _err := client.InitFaceVerify(request)
	if _err != nil {
		return _result
	}
	_result = _body
	return _result
}

type DoDescribeFaceVerifyReq struct {
	// Unique identifier for real-person authentication.
	//
	// example:
	//
	// 91707dc296d469ad38e4c5efa6a0f24b
	CertifyId *string `json:"CertifyId,omitempty" xml:"CertifyId,omitempty"`
	// Image return type.
	//
	// example:
	//
	// JPG
	PictureReturnType *string `json:"PictureReturnType,omitempty" xml:"PictureReturnType,omitempty"`
	// Authentication scene ID.
	//
	// example:
	//
	// 1000000006
	SceneId *int64 `json:"SceneId,omitempty" xml:"SceneId,omitempty"`

	AccessKeyId     string `json:"accessKeyId"`     //AccessKeyId
	AccessKeySecret string `json:"accessKeySecret"` //AccessKeySecret
}

func DoDescribeFaceVerify(req DoDescribeFaceVerifyReq) (result *sdk.DescribeFaceVerifyResponse, err error) {
	request := sdk.DescribeFaceVerifyRequest{}
	request.CertifyId = req.CertifyId
	request.PictureReturnType = req.PictureReturnType
	request.SceneId = req.SceneId
	// request.UserId = &req.UserId
	// request.FaceContrastPictureUrl = &req.FaceContrastPictureUrl
	result = DescribeFaceVerifyAutoRoute(&request, req.AccessKeyId, req.AccessKeySecret)
	// if tea.BoolValue(util.EqualNumber(tea.ToInt(result.StatusCode), tea.Int(200))) {
	// 	res = result.Body
	// } else {
	// 	return res, err
	// }

	return result, err
}

func DescribeFaceVerifyAutoRoute(request *sdk.DescribeFaceVerifyRequest, accessKeyId, accessKeySecret string) (_result *sdk.DescribeFaceVerifyResponse) {
	endpoints := []*string{tea.String("cloudauth.cn-shanghai.aliyuncs.com"), tea.String("cloudauth.cn-beijing.aliyuncs.com")}
	var lastResponse *sdk.DescribeFaceVerifyResponse
	for _, endpoint := range endpoints {
		_, tryErr := func() (_r *sdk.Id2MetaVerifyResponse, _e error) {
			defer func() {
				if r := tea.Recover(recover()); r != nil {
					_e = r
				}
			}()
			// 调用服务。
			response := DescribeFaceVerify(endpoint, request, accessKeyId, accessKeySecret)
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

// 获取服务Client实例，调用验证方法。
func DescribeFaceVerify(endpoint *string, request *sdk.DescribeFaceVerifyRequest, accessKeyId, accessKeySecret string) (_result *sdk.DescribeFaceVerifyResponse) {
	// 获取SDK Client实例。
	client := CreateClient(endpoint, accessKeyId, accessKeySecret)
	// 连接
	_result = &sdk.DescribeFaceVerifyResponse{}
	_body, _err := client.DescribeFaceVerify(request)
	if _err != nil {
		return _result
	}
	_result = _body
	return _result
}

package tx

import (
	"fmt"
	"strconv"

	"github.com/labstack/gommon/log"
	captcha "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/captcha/v20190722"
	"github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common"
	"github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common/errors"
	"github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common/profile"
)

// VerifyCaptchaTicket 校验腾讯云验证码票据
func VerifyCaptchaTicket(
	secretID string,
	secretKey string,
	appID string,
	appSecretKey string,
	ticket string,
	randStr string,
	userIP string,

) (bool, error) {

	credential := common.NewCredential(secretID, secretKey)

	cpf := profile.NewClientProfile()
	cpf.HttpProfile.Endpoint = "captcha.tencentcloudapi.com"

	client, err := captcha.NewClient(credential, "", cpf)
	if err != nil {
		return false, err
	}
	captchaType := uint64(9)
	request := captcha.NewDescribeCaptchaResultRequest()
	request.CaptchaType = &captchaType
	request.Ticket = &ticket
	request.Randstr = &randStr
	request.UserIp = &userIP

	//string转 uint64
	appIDInt, err := strconv.ParseUint(appID, 10, 64)
	if err != nil {
		return false, err
	}

	request.CaptchaAppId = &appIDInt
	// request.BusinessId = &bizID
	// request.SceneId = common.Int64Ptr(0)
	request.AppSecretKey = common.StringPtr(appSecretKey) // 如未配置可不填

	response, err := client.DescribeCaptchaResult(request)
	log.Info("DescribeCaptchaResult response: ", response)
	if err != nil {
		if sdkErr, ok := err.(*errors.TencentCloudSDKError); ok {
			return false, fmt.Errorf("sdk error: %s", sdkErr.GetMessage())
		}
		return false, err
	}

	// CaptchaCode == 1 表示校验成功
	if response.Response != nil &&
		response.Response.CaptchaCode != nil &&
		*response.Response.CaptchaCode == 1 {
		return true, nil
	}

	return false, nil
}

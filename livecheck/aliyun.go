package livecheck

/**
最低SDK版本要求：facebody-20191230的SDK版本需大于等于4.0.7。
可以在此仓库地址中引用最新版本SDK：https://pkg.go.dev/github.com/alibabacloud-go/facebody-20191230/v4
依赖github.com/alibabacloud-go/facebody-20191230
建议使用go mod tidy安装依赖
*/

import (
	"fmt"
	"net/http"

	openapi "github.com/alibabacloud-go/darabonba-openapi/v2/client"
	facebody20191230 "github.com/alibabacloud-go/facebody-20191230/v4/client"
	util "github.com/alibabacloud-go/tea-utils/v2/service"
	"github.com/alibabacloud-go/tea/tea"
)

// const (
// 	accessKeyId     = "LTAI5t999999999999999"
// 	accessKeySecret = "99999999999999999999999999999999"
// )

// 单图片
func AliyunLiveCheck(accessKeyId, accessKeySecret, fileUrl string) {
	// 创建AccessKey ID和AccessKey Secret，请参考https://help.aliyun.com/document_detail/175144.html。
	// 如果您用的是RAM用户的AccessKey，还需要为RAM用户授予权限AliyunVIAPIFullAccess，请参考https://help.aliyun.com/document_detail/145025.html。
	// 从环境变量读取配置的AccessKey ID和AccessKey Secret。运行示例前必须先配置环境变量。
	// accessKeyId := os.Getenv("ALIBABA_CLOUD_ACCESS_KEY_ID")
	// accessKeySecret := os.Getenv("ALIBABA_CLOUD_ACCESS_KEY_SECRET")
	// 初始化配置对象 &openapi.Config。Config对象存放AccessKeyId、AccessKeySecret、Endpoint等配置。
	config := &openapi.Config{
		AccessKeyId:     tea.String(accessKeyId),
		AccessKeySecret: tea.String(accessKeySecret),
	}
	// 访问的域名
	config.Endpoint = tea.String("facebody.cn-shanghai.aliyuncs.com")
	client, err := facebody20191230.NewClient(config)
	if err != nil {
		panic(err)
	}

	//场景二，使用任意可访问的url
	httpClient := http.Client{}
	file1, _ := httpClient.Get(fileUrl)

	tasks1 := &facebody20191230.DetectLivingFaceAdvanceRequestTasks{
		ImageURLObject: file1.Body,
	}
	detectLivingFaceAdvanceRequest := &facebody20191230.DetectLivingFaceAdvanceRequest{
		Tasks: []*facebody20191230.DetectLivingFaceAdvanceRequestTasks{tasks1},
	}
	runtime := &util.RuntimeOptions{}
	detectLivingFaceResponse, err := client.DetectLivingFaceAdvance(detectLivingFaceAdvanceRequest, runtime)
	if err != nil {
		// 获取整体报错信息
		fmt.Println(err.Error())
	} else {
		// 获取整体结果
		fmt.Println(detectLivingFaceResponse)
	}
}

// 视频
func AliyunVideoLiveCheck(accessKeyId, accessKeySecret, fileUrl string) {
	// 创建AccessKey ID和AccessKey Secret，请参考https://help.aliyun.com/document_detail/175144.html。
	// 如果您用的是RAM用户的AccessKey，还需要为RAM用户授予权限AliyunVIAPIFullAccess，请参考https://help.aliyun.com/document_detail/145025.html。
	// 从环境变量读取配置的AccessKey ID和AccessKey Secret。运行示例前必须先配置环境变量。
	// accessKeyId := os.Getenv("ALIBABA_CLOUD_ACCESS_KEY_ID")
	// accessKeySecret := os.Getenv("ALIBABA_CLOUD_ACCESS_KEY_SECRET")
	// 初始化配置对象 &openapi.Config。Config对象存放AccessKeyId、AccessKeySecret、Endpoint等配置。
	config := &openapi.Config{
		AccessKeyId:     tea.String(accessKeyId),
		AccessKeySecret: tea.String(accessKeySecret),
	}
	// 访问的域名
	config.Endpoint = tea.String("facebody.cn-shanghai.aliyuncs.com")
	client, err := facebody20191230.NewClient(config)
	if err != nil {
		panic(err)
	}

	detectVideoLivingFaceRequest := &facebody20191230.DetectVideoLivingFaceRequest{
		VideoUrl: tea.String(fileUrl),
	}

	//DetectVideoLivingFace
	detectLivingFaceResponse, err := client.DetectVideoLivingFace(detectVideoLivingFaceRequest)
	if err != nil {
		// 获取整体报错信息
		fmt.Println(err.Error())
	} else {
		// 获取整体结果
		fmt.Println(detectLivingFaceResponse)
	}
}

//1:1
func AliyunPicCompare(accessKeyId, accessKeySecret, fileA, fileB string) {
	// 创建AccessKey ID和AccessKey Secret，请参考https://help.aliyun.com/document_detail/175144.html。
	// 如果您用的是RAM用户的AccessKey，还需要为RAM用户授予权限AliyunVIAPIFullAccess，请参考https://help.aliyun.com/document_detail/145025.html。
	// 从环境变量读取配置的AccessKey ID和AccessKey Secret。运行示例前必须先配置环境变量。
	// accessKeyId := os.Getenv("ALIBABA_CLOUD_ACCESS_KEY_ID")
	// accessKeySecret := os.Getenv("ALIBABA_CLOUD_ACCESS_KEY_SECRET")
	// 初始化配置对象 &openapi.Config。Config对象存放AccessKeyId、AccessKeySecret、Endpoint等配置。
	config := &openapi.Config{
		AccessKeyId:     tea.String(accessKeyId),
		AccessKeySecret: tea.String(accessKeySecret),
	}
	// 访问的域名
	config.Endpoint = tea.String("facebody.cn-shanghai.aliyuncs.com")
	client, err := facebody20191230.NewClient(config)
	if err != nil {
		panic(err)
	}

	detectVideoLivingFaceRequest := &facebody20191230.CompareFaceRequest{
		QualityScoreThreshold: tea.Float32(98.5), // 质量分阈值
		ImageURLA:             tea.String(fileA), // 图片A
		ImageURLB:             tea.String(fileB), // 图片B
	}

	//DetectVideoLivingFace
	response, err := client.CompareFace(detectVideoLivingFaceRequest)
	if err != nil {
		// 获取整体报错信息
		fmt.Println(err.Error())
	} else {
		// 获取整体结果
		fmt.Println(response)
	}
}

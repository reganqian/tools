package yx

import (
	"fmt"
	"time"
)

// 创建房间
func CreateRtcRoom(appKey, appSecret string, uid, targetUid uint) {
	//
	// 参数	类型	是否必选	示例	描述
	// channelName	String	必选	abc	房间名称。
	// 字符串格式，长度为 1~64 字节。
	// 支持以下字符类型：
	// 大写英文字母
	// 小写英文字母
	// 数字
	// 空格
	// 特殊字符，支持!#$%&()+-:;≤.,>? @[]^_{|}~”
	// mode	Integer	必选	2	固定为 2。
	// uid	Long	必选	163	房间创建者的用户 ID，是您的业务系统中的实际用户 ID。
	// 长度限制为 64 位二进制。
	// 实现创建房间接口
	channelName := MakeRtcChannelName(uid, targetUid)
	// 构造请求参数
	params := map[string]interface{}{
		"channelName": channelName,
		"mode":        2,
		"uid":         uid,
	}
	baseUrl := "https://rtc.yunxinapi.com/v2/api/room"
	// 发送 POST 请求到云信 RTC 服务
	SendYunxinPostApi(appKey, appSecret, baseUrl, params)
}

func MakeRtcChannelName(uid, targetUid uint) string {
	//数字小的在前面
	if uid > targetUid {
		uid, targetUid = targetUid, uid
	}

	channelName := fmt.Sprintf("%v-%v", uid, targetUid)

	return channelName
}

func MakeChannelNameWithTime(uid, targetUid uint) string {
	//数字小的在前面
	if uid > targetUid {
		uid, targetUid = targetUid, uid
	}

	channelName := fmt.Sprintf("%v-%v-%v", uid, targetUid, time.Now().Unix())

	return channelName
}

func GetYunxinToken(appkey, appSecret, channelName string, uid uint64, ttlSec int) (string, error) {
	// 建议项目初始化时候调用一次 NewTokenServer， 通过单例全局维护一个 TokenServer 对象
	tokenServer, err := NewTokenServer(appkey, appSecret, 7200)
	if err != nil {
		return "", err
	}
	// 在需要的时候，提供 channelName（房间名）、uid（用户标识）、ttlSec（有效时间，单位秒） 参数，生成 token
	token, err := tokenServer.GetToken(channelName, uid, ttlSec)
	if err != nil {
		return "", err
	}

	return token, nil
}

// 获取云信音视频permissiontoken
func GetYunxinPermissionToken(appkey, appSecret, channelName, permSecret string, uid uint64) (string, error) {
	// 建议项目初始化时候调用一次 NewTokenServer， 通过单例全局维护一个 TokenServer 对象
	// appKey、appSecret 请替换成自己的，具体在云信管理后台查看。
	// 7200 代表默认有效的时间，单位秒。不能超过 86400，即 24 小时
	tokenServer, err := NewTokenServer(appkey, appSecret, 7200)
	if err != nil {
		return "", err
	}
	// 在需要的时候，提供 channelName（房间名）、uid（用户标识）、ttlSec（有效时间，单位秒） 参数，生成 token
	// token, err := tokenServer.GetToken(channelName, uid, ttlSec)
	// if err != nil {
	// 	return err
	// }
	// 具体权限说明见函数注释
	privilege := uint8(1)
	ttlSec := int64(1800)
	// permSecret 见云信管理后台，具体见文档说明：https://doc.yunxin.163.com/nertc/server-apis/DU3Mjk0MzQ?platform=server
	permissionToken, err := tokenServer.GetPermissionKey(channelName, permSecret, uid, privilege, ttlSec)

	return permissionToken, err
}

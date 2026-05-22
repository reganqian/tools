package tx

import (
	"fmt"

	"github.com/tencentyun/tls-sig-api-v2-golang/tencentyun"
)

type TxImConf struct {
	SDKAppID  int    `json:"SDKAppID"`  // 必填，应用ID
	APIHost   string `json:"APIHost"`   // 必填，API地址
	SDKAppKey string `json:"SDKAppKey"` // 必填，应用密钥
	AdminUser string `json:"AdminUser"` // 必填，管理员账号
}

func GetTxSig(conf *TxImConf, userId string) (string, error) {

	sig, err := tencentyun.GenUserSig(conf.SDKAppID, conf.SDKAppKey, userId, 86400*180)
	if err != nil {
		fmt.Println(err.Error())
		return "", err
	} else {
		fmt.Println(sig)
	}

	return sig, nil
}

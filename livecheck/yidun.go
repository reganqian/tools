/**
 * 单活体检测接口示例代码
 * 接口文档: https://support.dun.163.com/documents/391676076156063744?docId=456535635703713792
 */

package livecheck

import (
	"crypto/md5"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"io"
	"math/rand"
	"net/http"
	"net/url"
	"sort"
	"strconv"
	"strings"
	"time"

	"github.com/tjfoc/gmsm/sm3"
)

const (
	apiURL     = "https://verify.dun.163.com/v1/liveperson/recheck"
	version    = "v1"
	secretID   = "your_secretId"   //产品密钥ID，产品标识
	secretKey  = "your_secretKey"  //产品私有密钥，服务端生成签名信息使用，请严格保管，避免泄露
	businessID = "your_businessId" //业务ID，易盾根据产品业务特点分配
)

// 请求易盾接口
func check(params url.Values) (reply YidunLivePersionCheckReply) {
	params["secretId"] = []string{secretID}
	params["businessId"] = []string{businessID}
	params["version"] = []string{version}
	params["timestamp"] = []string{strconv.FormatInt(time.Now().UnixNano()/1000000, 10)}
	params["nonce"] = []string{strconv.FormatInt(rand.New(rand.NewSource(time.Now().UnixNano())).Int63n(10000000000), 10)}
	params["signature"] = []string{genSignature(params)}

	resp, err := http.Post(apiURL, "application/x-www-form-urlencoded", strings.NewReader(params.Encode()))

	if err != nil {
		fmt.Println("调用API接口失败:", err)
		reply.Code = 500
		reply.Message = "调用API接口失败:" + err.Error()
		return reply
	}

	defer resp.Body.Close()

	contents, _ := io.ReadAll(resp.Body)
	err = json.Unmarshal(contents, &reply)
	if err != nil {
		fmt.Println("解析JSON失败:", err)
		reply.Code = 500
		reply.Message = "解析JSON失败:" + err.Error()
		return reply
	}
	fmt.Println("解析JSON成功:", reply)
	return reply
}

// 生成签名信息
func genSignature(params url.Values) string {
	var paramStr string
	keys := make([]string, 0, len(params))
	for k := range params {
		keys = append(keys, k)
	}
	sort.Strings(keys)
	for _, key := range keys {
		paramStr += key + params[key][0]
	}
	paramStr += secretKey
	if params["signatureMethod"] != nil && params["signatureMethod"][0] == "SM3" {
		sm3Reader := sm3.New()
		sm3Reader.Write([]byte(paramStr))
		return hex.EncodeToString(sm3Reader.Sum(nil))
	}
	md5Reader := md5.New()
	md5Reader.Write([]byte(paramStr))
	return hex.EncodeToString(md5Reader.Sum(nil))
}

type YidunLivePersionCheckReply struct {
	Code    int                         `json:"code"`
	Message string                      `json:"message"`
	Result  YidunLivePersionCheckResult `json:"result"`
}

type YidunLivePersionCheckResult struct {
	LpCheckStatus string   `json:"lpCheckStatus"` //检测结果，1-通过 2-不通过 3-查无结果
	TaskId        string   `json:"taskId"`        //任务ID
	Status        string   `json:"status"`        //检测状态
	ReasonType    string   `json:"reasonType"`    //原因详情，1-通过 2-活体检测不通过 4-云端检测结果请求超时,请重试 5-云端检测图片上传失败,请重试 7-检测超时或其他异常 10-疑似攻击，建议拦截 11-检测对象为未成年人 13-人脸被遮挡（h5活体暂不返回）
	Avatar        string   `json:"avatar"`        //抓取头像照片
	PicType       string   `json:"picType"`       //图片类型
	PicUrl        []string `json:"picUrl"`        //图片URL
	PicUrlType    string   `json:"picUrlType"`    //图片URL类型
}

func YidunLivePersionCheck() (reply YidunLivePersionCheckReply) {
	params := url.Values{
		"token":      []string{"01421147bf6631522fc0d38b123456"},
		"needAvatar": []string{"false"},
	}

	reply = check(params)

	return reply
}

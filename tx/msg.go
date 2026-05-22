package tx

import (
	"bytes"
	"encoding/json"
	"math/rand"
	"net/http"
)

// MsgElem 单条消息元素（按需扩展）
type MsgElem struct {
	MsgType    string      `json:"MsgType"`
	MsgContent interface{} `json:"MsgContent"`
}

// MsgTextContent 文本消息内容
type MsgTextContent struct {
	Text string `json:"Text"`
}

func FormatContent(content string) (MsgTextContent, error) {
	var textContent MsgTextContent
	err := json.Unmarshal([]byte(content), &textContent)
	if err != nil {
		return textContent, err
	}
	return textContent, nil
}

// // MsgImageContent 图片消息内容
type MsgImageContent struct {
	UUID           string         `json:"UUID"`           // 图片唯一标识符
	ImageFormat    int            `json:"ImageFormat"`    // 图片格式，默认1
	ImageInfoArray []MsgImageInfo `json:"ImageInfoArray"` // 图片信息数组
}

func FormatImageContent(content string) (MsgImageContent, error) {
	var imageContent MsgImageContent
	err := json.Unmarshal([]byte(content), &imageContent)
	if err != nil {
		return imageContent, err
	}
	return imageContent, nil
}

type MsgImageInfo struct {
	Type   int    `json:"Type"`   //大图
	Size   int    `json:"Size"`   //图片大小
	Width  int    `json:"Width"`  //图片宽度
	Height int    `json:"Height"` //图片高度
	URL    string `json:"URL"`    //图片URL
}

// MsgVoiceContent 语音消息内容
type MsgVoiceContent struct {
	URL           string `json:"URL"`           // 语音URL
	UUID          string `json:"UUID"`          // 语音唯一标识符
	Size          int    `json:"Size"`          // 语音大小
	Second        int    `json:"Second"`        // 语音时长
	Download_Flag int    `json:"Download_Flag"` // 下载标志，默认2
}

func FormatVoiceContent(content string) (MsgVoiceContent, error) {
	var voiceContent MsgVoiceContent
	err := json.Unmarshal([]byte(content), &voiceContent)
	if err != nil {
		return voiceContent, err
	}
	return voiceContent, nil
}

// MsgVideoContent 视频消息内容
type MsgVideoContent struct {
	VideoURL          string `json:"VideoURL"`          // 视频URL
	VideoUUID         string `json:"VideoUUID"`         // 视频唯一标识符
	VideoSize         int    `json:"VideoSize"`         // 视频大小
	VideoSecond       int    `json:"VideoSecond"`       // 视频时长
	VideoFormat       string `json:"VideoFormat"`       // 视频格式，默认mp4
	VideoDownloadFlag int    `json:"VideoDownloadFlag"` // 视频下载标志，默认2
	ThumbURL          string `json:"ThumbURL"`          // 缩略图URL
	ThumbUUID         string `json:"ThumbUUID"`         // 缩略图唯一标识符
	ThumbSize         int    `json:"ThumbSize"`         // 缩略图大小
	ThumbWidth        int    `json:"ThumbWidth"`        // 缩略图宽度
	ThumbHeight       int    `json:"ThumbHeight"`       // 缩略图高度
	ThumbFormat       string `json:"ThumbFormat"`       // 缩略图格式，默认JPG
	ThumbDownloadFlag int    `json:"ThumbDownloadFlag"` // 缩略图下载标志，默认2
}

func FormatVideoContent(content string) (MsgVideoContent, error) {
	var videoContent MsgVideoContent
	err := json.Unmarshal([]byte(content), &videoContent)
	if err != nil {
		return videoContent, err
	}
	return videoContent, nil
}

// SendC2CMsgReq 发送单聊消息请求参数（常用字段）
type SendC2CMsgReq struct {
	SyncOtherMachine      int       `json:"SyncOtherMachine,omitempty"` // 是否同步到其他机器，默认0
	FromAccount           string    `json:"From_Account,omitempty"`     // 发送方账号
	ToAccount             string    `json:"To_Account"`                 // 接收方账号
	MsgLifeTime           int       `json:"MsgLifeTime,omitempty"`      // 消息过期时间，默认0，单位秒
	MsgRandom             int       `json:"MsgRandom"`
	MsgSeq                int64     `json:"MsgSeq,omitempty"`
	MsgBody               []MsgElem `json:"MsgBody"`
	MsgType               string    `json:"MsgType,omitempty"`
	CloudCustomData       string    `json:"CloudCustomData,omitempty"`
	ForbidCallbackControl []string  `json:"ForbidCallbackControl,omitempty"`
}

// SendC2CMsgResp 发送单聊消息返回
type SendC2CMsgResp struct {
	ActionStatus string `json:"ActionStatus"`
	ErrorCode    int    `json:"ErrorCode"`
	ErrorInfo    string `json:"ErrorInfo"`
	MsgKey       string `json:"MsgKey"`
	MsgID        string `json:"MsgID"`
}

type SendTxMsgReq struct {
	FromAccount string `json:"From_Account,omitempty"` // 发送方账号
	ToAccount   string `json:"To_Account"`             // 接收方账号
	MsgContent  string `json:"MsgContent"`
	MsgType     string `json:"MsgType,omitempty"`
}

// SendC2CMsg 发送单聊消息（可复用）
func SendC2CMsg(
	conf TxImConf,
	request SendTxMsgReq,
) (*SendC2CMsgResp, error) {
	d := "console.tim.qq.com"

	userSig, err := GetTxSig(&conf, conf.AdminUser)
	if err != nil {
		return nil, err
	}
	random := rand.Int63()

	req := SendC2CMsgReq{}
	req.FromAccount = request.FromAccount
	req.ToAccount = request.ToAccount
	req.MsgType = request.MsgType

	// 校式化消息内容
	content, err := FormatContent(request.MsgContent)
	if err != nil {
		return nil, err
	}

	req.MsgBody = []MsgElem{
		{
			MsgType:    request.MsgType,
			MsgContent: content,
		},
	}

	//
	req.SyncOtherMachine = 2 // 是否同步到其他机器，默认0
	req.ForbidCallbackControl = []string{
		"ForbidBeforeSendMsgCallback",
	}

	body, _ := json.Marshal(req)
	url := "https://" + d + "/v4/openim/sendmsg" +
		"?sdkappid=" + intToStr(conf.SDKAppID) +
		"&identifier=" + conf.AdminUser +
		"&usersig=" + userSig +
		"&random=" + intToStr(int(random)) +
		"&contenttype=json"

	resp, err := http.Post(url, "application/json", bytes.NewReader(body))
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	var result SendC2CMsgResp
	if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
		return nil, err
	}
	return &result, nil
}

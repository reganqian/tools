package email

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"

	"github.com/labstack/gommon/log"
)

type ApiSmtp struct {
	Mine       string         `json:"mime"`
	Channel    string         `json:"channel"`
	Recipients RecipientInfo  `json:"recipients,omitempty"`
	Originator OriginatorInfo `json:"originator,omitempty"`
}

type RecipientInfo struct {
	To       []NameAddr `json:"to,omitempty"`
	Cc       []NameAddr `json:"cc,omitempty"`
	Bcc      []NameAddr `json:"bcc,omitempty"`
	BulkList []NameAddr `json:"bulk_list,omitempty"`
}

type OriginatorInfo struct {
	From    NameAddr `json:"from"`
	ReplyTo NameAddr `json:"reply_to"`
}

type NameAddr struct {
	Name    string `json:"name"`
	Address string `json:"address"`
}

func (s *ApiSmtp) Init(channal, fromName, fromAddr, replyName, replyAddr string) {
	// s.Channel = "BuymmogEdm"
	s.Channel = channal
	org := OriginatorInfo{}
	// org.From = NameAddr{Name: "BuyMMOG", Address: "sender@buymmog.com"}
	// org.ReplyTo = NameAddr{Name: "support", Address: "admin@buymmog"}
	org.From = NameAddr{Name: fromName, Address: fromAddr}
	org.ReplyTo = NameAddr{Name: replyName, Address: replyAddr}
	s.Originator = org
}

func (s *ApiSmtp) SetContent(content string) {
	s.Mine = content
}

func (s *ApiSmtp) SetTo(toEmils []string, bccEmails []string) {
	recipt := RecipientInfo{}
	toList := []NameAddr{}
	for _, toEmail := range toEmils {
		toInfo := NameAddr{}
		toInfo.Address = toEmail
		toInfo.Name = GetEmailName(toEmail)
		toList = append(toList, toInfo)
	}
	bccList := []NameAddr{}
	for _, bccEmail := range bccEmails {
		bccInfo := NameAddr{}
		bccInfo.Address = bccEmail
		bccInfo.Name = GetEmailName(bccEmail)
		bccList = append(bccList, bccInfo)
	}
	recipt.To = toList
	recipt.Bcc = bccList
	s.Recipients = recipt
}

func GetEmailName(email string) (name string) {
	dataList := strings.Split(email, "@")
	if len(dataList) < 1 {
		return email
	}
	name = dataList[0]
	return name
}

type EmailApiInfo struct {
	LimitStr  string `json:"limitStr"`
	Channel   string `json:"channel"`
	FromName  string `json:"fromName"`
	FromAddr  string `json:"fromAddr"`
	ReplyName string `json:"replyName"`
	ReplyAddr string `json:"replyAddr"`
	ApiCode   string `json:"apiCode"`
}

func SendByApi(toEmails, bccEmails []string, content string, api EmailApiInfo) (map[string]interface{}, error) {
	apismtp := ApiSmtp{}
	// apismtp.Init("BuyMMOG", "sender@buymmog.com", "support", "admin@buymmog")
	apismtp.Init(api.Channel, api.FromName, api.FromAddr, api.ReplyName, api.ReplyAddr)
	apismtp.SetTo(toEmails, bccEmails)
	apismtp.SetContent(content)

	jsonData, err := json.Marshal(apismtp)
	if err != nil {
		return nil, err
	}
	log.Info("json params========", string(jsonData))
	// apiCode := "e19ddbea165786557dd2a0db0cf3d8ef3630829d"
	apiCode := api.ApiCode
	httpposturl := fmt.Sprintf("https://api:%v@api.smtp.com/v4/messages/mime", apiCode)
	// httpposturl := "https://api:e19ddbea165786557dd2a0db0cf3d8ef3630829d@api.smtp.com/v4/messages/mime"
	request, err := http.NewRequest("POST", httpposturl, bytes.NewBuffer(jsonData))
	request.Header.Set("Content-Type", "application/json; charset=UTF-8")
	request.Header.Set("accept", "application/json")

	client := &http.Client{}
	response, err := client.Do(request)
	if err != nil {
		fmt.Println("client.Do error ===", err.Error())
		return nil, err
	}
	defer response.Body.Close()

	fmt.Println("response Status:", response.Status)
	fmt.Println("response Headers:", response.Header)
	body, _ := ioutil.ReadAll(response.Body)
	fmt.Println("response Body==", string(body))

	stringBody := string(body)
	log.Info("response body=", stringBody)
	// Log.Info(stringBody)
	var result map[string]interface{}
	if err == nil {
		err = json.Unmarshal(body, &result)
		return nil, err
	}
	if result["status"] == nil {
		return nil, errors.New("send email failed")
	}
	if result["status"].(string) != "success" {
		errmsg := fmt.Sprintf("%v", result["data"])
		return nil, errors.New(errmsg)
	}

	return nil, err
}

func DoApiSendEmail(apiList []EmailApiInfo, to, bcc, message string) (isSend bool, err error) {
	for _, emailApi := range apiList {
		isSend, err = DoApiSendEmailDetail(emailApi, to, bcc, message)
		if isSend {
			return isSend, err
		}
	}
	return isSend, err
}

func DoApiSendEmailDetail(emailApi EmailApiInfo, to, bcc, message string) (isSend bool, err error) {
	to = strings.ToLower(to)
	if !strings.Contains(to, "@") { //如果不包含， 则邮箱格式不正确， 直接退出
		return isSend, err
	}
	toArray := strings.Split(to, "@")
	if len(toArray) != 2 {
		return isSend, err
	}
	suffix := toArray[1]
	if strings.Contains(suffix, emailApi.LimitStr) { //是否走此api流程， 域名判断
		toList := []string{to}
		bccList := []string{}
		if bcc != "" {
			bccList = strings.Split(bcc, ",")
			toList = append(toList, bccList...)
		}

		data, err := SendByApi(toList, bccList, message, emailApi)
		if err != nil {
			log.Error("send error:", err)
			return isSend, err
		} else {
			isSend = true
			log.Info("send SUCCESS", data)
			return isSend, err
		}
	}

	return isSend, err
}

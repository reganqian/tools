package email

import (
	// "fmt"
	"crypto/tls"
	"fmt"
	"net"
	"net/smtp"
	"strings"

	// "strconv"
	// "encoding/base64"

	"github.com/emersion/go-sasl"
	smtp1 "github.com/emersion/go-smtp"
	"github.com/labstack/gommon/log"
)

func SendSmtpMail2(user, password, host, port, to, subject, body, mailtype string) error {
	auth := sasl.NewPlainClient("", user, password)
	msg := strings.NewReader(
		"From: " + user + "\r\n" +
			"To: " + to + "\r\n" +
			"Subject: " + subject + "\r\n" +
			"\r\n" +
			body + "\r\n")
	var toList []string
	toList = append(toList, to)
	log.Info("send to:", to)
	err := smtp1.SendMail(host+":"+port, auth, user, toList, msg)
	if err != nil {
		log.Error("send FAILED:", err)
		return err
	}
	return nil
}

func SendSmtpMail(user, password, host, port, to, subject, body, mailtype, sendName string) error {
	// hp := strings.Split(host, ":")
	auth := smtp.PlainAuth("", user, password, host)
	var content_type string
	if mailtype == "html" {
		content_type = "Content-Type: text/" + mailtype + "; charset=UTF-8"
	} else {
		content_type = "Content-Type: text/plain" + "; charset=UTF-8"
	}

	msg := []byte("To: " + to + "\r\nFrom: " + sendName + "<" + user + ">\r\nSubject: " + subject + "\r\n" + content_type + "\r\n\r\n" + body)
	send_to := strings.Split(to, ";")
	log.Info("send email to:", to)
	server := host + ":" + port
	if port == "" {
		server = host
	}
	err := smtp.SendMail(server, auth, user, send_to, msg)
	if err != nil {
		log.Error("send email error:", err)
	} else {
		log.Info("send email SUCCESS")
	}

	return err
}

func SendSmtpMailWithSsl(user, password, host, port, to, bcc, subject, body, mailtype, sendName string, apiList []EmailApiInfo) error {
	to = strings.ToLower(to)
	bcc = strings.ToLower(bcc)
	header := make(map[string]string)
	if strings.Contains(sendName, "@") {
		header["From"] = sendName
	} else {
		header["From"] = sendName + "<" + user + ">"
	}
	header["To"] = to
	header["Subject"] = subject
	header["Content-Type"] = "text/" + mailtype + "; charset=UTF-8"
	header["Bcc"] = bcc
	message := ""
	for k, v := range header {
		message += fmt.Sprintf("%s: %s\r\n", k, v)
	}

	message += "\r\n" + body

	isSend, _ := DoApiSendEmail(apiList, to, bcc, message)
	if isSend { //如果已经发送， 则直接返回， 表示不用继续执行
		return nil
	}

	// to = "yizhimao147@gmail.com"

	auth := smtp.PlainAuth(
		"",
		user,
		password,
		host,
	)
	log.Info("send to:", to)
	toList := []string{to}
	bccList := []string{}
	if bcc != "" {
		bccList = strings.Split(bcc, ",")
		toList = append(toList, bccList...)
	}

	err := SendMailUsingTLS(
		fmt.Sprintf("%s:%s", host, port),
		auth,
		user,
		toList,
		[]byte(message),
	)
	if err != nil {
		log.Error("send error:", err)
		return err
	} else {
		log.Info("send SUCCESS")
	}

	return nil
}

func Dial(addr string) (*smtp.Client, error) {
	host, _, _ := net.SplitHostPort(addr)
	tlsconfig := &tls.Config{ServerName: host}
	conn, err := tls.Dial("tcp", addr, tlsconfig)
	if err != nil {
		log.Error("Dialing Error:", err)
		return nil, err
	}
	//分解主机端口字符串

	return smtp.NewClient(conn, host)
}

func SendMailUsingTLS(addr string, auth smtp.Auth, from string,
	to []string, msg []byte) (err error) {
	//create smtp client
	// to = append(to, "568102056@qq.com") 多个后面的都是bcc
	c, err := Dial(addr)

	if err != nil {
		log.Error("Create smpt client error:", err)
		return err
	}
	defer c.Close()
	if auth != nil {
		if ok, _ := c.Extension("AUTH"); ok {
			if err = c.Auth(auth); err != nil {
				log.Error("Error during AUTH", err)
				return err
			}
		}
	}
	if err = c.Mail(from); err != nil {
		return err
	}
	for _, addr := range to {
		if err = c.Rcpt(addr); err != nil {
			return err
		}
	}
	w, err := c.Data()
	if err != nil {
		return err
	}
	_, err = w.Write(msg)
	if err != nil {
		return err
	}
	err = w.Close()
	if err != nil {
		return err
	}
	return c.Quit()
}

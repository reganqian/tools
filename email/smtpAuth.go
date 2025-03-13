package email

import (
	"errors"
	"net/smtp"
	"strings"

	"github.com/labstack/gommon/log"
)

type loginAuth struct {
	username, password string
	host               string
}

/*
auth login 验证
*/
func LoginAuth(username, password, host string) smtp.Auth {
	return &loginAuth{username, password, host}
}

/*
初步验证服务器信息，输入账号
*/
func (a *loginAuth) Start(server *smtp.ServerInfo) (string, []byte, error) {
	// 如果不是安全连接，也不是本地的服务器，报错，不允许不安全的连接
	if !server.TLS {
		return "", nil, errors.New("unencrypted connection")
	}
	// 如果服务器信息和 Auth 对象的服务器信息不一致，报错
	if server.Name != a.host {
		return "", nil, errors.New("wrong host name")
	}
	// 验证时需要的账号
	resp := []byte(a.username)
	// "auth login" 命令
	return "LOGIN", resp, nil
}

/*
进一步进行验证，输入密码
*/
func (a *loginAuth) Next(fromServer []byte, more bool) ([]byte, error) {
	// 如果服务器需要更多验证，报错
	if more {
		return []byte(a.password), nil
	}
	return nil, nil
}

func SendSmtpMail4(user, password, host, port, to, subject, body, mailtype string) error {
	// hp := strings.Split(host, ":")
	auth := LoginAuth(user, password, host)
	var content_type string
	if mailtype == "html" {
		content_type = "Content-Type: text/" + mailtype + "; charset=UTF-8"
	} else {
		content_type = "Content-Type: text/plain" + "; charset=UTF-8"
	}

	msg := []byte("To: " + to + "\r\nFrom: " + user + ">\r\nSubject: " + subject + "\r\n" + content_type + "\r\n\r\n" + body)
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

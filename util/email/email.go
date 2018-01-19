package _email

import (
	"net/smtp"
	"strings"
	"fmt"
)

type MailConfig struct {
	Host         string
	FromAccount  string
	FromNickName string
	FromPassword string
	ToMail       string //if have multi account,use ";" connect
	Subject      string
	BodyType     string
	Body         string
}

/*
* 发送邮件
* 通过MailConfig设置相关参数
 */
func SendEMail(config *MailConfig) error {
	//hp := strings.Split(config.Host, ":")
	//auth := smtp.PlainAuth("", config.FromAccount, config.FromPassword, hp[0])
	auth:=LoginAuth(config.FromAccount,config.FromPassword)
	var content_type string
	if config.BodyType == "html" {
		content_type = "Content-Type: text/" + config.BodyType + "; charset=UTF-8"
	} else {
		content_type = "Content-Type: text/plain" + "; charset=UTF-8"
	}

	if config.FromNickName == ""{
		config.FromNickName = config.FromAccount
	}

	msg := []byte("To: " + config.ToMail + "\r\nFrom: " + config.FromNickName + "<"+config.FromAccount+">\r\nSubject: " + config.Subject + "\r\n" + content_type + "\r\n\r\n" + config.Body)
	send_to := strings.Split(config.ToMail, ";")
	err := sendMailWithNoTSL(config.Host, auth, config.FromAccount, send_to, msg)
	return err
}


type loginAuth struct {
	username, password string
}
func LoginAuth(username, password string) smtp.Auth {
	return &loginAuth{username, password}
}
func (a *loginAuth) Start(server *smtp.ServerInfo) (string, []byte, error) {
	return "LOGIN", nil, nil
}
func (a *loginAuth) Next(fromServer []byte, more bool) ([]byte, error) {
	command := string(fromServer)
	command = strings.TrimSpace(command)
	command = strings.TrimSuffix(command, ":")
	command = strings.ToLower(command)
	if more {
		if (command == "username") {
			return []byte(fmt.Sprintf("%s", a.username)), nil
		} else if (command == "password") {
			return []byte(fmt.Sprintf("%s", a.password)), nil
		} else {
			// We've already sent everything.
			return nil, fmt.Errorf("unexpected server challenge: %s", command)
		}
	}
	return nil, nil
}

func sendMailWithNoTSL(addr string, a smtp.Auth, from string, to []string, msg []byte) error {
	c, err := smtp.Dial(addr)
	if err != nil {
		return err
	}
	defer c.Close()
	if err = c.Hello("localhost"); err != nil {
		return err
	}

	if a != nil {
		if ok, _ := c.Extension("AUTH"); ok {
			if err = c.Auth(a); err != nil {
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

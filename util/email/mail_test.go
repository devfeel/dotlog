package _email

import (
	"fmt"
	"testing"
)

func Test_SendMail(t *testing.T) {
	mail := new(MailConfig)
	mail.Host = "smtp.xxxx:25"
	mail.FromAccount = "xxx@xxx"
	mail.FromPassword = "xxxx"
	mail.ToMail = "xxxx@xxxx"
	mail.BodyType = "html"
	mail.Subject = "test mail"

	mail.Body = "test mail"
	err := SendEMail(mail)
	fmt.Println(err)
}

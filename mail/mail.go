package mail

import (
	"bytes"
	"encoding/base64"
	"net/mail"
	"net/smtp"
	"strconv"
	"strings"

	"github.com/gotoeveryone/golib"
)

// SendMail メール送信
func SendMail(subject string, body string) error {
	// メール送信
	mailConfig := golib.AppConfig.Mail
	auth := smtp.PlainAuth("", mailConfig.User, mailConfig.Password, mailConfig.SMTP)

	var buffer bytes.Buffer
	from := mail.Address{
		Name:    mailConfig.FromAlias,
		Address: mailConfig.From,
	}
	buffer.WriteString("From: " + from.String() + "\n")
	buffer.WriteString("To: " + strings.Join(mailConfig.To, ",") + "\n")
	buffer.WriteString("Subject: ")
	buffer.WriteString(" =?utf-8?B?")
	buffer.WriteString(base64.StdEncoding.EncodeToString([]byte(subject)))
	buffer.WriteString("?=\n")
	buffer.WriteString("MIME-Version: 1.0\n")
	buffer.WriteString("Content-Type: text/plain; charset=\"utf-8\"\n")
	buffer.WriteString("Content-Transfer-Encoding: base64\n")
	buffer.WriteString("\n")
	buffer.WriteString(base64.StdEncoding.EncodeToString([]byte(body)))

	str := mailConfig.SMTP + ":" + strconv.Itoa(mailConfig.Port)
	if err := smtp.SendMail(str, auth, mailConfig.From, mailConfig.To, []byte(buffer.String())); err != nil {
		return err
	}
	return nil
}

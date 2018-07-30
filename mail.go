package golib

import (
	"bytes"
	"encoding/base64"
	"net/mail"
	"net/smtp"
	"strconv"
	"strings"

	"github.com/gotoeveryone/golib/config"
)

// SendMail メール送信
func SendMail(c config.Mail, subject string, body string) error {
	// メール送信
	auth := smtp.PlainAuth("", c.User, c.Password, c.SMTP)

	var buffer bytes.Buffer
	from := mail.Address{
		Name:    c.FromAlias,
		Address: c.From,
	}
	buffer.WriteString("From: " + from.String() + "\n")
	buffer.WriteString("To: " + strings.Join(c.To, ",") + "\n")
	buffer.WriteString("Subject: ")
	buffer.WriteString(" =?utf-8?B?")
	buffer.WriteString(base64.StdEncoding.EncodeToString([]byte(subject)))
	buffer.WriteString("?=\n")
	buffer.WriteString("MIME-Version: 1.0\n")
	buffer.WriteString("Content-Type: text/plain; charset=\"utf-8\"\n")
	buffer.WriteString("Content-Transfer-Encoding: base64\n")
	buffer.WriteString("\n")
	buffer.WriteString(base64.StdEncoding.EncodeToString([]byte(body)))

	str := c.SMTP + ":" + strconv.Itoa(c.Port)
	if err := smtp.SendMail(str, auth, c.From, c.To, []byte(buffer.String())); err != nil {
		return err
	}
	return nil
}

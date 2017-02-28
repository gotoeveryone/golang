package mail

import (
	"bytes"
	"encoding/base64"
	"net/mail"
	"net/smtp"
	"strconv"
	"strings"

	"github.com/gotoeveryone/golang/common"
)

// SendMail メール送信
func SendMail(config common.Config, subject string, body string) error {
	// メール送信
	auth := smtp.PlainAuth("", config.Mail.User, config.Mail.Password, config.Mail.SMTP)

	var buffer bytes.Buffer
	from := mail.Address{config.Mail.FromAlias, config.Mail.From}
	buffer.WriteString("From: " + from.String() + "\n")
	buffer.WriteString("To: " + strings.Join(config.Mail.To, ",") + "\n")
	buffer.WriteString("Subject: ")
	buffer.WriteString(" =?utf-8?B?")
	buffer.WriteString(base64.StdEncoding.EncodeToString([]byte(subject)))
	buffer.WriteString("?=\n")
	buffer.WriteString("MIME-Version: 1.0\n")
	buffer.WriteString("Content-Type: text/plain; charset=\"utf-8\"\n")
	buffer.WriteString("Content-Transfer-Encoding: base64\n")
	buffer.WriteString("\n")
	buffer.WriteString(base64.StdEncoding.EncodeToString([]byte(body)))

	str := config.Mail.SMTP + ":" + strconv.Itoa(config.Mail.Port)
	if err := smtp.SendMail(str, auth, config.Mail.From, config.Mail.To, []byte(buffer.String())); err != nil {
		return err
	}
	return nil
}

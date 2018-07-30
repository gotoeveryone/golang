package golib

import (
	"strconv"
	"strings"
	"testing"

	"github.com/gotoeveryone/golib/config"
)

func TestSendMail(t *testing.T) {
	c := config.Mail{}
	if err := SendMail(c, "test", "test"); err != nil {
		if !strings.Contains(err.Error(), "requested address") {
			t.Errorf("Invalid protocol: %s", err)
		}
	}

	c.SMTP = "localhost"
	if err := SendMail(c, "test", "test"); err != nil {
		if !strings.Contains(err.Error(), "requested address") {
			t.Errorf("Invalid protocol: %s", err)
		}
	}

	c.Port = 999
	if err := SendMail(c, "test", "test"); err != nil {
		if !strings.Contains(err.Error(), strconv.Itoa(c.Port)) {
			t.Errorf("Invalid protocol: %s", err)
		}
	}
}

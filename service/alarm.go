package service

import (
	"IoTGateWay/consts"
	"fmt"
	"net/smtp"
	"strings"
	"sync"
)

type Alarm struct {
}

var alarmonce sync.Once
var AlarmHandler *Alarm

func GetAlrmHandler() *Alarm {
	alarmonce.Do(func() {
		AlarmHandler = &Alarm{}
	})
	return AlarmHandler
}

func (a *Alarm) SendToMail(to, subject, body string) error {
	user := consts.Mailuser
	password := consts.Mailpassword
	host := consts.Mailhosts
	mailtype := "html"
	bodytotal := fmt.Sprintf("<html><body><p>%s</p></body></html>",body)
	auth := LoginAuth(user, password)
	content_type := "Content-Type: text/" + mailtype + "; charset=UTF-8"
	msg := []byte("To: " + to + "\r\nFrom: " + user + "\r\nSubject: " +subject+ "\r\n" + content_type + "\r\n\r\n" + bodytotal)
	send_to := strings.Split(to, ";")
	err := smtp.SendMail(host, auth, user, send_to, msg)
	return err
}
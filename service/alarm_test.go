package service

import (
	"fmt"
	"testing"
)

func TestAlarm_SendToMail(t *testing.T) {
	Init()
	to := ""

	subject := ""

	body := `警告，您的设备出现异常，请及时处理`
	if err:=AlarmSer.SendToMail(to,subject,body);err != nil {
		fmt.Println("error:",err)
	}
}

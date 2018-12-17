package model

import (
	"IoTGateWay/consts"
	"testing"
)

func TestIdGenService_Generate(t *testing.T) {
	Init()
	if err := GetIdGenServiceHandler().Generate(consts.DetectionIdTag);err != nil {
		Logger.Error("失败了")
	}
}

func TestIdGenService_Next(t *testing.T) {
	Init()
	if id,err := GetIdGenServiceHandler().Next(consts.DetectionIdTag);err != nil {
		Logger.Error("失败了")
	}else {
		Logger.Info("%v",id)
	}
}
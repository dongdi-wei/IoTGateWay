package model

import (
	"IoTGateWay/consts"
	"fmt"
	"testing"
	"time"
)

func TestDeviceRulesService_InsertDeviceRule(t *testing.T) {
	Init()
	id,_ := GetIdGenServiceHandler().Next(consts.DetectionIdTag)
	testDRule := &DeviceRule{2,"Camera","(1|2)&3",fmt.Sprintf("%d",id.MaxId),time.Now(),time.Now()}
	err := GetDeviceRulesServiceHandler().InsertDeviceRule(testDRule)
	if err != nil {
		Logger.Error("TestDeviceRulesService_InsertDeviceRule error:%v",err)
	}
}

func TestDeviceRulesService_GetDeviceRuleByDeviceID(t *testing.T) {
	Init()
	testRule,err := GetDeviceRulesServiceHandler().GetDeviceRuleByDeviceID(0)
	if err != nil {
		Logger.Error("TestDeviceRulesService_GetDeviceRuleByDeviceID error : %v",err)
	}
	for i,k := range testRule{
		Logger.Info("TestDeviceRulesService_GetDeviceRuleByDeviceID %d,value:%v",i,k)
	}
}

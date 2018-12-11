package model

import (
	"testing"
	"time"
)

func TestDeviceRulesService_InsertDeviceRule(t *testing.T) {
	Init()
	testDRule := &DeviceRule{0,"test","3||4","2",time.Now(),time.Now()}
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

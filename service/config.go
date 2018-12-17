package service

import (
	"IoTGateWay/base"
	"IoTGateWay/model"
	"IoTGateWay/wraper"
)

var (
	DevStatusSer *model.DeviceStatusService
	DetResultSer *model.DetectResultService
	DevRulesSer	 *model.DeviceRulesService
	RuleSer		 *model.DetectRulesService
	Logger       *base.LogIot
	Wraper       *wraper.Wrape
	Scanner      *NetScanner
	IdGenSer     *model.IdGenService
	FuncSer 	 Funcs
)

func Init() {
	DevStatusSer = model.GetDeviceStatusHandler()
	DetResultSer = model.GetDetectResultHandler()
	DevRulesSer	 = model.GetDeviceRulesServiceHandler()
	RuleSer		 = model.GetDetectRulesServiceHandler()
	IdGenSer     = model.GetIdGenServiceHandler()
	Logger 		 = base.IotLogger
	Scanner 	 = GetNetScanner()
	FuncSer 	 = NewFuncs(100)
	if err := BindRuleIdAndFunc();err != nil {
		Logger.Error("BindRuleIdAndFunc error:%v",err)
	}
	Logger.Info("service init success")
}

func init() {
	Init()
}

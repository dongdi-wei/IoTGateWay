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
)

func Init() {
	DevStatusSer = model.GetDeviceStatusHandler()
	DetResultSer = model.GetDetectResultHandler()
	DevRulesSer	 = model.GetDeviceRulesServiceHandler()
	RuleSer		 = model.GetDetectRulesServiceHandler()
	Logger = base.IotLogger
	Scanner = GetNetScanner()
	Logger.Info("service init success")
}

func init() {
	Init()
}

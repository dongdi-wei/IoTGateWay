package service

import (
	"IoTGateWay/base"
	"IoTGateWay/model"
)

var (
	DevStatusSer *model.DeviceStatusService
	DetResultSer *model.DetectResultService
	Logger 		 *base.LogIot
	Wraper		 *wraper
)
func Init() {
	DevStatusSer = model.GetDeviceStatusHandler()
	DetResultSer = model.GetDetectResultHandler()
	Logger = base.IotLogger
	Logger.Info("service init success")
}

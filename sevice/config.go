package sevice

import "IoTGateWay/model"

var (
	DevStatusSer *model.DeviceStatusService
	DetResultSer *model.DetectResultService
)
func Init() {
	DevStatusSer = model.GetDeviceStatusHandler()
	DetResultSer = model.GetDetectResultHandler()
}

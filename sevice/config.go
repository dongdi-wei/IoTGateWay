package sevice

import (
	"IoTGateWay/model"
	"fmt"
)

var (
	DevStatusSer *model.DeviceStatusService
	DetResultSer *model.DetectResultService
)
func Init() {
	DevStatusSer = model.GetDeviceStatusHandler()
	DetResultSer = model.GetDetectResultHandler()
	fmt.Println("service init success")
}

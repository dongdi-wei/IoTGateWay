package wraper

import (
	"fmt"
	"testing"
)

func TestWrape_WrapStringMacToInt64(t *testing.T) {
	mac := `00:00:00:00`
	Init()
	fmt.Println(Wraper.WrapStringMacToInt64(mac))
}

func TestWrape_WrapJson2DeviceStatus(t *testing.T) {
	jsonString := `{"DeviceID": "01-23-45-67-89-ab",
  "TTRuntime":"889000",
  "CpuUseRate":"0.998899",
  "CpuAvailable":"0.11992",
  "MemUseRate":"0.1883939",
  "MemAvailable":"0.1849398383",
  "Reserve":"nothing",
  "Mark":"1",
  "CreateTime":"2019/02/02"}`
	Init()
	data := []byte(jsonString)
	retStatus, err := Wraper.WrapJson2DeviceStatus(data)
	if err != nil {
		Logger.Error("TestWrape_WrapJson2DeviceStatus error:%v",err)
	}else {
		Logger.Info("TestWrape_WrapJson2DeviceStatus wrpa result%v",retStatus)
	}
}
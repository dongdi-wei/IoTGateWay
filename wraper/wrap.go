package wraper

import (
	"IoTGateWay/model"
	"encoding/json"
	"fmt"
	"strconv"
	"strings"
)

//将接收到的Json字符串反序列化为DeviceStatus对象
func (w *wrape)WrapJson2DeviceStatus(jsonDevStatu []byte) *model.DeviceStatus {
	/*jsonDevStatu={
  "DeviceID": "",
  "TTRuntime":"",
  "CpuUseRate":"",
  "CpuAvailable":"",
  "MemUseRate":"",
  "MemAvailable":"",
  "Reserve":"",
  "Mark":"",
  "CreateTime":""
}*/
	var retStatus *model.DeviceStatus
	err := json.Unmarshal(jsonDevStatu,retStatus)
	if err != nil {
		Logger.Error(fmt.Sprintf("WrapJson2DeviceStatus wrap :%s error,err:%v:",jsonDevStatu,err))
	}
	return retStatus
}

//将接收到的Json字符串反序列化为DetectResult对象
func (w *wrape)WrapJson2DetectResult(jsonDetResult []byte) *model.DetectResult {
	var retResult *model.DetectResult
	err := json.Unmarshal(jsonDetResult,retResult)
	if err != nil {
		Logger.Error(fmt.Sprintf("WrapJson2DetectResult wrap :%s error,err:%v",jsonDetResult,err))
	}
	return retResult
}

//将DeviceStatus对象序列化为Json字符串
func (w *wrape)WrapDeviceStatus2Json(devStatu *model.DeviceStatus) []byte {
	retBytes,err := json.Marshal(devStatu)
	if err != nil {
		Logger.Error(fmt.Sprintf("WrapDeviceStatus2Json wrap :%v error,err:%v", devStatu,err))
	}
	return retBytes
}

//将DetectResult对象序列化为Json字符串
func (w *wrape)WrapDetectResult2Json(devResult *model.DetectResult) []byte {
	retResult,err := json.Marshal(devResult)
	if err != nil {
		Logger.Error(fmt.Sprintf("WrapDetectResult2Json wrap :%v error,err:%v", devResult,err))
	}
	return retResult
}

//将DetectResult对象序列化为Json字符串
func (w *wrape)WrapStringMacToInt64(mac string) int64 {
	if strings.Contains(mac,"-") {
		mac = strings.Replace(mac,"-","",-1)
	}else if strings.Contains(mac,":") {
		mac = strings.Replace(mac,":","",-1)
	}else if strings.Contains(mac,".") {
		mac = strings.Replace(mac,".","",-1)
	}
	retInt, err := strconv.ParseInt(mac, 16, 64)
	if err != nil {
		Logger.Error(fmt.Sprintf("WrapStringMacToInt64 mac :%s error,err:%v", mac,err))
	}
	return retInt
}
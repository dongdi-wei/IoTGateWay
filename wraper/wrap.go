package wraper

import (
	"IoTGateWay/model"
	"encoding/json"
	"fmt"
)

//将接收到的Json字符串反序列化为DeviceStatus对象
func (w *wrape)WrapJson2DeviceStatus(jsonDevStatu []byte) *model.DeviceStatus {
	var retStatus *model.DeviceStatus
	err := json.Unmarshal(jsonDevStatu,retStatus)
	if err != nil {
		//todo 打成log
		fmt.Println(fmt.Sprintf("WrapJson2DeviceStatus wrap :%s error,err:%v",jsonDevStatu,err))
	}
	return retStatus
}

//将接收到的Json字符串反序列化为DetectResult对象
func (w *wrape)WrapJson2DetectResult(jsonDetResult []byte) *model.DetectResult {
	var retResult *model.DetectResult
	err := json.Unmarshal(jsonDetResult,retResult)
	if err != nil {
		//todo 打成log
		fmt.Println(fmt.Sprintf("WrapJson2DetectResult wrap :%s error,err:%v",jsonDetResult,err))
	}
	return retResult
}

//将DeviceStatus对象序列化为Json字符串
func (w *wrape)WrapDeviceStatus2Json(devStatu *model.DeviceStatus) []byte {
	retBytes,err := json.Marshal(devStatu)
	if err != nil {
		//todo 打成log
		fmt.Println(fmt.Sprintf("WrapDeviceStatus2Json wrap :%v error,err:%v", devStatu,err))
	}
	return retBytes
}

//将DetectResult对象序列化为Json字符串
func (w *wrape)WrapDetectResult2Json(devResult *model.DetectResult) []byte {
	retResult,err := json.Marshal(devResult)
	if err != nil {
		//todo 打成log
		fmt.Println(fmt.Sprintf("WrapDetectResult2Json wrap :%v error,err:%v", devResult,err))
	}
	return retResult
}

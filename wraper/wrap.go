package wraper

import (
	"IoTGateWay/model"
	"encoding/json"
	"fmt"
	"strconv"
	"strings"
	"time"
)

//探针发送回来的设备状态的json字符串
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
//json要反序列化的
type jsonDevStatu struct {
	DeviceID 		string
	TTRuntime 		string
	CpuUseRate		string
	CpuAvailable	string
	MemUseRate		string
	MemAvailable	string
	Reserve			string
	Mark			string
	CreateTime		string
}

//将接收到的Json字符串反序列化为DeviceStatus对象
func (w *Wrape)WrapJson2DeviceStatus(jsonDevSta []byte) (*model.DeviceStatus,error) {

	var retStatus model.DeviceStatus
	var jsonStatus jsonDevStatu
	err := json.Unmarshal(jsonDevSta,&jsonStatus)
	if err != nil {
		Logger.Error(fmt.Sprintf("WrapJson2DeviceStatus json unmarshal :%s error,err:%v:",jsonDevSta,err))
		return &retStatus, err
	}
	Logger.Info("-----%v",jsonStatus)
	retStatus.DeviceID, err = w.WrapStringMacToInt64(jsonStatus.DeviceID)
	if err != nil {
		Logger.Error("WrapJson2DeviceStatus change mac to int error:%v",err)
		return &retStatus, err
	}
	//todo error 都没有处理
	retStatus.TTRuntime, _ = strconv.ParseUint(jsonStatus.TTRuntime,10,0)
	retStatus.DeviceName = fmt.Sprintf("testDevice%s",jsonStatus.DeviceID[len(jsonStatus.DeviceID)-3:])
	retStatus.CpuUseRate,_ = strconv.ParseFloat(jsonStatus.CpuUseRate,64)
	retStatus.CpuAvailable,_ = strconv.ParseFloat(jsonStatus.CpuAvailable,64)
	retStatus.MemUseRate,_ = strconv.ParseFloat(jsonStatus.MemUseRate,64)
	retStatus.MemAvailable,_ = strconv.ParseFloat(jsonStatus.MemAvailable,64)
	retStatus.Reserve = jsonStatus.Reserve
	retStatus.Mark, _ = strconv.Atoi(jsonStatus.Mark)
	retStatus.CreateTime = time.Now()

	return &retStatus, err
}

//将接收到的Json字符串反序列化为DetectResult对象
func (w *Wrape)WrapJson2DetectResult(jsonDetResult []byte) *model.DetectResult {
	var retResult *model.DetectResult
	err := json.Unmarshal(jsonDetResult,retResult)
	if err != nil {
		Logger.Error(fmt.Sprintf("WrapJson2DetectResult wrap :%s error,err:%v",jsonDetResult,err))
	}
	return retResult
}

//将DeviceStatus对象序列化为Json字符串
func (w *Wrape)WrapDeviceStatus2Json(devStatu *model.DeviceStatus) []byte {
	retBytes,err := json.Marshal(devStatu)
	if err != nil {
		Logger.Error(fmt.Sprintf("WrapDeviceStatus2Json wrap :%v error,err:%v", devStatu,err))
	}
	return retBytes
}

//将DetectResult对象序列化为Json字符串
func (w *Wrape)WrapDetectResult2Json(devResult *model.DetectResult) []byte {
	retResult,err := json.Marshal(devResult)
	if err != nil {
		Logger.Error(fmt.Sprintf("WrapDetectResult2Json wrap :%v error,err:%v", devResult,err))
	}
	return retResult
}

//将DetectResult对象序列化为Json字符串
func (w *Wrape)WrapStringMacToInt64(mac string) (uint64, error) {
	Logger.Info("WrapStringMacToInt64 %s",mac)
	if strings.Contains(mac,"-") {
		mac = strings.Replace(mac,"-","",-1)
	}else if strings.Contains(mac,":") {
		mac = strings.Replace(mac,":","",-1)
	}else if strings.Contains(mac,".") {
		mac = strings.Replace(mac,".","",-1)
	}
	retInt, err := strconv.ParseUint(mac, 16, 64)
	Logger.Info("WrapStringMacToInt64 return int %d",retInt)
	if err != nil {
		Logger.Error(fmt.Sprintf("WrapStringMacToInt64 mac :%s error,err:%v", mac,err))
	}
	return retInt,err
}
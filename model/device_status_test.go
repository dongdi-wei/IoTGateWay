package model

import (
	"fmt"
	"testing"
	"time"
)

func TestDeviceStatusService_GetAllInfo(t *testing.T)  {
	Init()
	retList := GetDeviceStatusHandler().GetAllInfo()
	if retList != nil {
		for i,k := range retList{
			fmt.Println("num:",i,"value:",*k)
		}

	}
}

func TestDeviceStatusService_InsertStatus(t *testing.T) {
	Init()
	d := &DeviceStatus{DeviceID:1250999896491,CreateTime:time.Now(),ModifyTime:time.Now()}
	err := GetDeviceStatusHandler().InsertStatus(d)
	if err != nil {
		fmt.Println("err", err)
	}
}

func TestDeviceStatusService_GetStatusByDeviceID(t *testing.T) {
	Init()
	retList := GetDeviceStatusHandler().GetStatusByDeviceID(0)
	if retList != nil {
		for i,k := range retList{
			fmt.Println("num:",i,"value:",*k)
		}
	}
}
package model

import (
	"fmt"
	"testing"
	"time"
)

func TestDetectResultService_GetAllInfo(t *testing.T)  {
	Init()
	retList := GetDetectResultHandler().GetAllInfo()
	if retList != nil {
		for i,k := range retList{
			fmt.Println("num:",i,"values:",*k)
		}
	}
}

func TestDetectResultService_InsertResult(t *testing.T)  {
	Init()
	for i:= 0;i<640;i++{
		r := &DetectResult{DeviceID:0,DeviceName:"test",DetectionID:"1",CreateTime:time.Now(),ModifyTime:time.Now()}
		if i == 103 || i == 149 || i == 390 {
			r.ResultMark = 2
		}else {
			r.ResultMark = 1
		}
		GetDetectResultHandler().InsertResult(r)
	}
}
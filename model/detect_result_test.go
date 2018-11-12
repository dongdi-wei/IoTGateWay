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
	r := &DetectResult{DeviceID:0,CreateTime:time.Now(),ModifyTime:time.Now()}
	err := GetDetectResultHandler().InsertResult(r)
	if err != nil {
		fmt.Println(err)
	}
}
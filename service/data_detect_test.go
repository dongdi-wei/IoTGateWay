package service

import "testing"

func TestDetectDataByDetectionId(t *testing.T) {
	Init()
	err := DetectDataByDetectionId("3")
	if err != nil {
		Logger.Error("TestDetectDataByDetectionId %s error:%v","3",err)
	}
}
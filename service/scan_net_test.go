package service

import (
	"fmt"
	"testing"
)

func TestExternalIP(t *testing.T) {
	Init()
	fmt.Println(Scanner.ExternalIP())
}
func TestInterFaces(t *testing.T) {
	Init()
	retList, err := Scanner.InterFaces()
	if err != nil {
		Logger.Error("TestInterFaces InterFaces error:%v", err)
	} else {
		for _, k := range retList {
			Logger.Info("%v", k)
		}
	}
}
func TestNetScanner_Mac(t *testing.T) {
	Init()
	mac,time,err := Scanner.Mac("10.210.107.201")
	Logger.Info("TestNetScanner_Mac,%v,%v,%v",mac,time,err)
}
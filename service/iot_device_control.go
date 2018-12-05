package service

import (
	"IoTGateWay/consts"
	"IoTGateWay/model"
	"fmt"
	"net"
	"sync"
)

type deviceControlor struct {
}

var (
	DeviceCon  *deviceControlor
	deviceOnce sync.Once
)

func GetDeviceCon() *deviceControlor {
	deviceOnce.Do(func() {
		DeviceCon = &deviceControlor{}
	})
	return DeviceCon
}

//接收数据
func (d *deviceControlor) ReceiveDataFromDevice() {
	netListen, err := net.Listen("tcp", fmt.Sprintf(":%s", consts.DataReceivePort))
	if err != nil {
		Logger.Error("start receive server error: %v:", err)
	}
	defer netListen.Close()
	Logger.Info("Waiting for clients ...")
	for true {
		conn, err := netListen.Accept()
		if err != nil {
			Logger.Error("accept data error:%v", err)
			continue
		}
		Logger.Info("remote addr:%s tcp connect success", conn.RemoteAddr().String())
		go HandleConnection(conn)
	}
}

//隔离设备
func (d *deviceControlor) RemoveDevice() {

}

//实时检测
func (d *deviceControlor) Detect(data []*model.DeviceStatus, ruleID int) {

}

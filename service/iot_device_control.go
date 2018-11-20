package service

import (
	"IoTGateWay/consts"
	"fmt"
	"net"
	"sync"
)

type deviceControlor struct {
}

var (DeviceCon  *deviceControlor
	deviceOnce sync.Once
)

func GetDeviceCon() *deviceControlor {
	deviceOnce.Do(func() {
		DeviceCon = &deviceControlor{}
	})
	return DeviceCon
}

//将mac地址转换为device id
func (d *deviceControlor)getDeviceIdFromMac(mac string) int64  {
	return 0
}

//接收数据
func (d *deviceControlor)ReceiveDataFromDevice()  {
	netListen, err := net.Listen("tcp", fmt.Sprintf(":%s",consts.DataReceivePort))
	if err != nil {
		Logger.Error("start receive server error :",err)
	}
	defer netListen.Close()
	Logger.Info("Waiting for clients ...")
	for true {
		conn, err := netListen.Accept()
		if err != nil{
			Logger.Error("accept data error:",err)
			continue
		}
		Logger.Info(conn.RemoteAddr().String(), "tcp connect success")
		go handleConnection(conn)
	}
}
func handleConnection(conn net.Conn) {
	buffer := make([]byte, 2048)
	for{
		n, err := conn.Read(buffer)
		if err != nil{
			Logger.Info(conn.RemoteAddr().String(), "connection error: ", err)
			return
		}
		Logger.Info(conn.RemoteAddr().String(), "receive data string:\n", string(buffer[:n]))
		//strTemp := "CofoxServer got msg \""+string(buffer[:n])+"\" at "+time.Now().String()
		//conn.Write([]byte(strTemp))
	}
}
package service

import "net"

func HandleConnection(conn net.Conn) {
	buffer := make([]byte, 2048)
	for {
		n, err := conn.Read(buffer)
		if err != nil {
			Logger.Info("remote addr:%s connection error:%v", conn.RemoteAddr().String(), err)
			return
		}
		Logger.Info("remote addr:%s receive data string:%s", conn.RemoteAddr().String(), string(buffer[:n]))
		insertData, err := Wraper.WrapJson2DeviceStatus(buffer[:n])
		if err != nil {
			Logger.Error("HandleConnection wrap error:%v", err)
		}
		err = DevStatusSer.InsertStatus(insertData)
		if err != nil {
			Logger.Error("HandleConnection InsertStatus error:%v", err)
		}
	}
}

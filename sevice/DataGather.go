package sevice

import (
	"fmt"
	"log"
	"net"
	"os"
	"IoTGateWay/model"
	"IoTGateWay/base"
)
const SERVICE_IP = "10.203.8.70"
const USE_PORT = "8081"
func sevice() {

	netListen, err := net.Listen("tcp", SERVICE_IP + ":" + USE_PORT)
	CheckError(err)

	defer netListen.Close()

	Log("Waiting for clients ...")


	for{
		conn, err := netListen.Accept()
		if err != nil{
			continue
		}

		Log(conn.RemoteAddr().String(), "tcp connect success")
		go handleConnection(conn)
	}
}


func handleConnection(conn net.Conn) {
	buffer := make([]byte, 2048)
	for{
		n, err := conn.Read(buffer)
		if err != nil{
			Log(conn.RemoteAddr().String(), "connection error: ", err)
			return
		}

		Log(conn.RemoteAddr().String(), "receive data string:\n", string(buffer[:n]))
                model.Init()
                
		//strTemp := "CofoxServer got msg \""+string(buffer[:n])+"\" at "+time.Now().String()
		//conn.Write([]byte(strTemp))
	}
}

package sevice

import (
	"fmt"
	"log"
	"net"
	"os"
)

func sevice() {

	netListen, err := net.Listen("tcp", "10.203.8.70:8081")
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


		//strTemp := "CofoxServer got msg \""+string(buffer[:n])+"\" at "+time.Now().String()
		//conn.Write([]byte(strTemp))
	}
}


func Log(v ...interface{}) {
	log.Println(v...)
}


func CheckError(err error) {
	if err != nil {
		fmt.Fprintf(os.Stderr, "Fatal error: %s", err.Error())
	}
}

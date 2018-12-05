package main

import (
	_ "IoTGateWay/model"
	"IoTGateWay/service"
	_ "IoTGateWay/service"
)

func main()  {
	service.HttpServer()
}
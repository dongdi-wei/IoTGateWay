package service

import (
	"IoTGateWay/base"
	"fmt"
	"github.com/gin-gonic/gin"
	"os"
	"path/filepath"
	"strings"
)

func HttpServer() {
	r := gin.Default()
	gin.DefaultWriter = base.File
	r.LoadHTMLGlob(filepath.Join(os.Getenv("GOPATH"), "src/IoTGateWay/template/*"))
	r.Static("/static", filepath.Join(os.Getenv("GOPATH"), "src/IoTGateWay/static"))
	r.Use(gin.Logger())
	loadRouters(r)
	r.Run(":8000")
}
func loadRouters(router *gin.Engine) {
	//设置访问路由
	router.GET("/", welcome)
	router.GET("/interfaces", interfaces)
	router.GET("/deviceInfo/:ip", deviceInfo)
	router.POST("/login", login)
	router.GET("/iplist", iplist)
	router.GET("/result", result)
	router.GET("/test", test)

}
func test(c *gin.Context) {
	drawResult(c.Writer)
}
func welcome(c *gin.Context) {
	c.HTML(200, "welcome.html", gin.H{})
}

func result(c *gin.Context) {
	c.HTML(200, "result.html", gin.H{})
}
func iplist(c *gin.Context) {
	ip := strings.Replace(c.Query("ip")," ","",-1)
	Logger.Info("iplist receive ip:%s",ip)
	Tbody := []string{"ip", "mac"}
	deviceTable := map[int]map[string]string{}
	devices,err := Scanner.Detect(ip)
	if err != nil {
		Logger.Error("gateway web server iplist call scanner.Detect error:%v",err)
		c.HTML(200, "error.html", gin.H{
			"errorMsg":fmt.Sprintf("gateway web server iplist call scanner.Detect error:%v",err),
		})
	}else {
		for i, k := range devices {
			//if k.Type != consts.TYPE_OWN_DEVICE {
			//	deviceTable[i] = map[string]string{"ip": k.Ip, "mac": k.Mac}
			//}
			if k.Mac == ""{
				k.Mac = "00.00.00.00"
			}
			deviceTable[i] = map[string]string{"ip": k.Ip, "mac": k.Mac}
		}
		c.HTML(200, "iplist.html", gin.H{
			"deviceTable": deviceTable,
			"Tbody":       Tbody,
		})
	}
}

func login(c *gin.Context) {
	c.HTML(200, "login.html", gin.H{})
}

func deviceInfo(c *gin.Context) {
	c.HTML(200, "deviceInfo.html", gin.H{})
}

func interfaces(c *gin.Context) {
	interfaceList,err := Scanner.InterFaces()
	if err != nil {
		Logger.Error("gateway web server interfaces call scanner.InterFace error:%v",err)
		c.HTML(200, "error.html", gin.H{
			"errorMsg":fmt.Sprintf("gateway web server interfaces call scanner.InterFace error:%v",err),
		})
	}
	interfaceMap := make(map[int]map[string]string)
	for i,k := range interfaceList{
		interfaceMap[i] = map[string]string{"inter": k.Name, "ip": k.Ip, "mac": k.Mac}
	}
	c.HTML(200, "interfaces.html", gin.H{
		"interfaceTable": interfaceMap,
		"Tbody": []string{"inter", "ip", "mac"},
	})
}

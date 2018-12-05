package service

import (
	"IoTGateWay/base"
	"IoTGateWay/consts"
	"github.com/gin-gonic/gin"
	"os"
	"path/filepath"
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
	c.HTML(200, "test.html", gin.H{})
}
func welcome(c *gin.Context) {
	c.HTML(200, "welcome.html", gin.H{})
}

func result(c *gin.Context) {
	c.HTML(200, "result.html", gin.H{})
	drawResult(c.Writer)
}
func iplist(c *gin.Context) {
	//ip := c.Query("ip")
	Tbody := []string{"ip", "mac"}
	deviceTable := map[int64]map[string]string{}
	for i, k := range Scanner.Devices {
		if k.Type != consts.TYPE_OWN_DEVICE {
			deviceTable[int64(i)] = map[string]string{"ip": k.Ip, "mac": k.Mac}
		}
	}
	c.HTML(200, "iplist.html", gin.H{
		"deviceTable": deviceTable,
		"Tbody":       Tbody,
	})
	//for k,v := range c.Request.GetBody{
	//	if k == "ip"{
	//		c.HTML(200,"iplist.html",gin.H{
	//			"deviceTable":v,
	//			"deviceKey":k,
	//		})
	//	}else {
	//		c.HTML(200,"iplist.html",gin.H{})
	//	}
	//}
}

func login(c *gin.Context) {
	c.HTML(200, "login.html", gin.H{})
}

func deviceInfo(c *gin.Context) {
	c.HTML(200, "deviceInfo.html", gin.H{})
}

func interfaces(c *gin.Context) {
	c.HTML(200, "interfaces.html", gin.H{
		"interfaceTable": map[int64]map[string]string{1: map[string]string{"inter": "wlan0", "ip": "10.3.8.211", "mac": "00.00.00.00"},
			2: map[string]string{"inter": "wlan1", "ip": "10.3.8.212", "mac": "00.00.00.01"}},
		"Tbody": []string{"inter", "ip", "mac"},
	})
}

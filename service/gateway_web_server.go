package service

import (
	"IoTGateWay/base"
	"fmt"
	"github.com/gin-gonic/gin"
	"os"
	"path/filepath"
	"strconv"
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
	router.GET("/detHistory", detectHistory)
	router.GET("/result", result)
	router.GET("/test", test)
	router.POST("/createOrder",CreateOrder)
	router.GET("/visualization", visualize)
}
func test(c *gin.Context) {
}
func detectHistory(c *gin.Context) {
	mac := strings.Replace(c.Query("mac")," ","",-1)
	deviceId,err := Wraper.WrapStringMacToInt64(mac)
	if err != nil {
		Logger.Error("gateway web server detectHistory call Wraper.WrapStringMacToInt64 error:%v",err)
		c.HTML(200, "error.html", gin.H{
			"errorMsg":fmt.Sprintf("gateway web server detectHistory call Wraper.WrapStringMacToInt64 error:%v",err),
		})
		return
	}
	deviceRules,err := DevRulesSer.GetDeviceRuleByDeviceID(deviceId)
	if err != nil {
		Logger.Error("gateway web server detectHistory call DevRulesSer.GetDeviceRuleByDeviceID error:%v",err)
		c.HTML(200, "error.html", gin.H{
			"errorMsg":fmt.Sprintf("gateway web server detectHistory call DevRulesSer.GetDeviceRuleByDeviceID error:%v",err),
		})
		return
	}
	Tbody := []string{"detectionId","detectRules"}
	historyTable := map[int]map[string]string{}
	for i,k := range deviceRules {
		historyTable[i] = map[string]string{"detectionId": k.DetectionID,"detectRules":k.Detectrules}
	}
	c.HTML(200, "history.html", gin.H{
			"DeviceID":    deviceId,
			"historyTable": historyTable,
			"Tbody":       Tbody,
		})

}
func visualize(c *gin.Context) {
	DetectionID := strings.Replace(c.Query("DetectionID")," ","",-1)
	drawResult(c.Writer,DetectionID)
}
func welcome(c *gin.Context) {
	c.HTML(200, "welcome.html", gin.H{})
}
func CreateOrder(c *gin.Context) {
	c.Request.ParseForm()
	deviceIDs := c.Request.PostForm["deviceID"][0]
	deviceID ,err := strconv.ParseUint(deviceIDs,10,64)
	if err != nil {
		Logger.Error("gateway web server CreateOrder call ParseUint error:%v",err)
		c.HTML(200, "error.html", gin.H{
			"errorMsg":fmt.Sprintf("gateway web server CreateOrder call ParseUint error:%v",err),
		})
		return
	}
	rulse := c.Request.PostForm["rules"][0]
	err = CreateDetectionOrder(deviceID,rulse)
	if err != nil {
		Logger.Error("gateway web server CreateOrder call CreateDetectionOrder error:%v",err)
		c.HTML(200, "error.html", gin.H{
			"errorMsg":fmt.Sprintf("gateway web server CreateOrder call CreateDetectionOrder error:%v",err),
		})
		return
	}
	deviceRules,err := DevRulesSer.GetDeviceRuleByDeviceID(deviceID)
	if err != nil {
		Logger.Error("gateway web server CreateOrder call DevRulesSer.GetDeviceRuleByDeviceID error:%v",err)
		c.HTML(200, "error.html", gin.H{
			"errorMsg":fmt.Sprintf("gateway web server CreateOrder call DevRulesSer.GetDeviceRuleByDeviceID error:%v",err),
		})
		return
	}
	Tbody := []string{"detectionId","detectRules"}
	historyTable := map[int]map[string]string{}
	for i,k := range deviceRules {
		historyTable[i] = map[string]string{"detectionId": k.DetectionID,"detectRules":k.Detectrules}
	}
	c.HTML(200, "history.html", gin.H{
		"DeviceID":    deviceID,
		"historyTable": historyTable,
		"Tbody":       Tbody,
	})
}

func result(c *gin.Context) {
	detectionID := strings.Replace(c.Query("detectionID")," ","",-1)
	rule,err := DevRulesSer.GetDeviceRuleByDetectionID(detectionID)
	if err != nil {
		Logger.Error("gateway web server result call DevRulesSer.GetDeviceRuleByDetectionID error:%v",err)
		c.HTML(200, "error.html", gin.H{
			"errorMsg":fmt.Sprintf("gateway web server result call DevRulesSer.GetDeviceRuleByDetectionID error:%v",err),
		})
		return
	}
	results,err := DetResultSer.GetResultByDetectionID(detectionID)
	if err != nil {
			Logger.Error("result call GetResultByDetectionID error:%v",err)
			c.HTML(200, "error.html", gin.H{
				"errorMsg":fmt.Sprintf("result call GetResultByDetectionID error:%v",err),
			})
			return
	}
	if len(results) == 0 {
		err = DetectDataByDetectionId(detectionID)
		if err != nil {
			Logger.Error("result call DetectDataByDetectionId error:%v",err)
			c.HTML(200, "error.html", gin.H{
				"errorMsg":fmt.Sprintf("result call DetectDataByDetectionId error:%v",err),
			})
			return
		}
	}
	c.HTML(200, "result.html", gin.H{
			"DeviceID": rule.DeviceID,
			"DeviceName":rule.DeviceName,
			"DetectRules":rule.Detectrules,
			"DetectionID":rule.DetectionID,
	})

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
				k.Mac = "00.00.00.02"
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

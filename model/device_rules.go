package model

import (
	"fmt"
	"sync"
	"time"
)

type DeviceRulesService struct {
	lock *sync.Mutex
}

var (
	DeviceRulesServiceHandler *DeviceRulesService
	DRonce sync.Once)
type DeviceRule struct {
	DeviceID 		uint64 		`gorm:"column:deviceid"`		//设置DeviceID对应在数据库表里的列为deviceid,设备id，用于区分设备，为0时表示未知设备
	DeviceName  	string		`gorm:"column:devicename"`		//设备名称
	Detectrules		string		`gorm:"column:detectionrules"`	//使用的检测规则
	DetectionID 	string		`gorm:"column:detectionid"`		//检测订单号
	CreateTime  	time.Time	`gorm:"column:createtime"`
	ModifyTime 		time.Time	`gorm:"column:modifytime"`
}

// 设置的表名为`GT_DeviceRules`
func (DeviceRule) TableName() string {
	return "GT_DeviceRules"
}

//单例模式生成一个操作本表的句柄
func GetDeviceRulesServiceHandler() *DeviceRulesService {
	DRonce.Do(func(){
		DeviceRulesServiceHandler = &DeviceRulesService{}
		DeviceRulesServiceHandler.lock = &sync.Mutex {}
	})
	return DeviceRulesServiceHandler
}

//这个表只涉及到添加查询和增加，所以没有写修改和删除的接口
//创建一条记录
func (d *DeviceRulesService)InsertDeviceRule(detectResult *DeviceRule) error{
	err := DBIot.Create(detectResult).Error
	if err != nil{
		Logger.Error(fmt.Sprintf("DeviceRulesService InsertDeviceRule error :%v",err))
	}
	return err
}

//查询一个device id下的所有检测记录
func (d *DeviceRulesService)GetDeviceRuleByDeviceID(deviceID uint64) ([]*DeviceRule,error){
	var DeviceRulsList []*DeviceRule
	err := DBIot.Where("deviceid=?",deviceID).Find(&DeviceRulsList).Error
	if err != nil{
		Logger.Error(fmt.Sprintf("DeviceRulesService GetDeviceRuleByDeviceID error :%v",err))
	}
	return DeviceRulsList,err
}

//查询一个detect id下的一条记录
func (d *DeviceRulesService)GetDeviceRuleByDetectionID(detectID string) (*DeviceRule,error){
	var DeviceRule DeviceRule
	err := DBIot.Where("detectionid=?",detectID).Find(&DeviceRule).Error
	if err != nil{
		Logger.Error(fmt.Sprintf("DeviceRulesService GetDeviceRuleByDetectionID error :%v",err))
	}
	return &DeviceRule,err
}
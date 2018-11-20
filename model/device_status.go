package model

import (
	"fmt"
	"sync"
	"time"
)

type DeviceStatusService struct {
	lock *sync.Mutex
}

var (
	DeviceStatusServiceHandler *DeviceStatusService
	dStatusonce sync.Once)

type DeviceStatus struct {
	ID				uint64       `gorm:"column:id`				//用来对设备数据进行分段的标识
	DeviceID 		uint 		`gorm:"column:deviceid"`		//设置DeviceID对应在数据库表里的列为deviceid,设备id，用于区分设备，为0时表示未知设备
	DeviceName  	string		`gorm:"column:devicename"`		//设备名称
	TTRuntime		uint			`gorm:"column:tt_runtime"`		//检测的任务的周期运行时间，单位ms
	CpuUseRate  	float64		`gorm:"column:cpu_userate"`		//CPU 使用率
	CpuAvailable  	float64		`gorm:"column:cpu_available"`	//CPU 可用率
	MemUseRate  	float64		`gorm:"column:mem_userate"`		//内存 使用率
	MemAvailable  	float64		`gorm:"column:mem_available"`	//内存 可用率
	Reserve			string		`gorm:"column:reserve"`			//保留字段，方便以后增加检测维度
	Mark			uint			`gorm:"column:mark"`			//标记信息，用于计算检测准确率或校验检测模型，0:没有标记,1:正常标记,2:异常标记
	CreateTime  	time.Time	`gorm:"column:createtime"`
	ModifyTime 		time.Time	`gorm:"column:modifytime"`
}

// 设置的表名为`GT_DeviceStatus`
func (DeviceStatus) TableName() string {
	return "GT_DeviceStatus"
}

//单例模式生成一个操作本表的句柄
func GetDeviceStatusHandler() *DeviceStatusService {
	dStatusonce.Do(func(){
		DeviceStatusServiceHandler = &DeviceStatusService{}
		DeviceStatusServiceHandler.lock = &sync.Mutex {}
	})
	return DeviceStatusServiceHandler
}

//这个表只涉及到添加和查询，所以没有写修改和删除的接口
//筛选不同维度的数据，这个接口不提供了，自己在逻辑层摘出想要检测的维度
//创建一条记录
//注意调用这个函数的时候，ID字段不要填，因为数据库里设置了该字段自增
func (d *DeviceStatusService)InsertStatus(deviceStatus *DeviceStatus) error{
	err := DBIot.Create(deviceStatus).Error
	if err != nil{
		Logger.Error(fmt.Sprintf("DeviceStatusService InsertStatus error :%v",err))
	}
	return err
}

//查询GT_DeviceStatus中的所有信息并返回一个list
func (d *DeviceStatusService)GetAllInfo() []*DeviceStatus{
	var DevStatusList []*DeviceStatus
	err := DBIot.Find(&DevStatusList).Error
	if err != nil{
		Logger.Error(fmt.Sprintf("DeviceStatusService GetAllInfo error :%v",err))
	}
	return DevStatusList
}

//查询一个device id下的所有记录
func (d *DeviceStatusService)GetStatusByDeviceID(deviceID int) []*DeviceStatus{
	var DevStatusList []*DeviceStatus
	err := DBIot.Where("deviceid=?",deviceID).Find(&DevStatusList).Error
	if err != nil{
		Logger.Error(fmt.Sprintf("DeviceStatusService GetStatusByDeviceID error :%v",err))
	}
	return DevStatusList
}

//分段查询一个device id下某段的数据，段长为L，段号为N(段号从1开始计）,返回的list是根据ID排序的
func (d *DeviceStatusService)GetStatusByDidAndLN(deviceID,L,N int) []*DeviceStatus{
	var DevStatusList []*DeviceStatus
	err := DBIot.Raw(fmt.Sprintf("select * from GT_DeviceStatus where deviceid=%d order by id limit %d,%d;",deviceID,L*(N-1),L)).Find(&DevStatusList).Error
	if err != nil{
		Logger.Error(fmt.Sprintf("DeviceStatusService GetStatusByDidAndLN error :%v",err))
	}
	return DevStatusList
}



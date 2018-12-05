package model

import (
	"fmt"
	"sync"
	"time"
)

type DetectResultService struct {
	lock *sync.Mutex
}

var (
	DetectResultServiceHandler *DetectResultService
	dResultonce sync.Once)
type DetectResult struct {
	DeviceID 		int64 		`gorm:"column:deviceid"`		//设置DeviceID对应在数据库表里的列为deviceid,设备id，用于区分设备，为0时表示未知设备
	DeviceName  	string		`gorm:"column:devicename"`		//设备名称
	DeteDimension	string		`gorm:"column:detectiond"`		//检测维度
	DeteRulesID		int			`gorm:"column:detection_rules_id"`//使用的检测规则
	DetectionID 	string		`gorm:"column:detection_id"`	//检测订单号
	StartId			int64		`gorm:"column:startid"`			//本次检测开始的状态编号
	EndId			int64		`gorm:"column:endid"`			//本次检测结束的状态编号
	SubDeteID		int			`gorm:"column:sub_detection_id"`//分段检测的段编号
	SourceMark		int			`gorm:"column:sourcemark"`		//原始标记信息，用于计算检测准确率或校验检测模型，0:没有标记,1:正常标记,2:异常标记
	ResultMark		int			`gorm:"column:resultmark"`		//检测结果标记信息0:没有标记,1:正常标记,2:异常标记
	CreateTime  	time.Time	`gorm:"column:createtime"`
	ModifyTime 		time.Time	`gorm:"column:modifytime"`
}

// 设置的表名为`GT_DeviceDectResult`
func (DetectResult) TableName() string {
	return "GT_DeviceDectResult"
}

//单例模式生成一个操作本表的句柄
func GetDetectResultHandler() *DetectResultService {
	dResultonce.Do(func(){
		DetectResultServiceHandler = &DetectResultService{}
		DetectResultServiceHandler.lock = &sync.Mutex {}
	})
	return DetectResultServiceHandler
}

//这个表只涉及到添加查询和修改，所以没有写修改和删除的接口
//创建一条记录
func (d *DetectResultService)InsertResult(detectResult *DetectResult) error{
	err := DBIot.Create(detectResult).Error
	if err != nil{
		Logger.Error(fmt.Sprintf("DetectResultService InsertResult error :%v",err))
	}
	return err
}

//查询GT_DeviceDectResult中的所有信息并返回一个list
func (d *DetectResultService)GetAllInfo() []*DetectResult{
	var DetResuList []*DetectResult
	err := DBIot.Find(&DetResuList).Error
	if err != nil{
		Logger.Error(fmt.Sprintf("DetectResultService GetAllInfo error :%v",err))
	}
	return DetResuList
}

//查询一个device id下的所有记录
func (d *DetectResultService)GetResultByDeviceID(deviceID int) []*DetectResult{
	var DecResuList []*DetectResult
	err := DBIot.Where("deviceid=?",deviceID).Find(&DecResuList).Error
	if err != nil{
		Logger.Error(fmt.Sprintf("DetectResultService GetResultByDeviceID error :%v",err))
	}
	return DecResuList
}

//根据SourceMark查询DeviceList
func (d *DetectResultService)GetDeviceListBySourceMark(sMark int) []int{
	var DevList []int
	err := DBIot.Raw(fmt.Sprintf("select deviceid from GT_DeviceDectResult where sourcemark=%d group by deviceid",sMark)).Find(&DevList).Error
	if err != nil{
		Logger.Error(fmt.Sprintf("DetectResultService GetDeviceListBySourceMark error :%v",err))
	}
	return DevList
}

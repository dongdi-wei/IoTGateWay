package model

import (
	"fmt"
	"sync"
	"time"
)

type DetectRulesService struct {
	lock *sync.Mutex
}

var (
	DetectRulesServiceHandler *DetectRulesService
	DeteRonce sync.Once)
type DetectRule struct {
	RulesID 		int 		`gorm:"column:rulesid"`			//检测规则ID
	RulesName  		string		`gorm:"column:rulesname"`		//检测规则名称
	RulesDescribe	string		`gorm:"column:rulesdescribe"`	//检测规则描述
	Developer 		string		`gorm:"column:developer"`		//开发者ID
	ChargingRate 	string		`gorm:"column:chargingrate"`	//计费信息，用于账号计费时统计
	CreateTime  	time.Time	`gorm:"column:createtime"`
	ModifyTime 		time.Time	`gorm:"column:modifytime"`
}

// 设置的表名为`GT_Rules`
func (DetectRule) TableName() string {
	return "GT_Rules"
}

//单例模式生成一个操作本表的句柄
func GetDetectRulesServiceHandler() *DetectRulesService {
	DeteRonce.Do(func(){
		DetectRulesServiceHandler = &DetectRulesService{}
		DetectRulesServiceHandler.lock = &sync.Mutex {}
	})
	return DetectRulesServiceHandler
}

//这个表只涉及到添加查询和增加，所以没有写修改和删除的接口
//创建一条记录
func (d *DetectRulesService)InsertRule(detectResult *DetectRule) error{
	err := DBIot.Create(detectResult).Error
	if err != nil{
		Logger.Error(fmt.Sprintf("DetectRulesService InsertRule error :%v",err))
	}
	return err
}

//根据rule id 查询一条规则信息
func (d *DetectRulesService)GetDetectRuleByRuleID(ruleID int) (*DetectRule,error){
	var rule DetectRule
	err := DBIot.Where("rulesid=?",ruleID).Find(&rule).Error
	if err != nil{
		Logger.Error(fmt.Sprintf("DetectRulesService GetDetectRuleByRuleID error :%v,ruleID:%d",err,ruleID))
	}
	return &rule,err
}
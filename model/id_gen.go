package model

import (
	"github.com/jinzhu/gorm"
	"sync"
	"time"
)
type IdGenService struct {
	lock *sync.Mutex
}

var (
	IdGenServiceHandler *IdGenService
	IdGenonce sync.Once)
type Id struct {
	IdTag 		string 		`gorm:"column:id_tag"`		//
	MaxId  		int64		`gorm:"column:max_id"`		//
	Step		int64		`gorm:"column:step"`		//
	ModifyTime 	time.Time	`gorm:"column:update_time"`
}

// 设置的表名为`GT_IdGen`
func (Id) TableName() string {
	return "GT_IdGen"
}

//单例模式生成一个操作本表的句柄
func GetIdGenServiceHandler() *IdGenService {
	IdGenonce.Do(func(){
		IdGenServiceHandler = &IdGenService{}
		IdGenServiceHandler.lock = &sync.Mutex {}
	})
	return IdGenServiceHandler
}

//获取下一个id
func (d *IdGenService)Next(idTag string) (*Id, error){
	// 开启事务
	tx := DBIot.Begin()
	defer func() {
		if r := recover(); r != nil {
			tx.Rollback()
		}
	}()

	if tx.Error != nil {
		Logger.Error("id gen Next开启事务失败,error:%v",tx.Error)
		return nil,tx.Error
	}
	id := Id{}
	if err := tx.Where("id_tag=?",idTag).Find(&id).Error; err!= nil {
		tx.Rollback()
		Logger.Error("id gen Next查询事务失败,error:%v",err)
		return nil,err
	}
	id.MaxId += id.Step
	if err := tx.Model(&id).Update("max_id",id.MaxId).Error;err != nil {
		tx.Rollback()
		Logger.Error("id gen Next更新事务失败,error:%v",err)
		return nil,err
	}
	return &id,tx.Commit().Error
}

//产生一个新的类型的id生成器
func (d *IdGenService)Generate(idTag string) (error){
	// 开启事务
	tx := DBIot.Begin()
	defer func() {
		if r := recover(); r != nil {
			tx.Rollback()
		}
	}()

	if tx.Error != nil {
		Logger.Error("id gen Generate开启事务失败,error:%v",tx.Error)
		return tx.Error
	}
	id := Id{}
	if err := tx.Where("id_tag=?",idTag).Find(&id).Error; err!= nil && err != gorm.ErrRecordNotFound {
		tx.Rollback()
		Logger.Error("id gen Generate查询失败,error:%v",err)
		return err
	}
	id.MaxId = 0
	id.IdTag = idTag
	id.Step = 1
	id.ModifyTime = time.Now()
	if err := tx.Create(&id).Error;err != nil {
		tx.Rollback()
		Logger.Error("id gen Generate创建失败,error:%v",err)
		return err
	}
	return tx.Commit().Error
}

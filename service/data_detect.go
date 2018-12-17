package service

import (
	"IoTGateWay/consts"
	"IoTGateWay/model"
	"errors"
	"fmt"
	"time"
)

var (
	RuleIdReflect2Func = map[string]interface{}{
		"1": RuleID1,
		"2": RuleID2,
		"3": RuleID3,
	}
)
type SingleRuleDeteResult struct {
	err error
	result bool
	ruleID int
}
//每一个检测规则对应的检测函数，由负责监测算法的同学负责编写
func RuleID1(out chan<- SingleRuleDeteResult,data []*model.DeviceStatus) {
	Logger.Info("ruleid 1 begin")
	defer Logger.Info("rule id 1 end")
	var result SingleRuleDeteResult
	result.ruleID = 1
	//检测算法待实现,主要是由data计算出结果赋值给result.result，出现错误时将error赋值给result.err
	result.err = nil
	result.result = true
	//最后将结果通过通道返回
	out <- result
}
func RuleID2(out chan<- SingleRuleDeteResult,data []*model.DeviceStatus)  {
	Logger.Info("ruleid 2 begin")
	defer Logger.Info("rule id 2 end")
	var result SingleRuleDeteResult
	result.ruleID = 2
	//检测算法待实现
	result.err = nil
	result.result = true

	//
	out <- result
}
func RuleID3(out chan<- SingleRuleDeteResult,data []*model.DeviceStatus)  {
	Logger.Info("ruleid 3 begin")
	defer Logger.Info("rule id 3 end")
	var result SingleRuleDeteResult
	result.ruleID = 3
	//检测算法待实现
	result.err = nil
	result.result = false
	//
	out <- result
}

func BindRuleIdAndFunc() (err error){
	for k, v := range RuleIdReflect2Func {
		err = FuncSer.Bind(k, v)
		if err != nil {
			Logger.Error("BindRuleIdAndFunc Bind %s: %v", k, err)
			return
		}
	}
	return
}

func DetectDataByDetectionId(detectID string) (error) {
	deviceRule,err := DevRulesSer.GetDeviceRuleByDetectionID(detectID)
	if err != nil {
		Logger.Error("DetectDataByDetectionId GetDeviceRuleByDetectionID err:%v",err)
		return err
	}
	//todo 应该是一段一段检测的，但是段长是多少还没有办法确定,目前按照DataSliceLength来划分
	datas := DevStatusSer.GetStatusByDeviceID(deviceRule.DeviceID)
	begin := int64(0)
	end := int64(len(datas))
	sliceEnd := int64(0)
	if end == 0{
		return errors.New("not found")
	}
	detectResult := &model.DetectResult{
		DeviceID:deviceRule.DeviceID,
		DeviceName:datas[0].DeviceName,
		DeteRulesID:deviceRule.Detectrules,
		DetectionID:detectID,
	}
	for i := 0; begin < end; i ++{
		sliceEnd += consts.DataSliceLength
		if sliceEnd > end {
			sliceEnd = end
		}
		result,err := detectdataByExp(deviceRule.Detectrules,datas[begin:sliceEnd])
		if err != nil {
			return err
		}
		detectResult.StartId = begin
		detectResult.EndId	 = sliceEnd
		detectResult.SubDeteID = i
		detectResult.SourceMark = consts.MarkNomal
		detectResult.CreateTime = time.Now()
		detectResult.ModifyTime = time.Now()
		if result {
			detectResult.ResultMark = consts.MarkAbnormal
		}else {
			detectResult.ResultMark = consts.MarkNomal
		}
		for {
			err := DetResultSer.InsertResult(detectResult)
			if err == nil {
				break
			}
		}
		begin = sliceEnd
	}
	return err
}
func detectdataByExp(exp string,data []*model.DeviceStatus) (bool,error) {
	arr,rules := GenerateRPN(exp)
	var resultMap = make(map[string]bool)
	var resultChan = make(chan SingleRuleDeteResult,len(rules))
	for _,k := range rules{
		Logger.Info("call rule:%s",k)
		if err := FuncSer.Call(k,resultChan,data); err != nil {
			Logger.Error("detectdataByExp FuncSer.Call rule %s:error %v", k, err)
			return false,err
		}
	}
	for range rules{
		rst := <- resultChan
		if rst.err != nil {
			resultMap[fmt.Sprintf("%d",rst.ruleID)] = false
			Logger.Error("detectdataByExp channel receive rule result error,rule id:%d,err:%v",rst.ruleID,rst.err)
		}else {
			resultMap[fmt.Sprintf("%d",rst.ruleID)] = rst.result
		}
	}
	return calculateRPN(arr,resultMap)
}

func CreateDetectionOrder(deviceID uint64,rule_exp string) error {
	id,_ := IdGenSer.Next(consts.DetectionIdTag)
	//todo get device info to fill device name
	createRule := &model.DeviceRule{deviceID,fmt.Sprintf("device%d",deviceID),rule_exp,fmt.Sprintf("%d",id.MaxId),time.Now(),time.Now()}
	return DevRulesSer.InsertDeviceRule(createRule)
}
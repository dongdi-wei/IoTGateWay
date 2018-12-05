package service

import (
	"fmt"
	"testing"
)

func TestExternalIP(t *testing.T) {
	fmt.Println(ExternalIP())
}
func TestInterFaces(t *testing.T) {
	retList, err := InterFaces()
	if err != nil {
		Logger.Error("TestInterFaces InterFaces error:%v", err)
	} else {
		for _, k := range retList {
			Logger.Info("%v", k)
		}
	}
}

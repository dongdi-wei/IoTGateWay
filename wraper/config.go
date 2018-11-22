package wraper

import (
	"IoTGateWay/base"
	"sync"
)

type wrape struct {
}

var (Wraper *wrape
	 wrapOnce sync.Once
	 Logger *base.LogIot
)

func Init()  {
	wrapOnce.Do(func() {
		Wraper = new(wrape)
	})
	Logger = base.IotLogger
}

func init()  {
	Init()
}
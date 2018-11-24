package wraper

import (
	"IoTGateWay/base"
	"sync"
)

type Wrape struct {
}

var (Wraper *Wrape
	 wrapOnce sync.Once
	 Logger *base.LogIot
)

func Init()  {
	wrapOnce.Do(func() {
		Wraper = new(Wrape)
	})
	Logger = base.IotLogger
}

func init()  {
	Init()
}
package base

import (
	"log"
	"os"
	"sync"
)

var (IotLogger *LogIot
	logOnce sync.Once
)
func init()  {
	logOnce.Do(func() {
		file, err := os.OpenFile("IotGateWay.log",
			os.O_CREATE|os.O_WRONLY|os.O_APPEND,0666)
		if err != nil {
			log.Fatalln("failed to open error log file:", err)
		}
		flag := log.Llongfile
		IotLogger = NewWriterLogger(file,flag,3)
	})
}

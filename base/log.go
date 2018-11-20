package base

import (
	"fmt"
	"io"
	"log"
	"IoTGateWay/consts"
	"os"
)

type LogIot struct {
	level int
	err   *log.Logger
	warn  *log.Logger
	info  *log.Logger
	debug *log.Logger
	depth int
}

//获取一个logIot的实例
func NewWriterLogger(w io.Writer, flag int, depth int) *LogIot {
	logger := new(LogIot)
	logger.depth = depth
	if logger.depth <= 0 {
		logger.depth = 2
	}
	logger.err = log.New(io.MultiWriter(w,os.Stderr), "[Error] ", flag)
	logger.warn = log.New(io.MultiWriter(w,os.Stdout), "[Warning] ", flag)
	logger.info = log.New(io.MultiWriter(w,os.Stdout), "[Info] ", flag)
	logger.debug = log.New(io.MultiWriter(w,os.Stdout), "[Debug] ", flag)

	logger.SetLevel(consts.LevelInformational)

	return logger
}

//设置日志level
func (li *LogIot) SetLevel(l int) int {
	li.level = l
	return li.level
}

// 打印Error级别的日志.
func (ll *LogIot) Error(format string, v ...interface{}) {
	if consts.LevelError > ll.level {
		return
	}
	ll.err.Output(ll.depth, fmt.Sprintf(format, v...))
}

// 打印Warn级别的日志.
func (ll *LogIot) Warn(format string, v ...interface{}) {
	if consts.LevelWarning > ll.level {
		return
	}
	ll.warn.Output(ll.depth, fmt.Sprintf(format, v...))
}

// 打印Info级别的日志.
func (ll *LogIot) Info(format string, v ...interface{}) {
	if consts.LevelInformational > ll.level {
		return
	}
	ll.info.Output(ll.depth, fmt.Sprintf(format, v...))
}

// 打印Debug级别的日志.
func (ll *LogIot) Debug(format string, v ...interface{}) {
	if consts.LevelDebug > ll.level {
		return
	}
	ll.debug.Output(ll.depth, fmt.Sprintf(format, v...))
}

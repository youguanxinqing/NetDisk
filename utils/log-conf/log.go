package log_conf

import (
	"os"
	"runtime"
	"strconv"

	"github.com/sirupsen/logrus"
)

var f *os.File

// isDisk 是否将日子存储磁盘上
func isDisk(flag bool, path ...string) (*os.File, error) {
	if !flag {
		return os.Stdout, nil
	}

	var err error
	if f, err = os.OpenFile(
		path[0], os.O_WRONLY|os.O_CREATE|os.O_APPEND, 0666); err != nil {
		return nil, err
	}
	return f, nil
}

func init() {
	// 性能消耗严重，非生产环境可选择关闭
	logrus.SetReportCaller(true)
	formatter := logrus.TextFormatter{
		TimestampFormat: "2006/01/02 15:04:05",
		CallerPrettyfier: func(f *runtime.Frame) (function, file string) {
			function = "" // 关闭函数的显示
			file = f.File + ":" + strconv.Itoa(f.Line)
			return
		},
	}
	logrus.SetFormatter(&formatter)

	if f, err := isDisk(false, "test.log"); err == nil {
		logrus.SetOutput(f)
	}
}

package log

import (
	"io/ioutil"
	"log"
	"os"
	"sync"
)

var (
	// \033[31m红色 \033[0m无色 \033[34m蓝色
	// log.LstdFlags 显示时间日期, Lshortfile 显示文件名以及代码行号
	errorLog = log.New(os.Stdout, "\033[31m[error]\033[0m ", log.LstdFlags|log.Lshortfile)
	infoLog  = log.New(os.Stdout, "\033[34m[info ]\033[0m ", log.LstdFlags|log.Lshortfile)
	loggers  = []*log.Logger{errorLog, infoLog}
	mu       sync.Mutex
)

// log方法
var (
	Error  = errorLog.Println
	Errorf = errorLog.Printf
	Info   = infoLog.Println
	Infof  = infoLog.Printf
)

// Log等级
const (
	InfoLevel = iota
	ErrorLevel
	Disabled
)

// 设置Log等级
func SetLevel(level int) {
	// TODO 此处为什么要设置Lock Unlock
	mu.Lock()
	defer mu.Unlock()

	// 默认都输出到控制台
	for _, logger := range loggers {
		logger.SetOutput(os.Stdout)
	}

	if ErrorLevel < level {
		errorLog.SetOutput(ioutil.Discard)
	}

	if InfoLevel < level {
		infoLog.SetOutput(ioutil.Discard)
	}
}

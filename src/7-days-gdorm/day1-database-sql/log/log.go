package log

import (
	"io/ioutil"
	"log"
	"os"
	"sync"
)

// Author:Boyn
// Date:2020/3/23

// log instances
// Used Lshortfile to print line number and file name
var (
	errorLog = log.New(os.Stdout, "\033[31m[error]\033[0m", log.LstdFlags|log.Lshortfile) // red color
	infoLog  = log.New(os.Stdout, "\033[34m[info]\033[0m", log.LstdFlags|log.Lshortfile)  // blue color
	loggers  = []*log.Logger{errorLog, infoLog}
	mu       sync.Mutex
)

// log methods.
// explode 4 methods
var (
	Error  = errorLog.Println
	Errorf = errorLog.Printf
	Info   = infoLog.Println
	Infof  = infoLog.Printf
)

// Support level setting in log system
const (
	InfoLevel = iota
	ErrorLevel
	Disabled
)

// we can use SetLevel(level) to set the log output level
// for example, if we use SetLevel(1), then the info log will be disabled
func SetLevel(level int) {
	mu.Lock()
	defer mu.Unlock()

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

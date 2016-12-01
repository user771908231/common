package log

import (
	"fmt"
	"log"
	"runtime"
	"strconv"
	"time"
)

import (
	"misc/timer"
	"casino_common/common/cfg"
)

const (
	PRINT_INTERVAL = 300
)

type LEVEL int32

var logLevel LEVEL = 1
var maxFileSize int64
var maxFileCount int32
var consoleAppender bool = true
var debugAppender bool = false
var RollingFile bool = false
var maxLogSeq int64 = 500

var TraceLoger *FileLogger
var ErrorLogger *FileLogger

var traceLogStrChan chan string
var errorLogStrChan chan string

const (
	ALL LEVEL = iota
	DEBUG
	INFO
	WARN
	ERROR
	FATAL
	OFF
)

var LevelMapping = map[string]LEVEL{
	"ALL":   ALL,
	"DEBUG": DEBUG,
	"INFO":  INFO,
	"WARN":  WARN,
	"ERROR": ERROR,
	"FATAL": FATAL,
	"OFF":   OFF,
}

var LevelMappingS = map[byte]LEVEL{
	'0': ALL,
	'1': DEBUG,
	'2': INFO,
	'3': WARN,
	'4': ERROR,
	'5': FATAL,
	'6': OFF,
}

func SetConsole(isConsole bool) {
	consoleAppender = isConsole
}

func SetDebug(ableDebug bool) {
	debugAppender = ableDebug
}

func SetMaxLogSeq(maxSeq int64) {
	maxLogSeq = maxSeq
}

func SetLevel(_level LEVEL) {
	logLevel = _level
}

func SetMaxFileCount(maxNumber int32) {
	maxFileCount = maxNumber
}

func SetMaxFileSize(maxSize int64, _unit UNIT) {
	maxFileSize = maxSize * int64(_unit)
}

func InitLoggers(fileDir, fileName string) {
	fmt.Println("日志的目录:", fileDir)
	fmt.Println("日志的名字:", fileName)
	TraceLoger = new(FileLogger)
	TraceLoger.SetRollingFile(fileDir, fileName + ".log")
	traceLogStrChan = make(chan string, maxLogSeq)
	go TraceLoger.logWriter(traceLogStrChan)

	ErrorLogger = new(FileLogger)
	ErrorLogger.SetRollingFile(fileDir, fileName + ".err")
	errorLogStrChan = make(chan string, maxLogSeq)
	go ErrorLogger.logWriter(errorLogStrChan)
}

func (f *FileLogger) logWriter(logStrChan chan string) {

	//timer
	config := cfg.Get()
	printInterval := PRINT_INTERVAL
	if len(config["log_seq_interval"]) != 0 {
		printInterval, _ = strconv.Atoi(config["log_seq_interval"])
	}
	seqTimer := make(chan int32, 1)
	timer.Add(0, time.Now().Unix() + int64(printInterval), seqTimer)

	for {
		select {
		case logStr := <-logStrChan:

			P(logStr, f)
		case <-seqTimer:
		//			log.Println("================ LOG SEQ SIZE:", len(logStrChan), "==================")
			timer.Add(0, time.Now().Unix() + int64(printInterval), seqTimer)
		}
	}
}

func P(logStr string, logfile *FileLogger) {

	logfile.mu.RLock()
	defer logfile.mu.RUnlock()

	logfile.lg.Output(2, logStr)
	console(logStr)
}

func console(logStr string) {
	if consoleAppender {
		log.Println(logStr)
	}
}

func shortFileName(file string) string {
	short := file
	for i := len(file) - 1; i > 0; i-- {
		if file[i] == '/' {
			short = file[i + 1:]
			break
		}
	}
	return short
}

func Trace(format string, v ...interface{}) {
	_, file, line, _ := runtime.Caller(2) //calldepth=3
	if logLevel <= ALL {
		traceLogStrChan <- fmt.Sprintf("[%v:%v]", shortFileName(file), line) + fmt.Sprintf("\033[1;35m[TRACE] " + format + " \033[0m ", v...)
	}
}

func T(format string, v ...interface{}) {
	Trace(format, v...)
}

func Debug(format string, v ...interface{}) {
	if debugAppender {
		_, file, line, _ := runtime.Caller(1) //calldepth=3
		if logLevel <= DEBUG {
			traceLogStrChan <- fmt.Sprintf("[%v:%v]", shortFileName(file), line) + fmt.Sprintf("\033[1;35m[DEBUG] " + format + " \033[0m ", v...)
		}
	}
}

func D(format string, v ...interface{}) {
	if debugAppender {
		_, file, line, _ := runtime.Caller(1) //calldepth=3
		if logLevel <= DEBUG {
			traceLogStrChan <- fmt.Sprintf("[%v:%v]", shortFileName(file), line) + fmt.Sprintf("\033[1;35m[DEBUG] " + format + " \033[0m ", v...)
		}
	}

}

func Normal(format string, v ...interface{}) {
	_, file, line, _ := runtime.Caller(1) //calldepth=3
	if logLevel <= INFO {
		traceLogStrChan <- fmt.Sprintf("[%v:%v]", shortFileName(file), line) + fmt.Sprintf("[NORMAL] " + format, v...)
	}
}

func N(format string, v ...interface{}) {
	Normal(format, v...)
}

func Info(format string, v ...interface{}) {
	Normal(format, v...)
}

func I(format string, v ...interface{}) {
	Normal(format, v...)
}

func Warn(format string, v ...interface{}) {
	_, file, line, _ := runtime.Caller(1) //calldepth=3
	if logLevel <= WARN {
		traceLogStrChan <- fmt.Sprintf("[%v:%v]", shortFileName(file), line) + fmt.Sprintf("\033[1;33m[WARN] " + format + " \033[0m ", v...)
	}
}

func W(format string, v ...interface{}) {
	Warn(format, v...)
}

func Error(format string, v ...interface{}) {
	_, file, line, _ := runtime.Caller(1) //calldepth=3
	if logLevel <= ERROR {
		errorLogStrChan <- fmt.Sprintf("%v:%v]", shortFileName(file), line) + fmt.Sprintf("\033[1;4;31m[ERROR] " + format + " \033[0m ", v...)
		traceLogStrChan <- fmt.Sprintf("%v:%v]", shortFileName(file), line) + fmt.Sprintf("\033[1;4;31m[ERROR] " + format + " \033[0m ", v...)
	}
}

func E(format string, v ...interface{}) {
	_, file, line, _ := runtime.Caller(1) //calldepth=3
	if logLevel <= ERROR {
		errorLogStrChan <- fmt.Sprintf("%v:%v]", shortFileName(file), line) + fmt.Sprintf("\033[1;4;31m[ERROR] " + format + " \033[0m ", v...)
		traceLogStrChan <- fmt.Sprintf("%v:%v]", shortFileName(file), line) + fmt.Sprintf("\033[1;4;31m[ERROR] " + format + " \033[0m ", v...)
	}
}

func Fatal(format string, v ...interface{}) {
	_, file, line, _ := runtime.Caller(1) //calldepth=2
	if logLevel <= FATAL {
		errorLogStrChan <- fmt.Sprintf("%v:%v", shortFileName(file), line) + fmt.Sprintf("\033[1;4;31m[FATAL]" + format + " \033[0m ", v...)
		traceLogStrChan <- fmt.Sprintf("%v:%v", shortFileName(file), line) + fmt.Sprintf("\033[1;4;31m[FATAL]" + format + " \033[0m ", v...)
	}
}

func F(format string, v ...interface{}) {
	_, file, line, _ := runtime.Caller(1) //calldepth=2
	if logLevel <= FATAL {
		errorLogStrChan <- fmt.Sprintf("%v:%v", shortFileName(file), line) + fmt.Sprintf("\033[1;4;31m[FATAL]" + format + " \033[0m ", v...)
		traceLogStrChan <- fmt.Sprintf("%v:%v", shortFileName(file), line) + fmt.Sprintf("\033[1;4;31m[FATAL]" + format + " \033[0m ", v...)
	}
}

func Printf(format string, v ...interface{}) {
	log.Printf(format, v...)
}

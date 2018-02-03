// Package: log
// File:
// Created by mint
// Useage:
// DATE: 14-7-25 15:26
package log

import (
	"casino_common/common/cfg"
	"fmt"
	"log"
	"os"
	"runtime/debug"
	"strconv"
	"sync"
	"time"
)

const DATEFORMAT = "2006-01-02"

var dailyRolling bool = true

type UNIT int64

const (
	_       = iota
	KB UNIT = 1 << (iota * 10)
	MB
	GB
	TB
)

type FileLogger struct {
	dir      string
	filename string
	_suffix  int
	isCover  bool
	_date    *time.Time
	mu       *sync.RWMutex
	logfile  *os.File
	lg       *log.Logger
}

func (f *FileLogger) SetRollingFile(fileDir, fileName string) {

	RollingFile = true
	dailyRolling = false

	f.dir = fileDir
	f.filename = fileName
	f.isCover = false
	f.mu = new(sync.RWMutex)

	f.mu.Lock()
	defer f.mu.Unlock()

	for i := 1; i <= int(maxFileCount); i++ {
		if isExist(fileDir + "/" + fileName + "." + strconv.Itoa(i)) {
			f._suffix = i
		}

		break
	}

	if !f.isMustRename() {
		if !isExist(fileDir) {
			os.Mkdir(fileDir, 0755)
		}
		f.logfile, _ = os.OpenFile(fileDir+"/"+fileName, os.O_RDWR|os.O_APPEND|os.O_CREATE, 0666)
		f.lg = log.New(f.logfile, "", log.LstdFlags|log.Lmicroseconds)
	} else {
		f.rename()
	}

	go f.fileMonitor()
}

func (f *FileLogger) SetRollingDaily(fileDir, fileName string) {
	RollingFile = false
	dailyRolling = true

	t, _ := time.Parse(DATEFORMAT, time.Now().Format(DATEFORMAT))

	f.dir = fileDir
	f.filename = fileName
	f._date = &t
	f.isCover = false
	f.mu = new(sync.RWMutex)

	f.mu.Lock()
	defer f.mu.Unlock()

	if !f.isMustRename() {
		if !isExist(fileDir) {
			os.Mkdir(fileDir, 0755)
		}
		f.logfile, _ = os.OpenFile(fileDir+"/"+fileName, os.O_RDWR|os.O_APPEND|os.O_CREATE, 0666)
		f.lg = log.New(f.logfile, "", log.LstdFlags|log.Lmicroseconds)
	} else {
		f.rename()
	}
}

func (f *FileLogger) isMustRename() bool {
	if dailyRolling {
		t, _ := time.Parse(DATEFORMAT, time.Now().Format(DATEFORMAT))
		if t.After(*f._date) {
			return true
		}
	} else {
		if maxFileCount > 1 {
			if fileSize(f.dir+"/"+f.filename) >= maxFileSize {
				return true
			}
		}
	}
	return false
}

func (f *FileLogger) rename() {
	if dailyRolling {
		fn := f.dir + "/" + f.filename + "." + f._date.Format(DATEFORMAT)
		if !isExist(fn) && f.isMustRename() {
			if f.logfile != nil {
				f.logfile.Close()
			}
			err := os.Rename(f.dir+"/"+f.filename, fn)
			if err != nil {
				f.lg.Println("rename err", err.Error())
			}
			t, _ := time.Parse(DATEFORMAT, time.Now().Format(DATEFORMAT))
			f._date = &t
			f.logfile, _ = os.Create(f.dir + "/" + f.filename)
			f.lg = log.New(f.logfile, "", log.LstdFlags|log.Lmicroseconds)
		}
	} else {
		f.coverNextOne()
	}
}

func (f *FileLogger) nextSuffix() int {
	return int(f._suffix%int(maxFileCount) + 1)
}

func (f *FileLogger) coverNextOne() {
	f._suffix = f.nextSuffix()
	if f.logfile != nil {
		f.logfile.Close()
	}
	if isExist(f.dir + "/" + f.filename + "." + strconv.Itoa(int(f._suffix))) {
		os.Remove(f.dir + "/" + f.filename + "." + strconv.Itoa(int(f._suffix)))
	}
	os.Rename(f.dir+"/"+f.filename, f.dir+"/"+f.filename+"."+strconv.Itoa(int(f._suffix)))
	f.logfile, _ = os.Create(f.dir + "/" + f.filename)
	f.lg = log.New(f.logfile, "", log.LstdFlags|log.Lmicroseconds)
}

func isExist(path string) bool {
	_, err := os.Stat(path)
	return err == nil || os.IsExist(err)
}

func (f *FileLogger) fileMonitor() {
	defer func() {
		if err := recover(); err != nil {
			fmt.Printf("地址[fileMonitor panic]开始打印日志err[%v] \n", err) // 这里的err其实就是panic传入的内容，55
			debug.PrintStack()
		}
	}()
	config := cfg.Get()
	logScan := 10
	if len(config["log_scan_interval"]) != 0 {
		logScan, _ = strconv.Atoi(config["log_scan_interval"])
	}

	timer := time.NewTicker(time.Duration(logScan) * time.Second)
	for {
		select {
		case <-timer.C:
			f.fileCheck()
		}
	}
}

func (f *FileLogger) fileCheck() {
	defer func() {
		if err := recover(); err != nil {
			log.Println(err)
		}
	}()
	if f.isMustRename() {
		f.mu.Lock()
		defer f.mu.Unlock()
		f.rename()
	}
}

func catchError() {
	if err := recover(); err != nil {
		log.Printf("ERR: %v", err)
	}
}

func fileSize(file string) int64 {
	f, e := os.Stat(file)
	if e != nil {
		//		fmt.Println(e.Error())
		return 0
	}
	//	log.Printf("Log[%v]:%.3fM\n", file, float32(f.Size())/1024/1024)
	return f.Size()
}

package ctlog

import (
	"fmt"
	"github.com/golang/glog"
	"log"
	"os"
	"path/filepath"
	"strconv"
	"sync"
	"time"
)

var createLogDirOnce = sync.Once{}

// level of debug
func Debugln(v ...interface{}) {
	if logLevel >= levelDEBUG {
		log.SetPrefix("DEBUG\t")
		log.Output(2, fmt.Sprintln(v))
	}
}

// level of info
func Infoln(v ...interface{}) {
	if logLevel >= levelINFO {
		log.SetPrefix("INFO\t")
		log.Output(2, fmt.Sprintln(v))
	}
}

// level of warning
func Warningln(v ...interface{}) {
	log.SetPrefix("WARN\t")
	if logLevel >= levelWARN {
		log.Output(2, fmt.Sprintln(v))
	}
}

// level of fatal
func Fatalln(v ...interface{}) {
	log.SetPrefix("FATAL\t")
	if logLevel >= levelFATAL {
		log.Output(2, fmt.Sprintln(v))
	}
}

func createLogDir() {
	// ignore error
	os.MkdirAll(ctlog.logDir, os.ModePerm)
}

func (l *logT) logFileName() (link string) {
	// set log file name,defauld pid
	if l.userName == "" {
		l.userName = strconv.Itoa(os.Getpid()) + ".log"
	}

	t := time.Now()
	fileName := fmt.Sprintf("%s.%04d%02d%02d-%02d%02d%02d.log",
		l.userName,
		t.Year(),
		t.Month(),
		t.Day(),
		t.Hour(),
		t.Minute(),
		t.Second(),
	)
	return fileName
}

func (l *logT) createLogFile() (err error) {
	if l.logDir == "" {
		l.logDir = "../output"
	}
	createLogDirOnce.Do(createLogDir)
	l.logFilePath = filepath.Join(l.logDir, l.logFileName())
	sysLink := filepath.Join(l.logDir, l.userName)
	l.f, err = os.Create(l.logFilePath)
	if err != nil {
		return fmt.Errorf("Create Log File Fail: %v", err)
	}
	os.Remove(sysLink)
	os.Symlink(l.logFilePath, sysLink)
	return nil
}

func (l *logT) outPut(calldepth int, s string) error {
	return log.Output(calldepth, s)
}

func (l *logT) Write(buf []byte) (n int, err error) {
	if l.f == nil {
		l.createLogFile()
	}
	l.logRotation()
	return l.f.Write(buf)
}

func (l *logT) logRotation() error {
	logFileInfo, err := os.Stat(l.logFilePath)
	if err != nil {
		fmt.Errorf("logRotation,get file sixe error.%v\n", err)
	}
	if logFileInfo.Size() >= l.maxLogSize {
		fmt.Println(logFileInfo.Size(), l.maxLogSize)
		l.createLogFile()
	}
	return nil
}

func Test() {
	glog.Error()
}

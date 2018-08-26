package ctlog

import (
	"fmt"
	"log"
	"os"
	"strings"
)

const (
	levelDEBUG int = iota
	levelINFO
	levelWARN
	levelERROR
	levelFATAL
	MaxLogSize int64 = 100 * 1024 * 1024
)

var (
	logLevel int
)

type logT struct {
	logLevel    int
	logDir      string
	logFilePath string
	programName    string
	f           *os.File
	maxLogSize  int64
}

var ctlog = new(logT)

func init() {
	// set log out method
	log.SetFlags(log.Ldate | log.Lmicroseconds | log.Lshortfile)

	// log out method
	log.SetOutput(ctlog)

	// set log remote size
	ctlog.maxLogSize = MaxLogSize

}

func SetLogLevel(level string) {
	level = strings.ToLower(level)
	switch level {
	case "debug":
		{
			logLevel = levelDEBUG
		}
	case "info":
		{
			logLevel = levelINFO
		}
	case "warning":
		{
			logLevel = levelWARN
		}
	case "error":
		{
			logLevel = levelERROR
		}
	case "fatal":
		{
			logLevel = levelFATAL
		}
	default:
		fmt.Println("Only Support LogLevel: debug info warning error fatal,Use default(error)")
		logLevel = levelERROR
	}
}

func SetLogDir(logDir, programName string) {
	ctlog.logDir = logDir
	ctlog.programName = programName + ".log"
}

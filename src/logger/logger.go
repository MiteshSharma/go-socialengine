package logger

import (
	"fmt"
	"time"
	"os"
	"log"
	"strconv"
	"path"
)

const (
	Critical = iota
	Error
	Warn
	Info
	Debug
)

type Logger interface {
	SetLevel(level int)
	Debug(msg string)
	Info(msg string)
	Warn(msg string)
	Error(msg string)
	Critical(msg string)
}

type logger struct {
	level int
}

func (l *logger) SetLevel(level int) {
	l.level = level
}

func (l *logger) Debug(msg string) {
	l.write(msg, Debug)
}

func (l *logger) Info(msg string) {
	l.write(msg, Info)
}

func (l *logger) Warn(msg string) {
	l.write(msg, Warn)
}

func (l *logger) Error(msg string) {
	l.write(msg, Error)
}

func (l *logger) Critical(msg string) {
	l.write(msg, Critical)
}

func (l *logger) write(msg string, lvl int) {
	if lvl > l.level {
		l.writeLog(msg)
	}
}

var logHandler *logger;

func Get() *logger {
	if logHandler == nil {
		logHandler = new (logger)
	}
	return logHandler
}

func (l *logger) SetLogLevel(level int) {
	l.level = level
}

func (l *logger) writeLog(logMessage string) {
	f := l.fileHandler(l.getFilePath())
	
	log.SetOutput(f)
	log.Println(logMessage)
	defer f.Close()
}

func (l *logger) fileHandler(filePath string) *os.File {
	if err := os.MkdirAll(path.Dir(filePath), 0777); err != nil {
        fmt.Println("Unable to create directory for tagfile! - " + err.Error())
	}
	
	f, err := os.OpenFile(filePath, os.O_RDWR | os.O_CREATE | os.O_APPEND, 0777)
	if err != nil {
	    fmt.Println("error opening file: %v", err)
	}
	
	return f
}

func (l *logger) getFilePath() string {
	filePath := "../log/"+strconv.Itoa(time.Now().Year()) +"/"+ time.Now().Month().String() +"/"+ strconv.Itoa(time.Now().Day())+ ".log"
	return filePath
}
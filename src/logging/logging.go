package logging

import (
	"container/list"
	"fmt"
)

/**
base for all kind of loggers
*/
type ILogger interface {
	Info(message string)
	Warn(message string)
	Error(message string)

	GetInfoCount() int
	GetWarnCount() int
	GetErrCount() int

	GetLastInfo() *string
	GetLastWarn() *string
	GetLastError() *string
}

const (
	info = 1
	warn = 2
	err  = 3
)

type LogData struct {
	lastInfo  *string
	lastWarn  *string
	lastError *string

	infoCount int
	warnCount int
	errCount  int
}

func (ld *LogData) init() {
	ld.lastInfo = nil
	ld.lastWarn = nil
	ld.lastError = nil

	ld.infoCount = 0
	ld.warnCount = 0
	ld.errCount = 0
}

func (ld *LogData) GetLastInfo() *string {
	return ld.lastInfo
}

func (ld *LogData) GetLastWarn() *string {
	return ld.lastWarn
}

func (ld *LogData) GetLastError() *string {
	return ld.lastError
}

func (ld *LogData) GetInfoCount() int {
	return ld.infoCount
}

func (ld *LogData) GetWarnCount() int {
	return ld.warnCount
}

func (ld *LogData) GetErrCount() int {
	return ld.errCount
}

/**
the dummy logger
*/
type DummyLogger struct {
	LogData
}

func NewDummyLogger() *DummyLogger {
	logger := new(DummyLogger)
	logger.init()
	return logger
}

func (l *DummyLogger) Info(message string) {
	l.lastInfo = &message
	l.infoCount++
}

func (l *DummyLogger) Warn(message string) {
	l.lastWarn = &message
	l.warnCount++
}

func (l *DummyLogger) Error(message string) {
	l.lastError = &message
	l.errCount++
}

/**
the console logger
*/
type ConsoleLogger struct {
	LogData
}

func NewConsoleLogger() *ConsoleLogger {
	logger := new(ConsoleLogger)
	logger.init()
	return logger
}

func (l *ConsoleLogger) print(logType string, message string) {
	fmt.Println(logType, " : ", message)
}

func (l *ConsoleLogger) Info(message string) {
	print("INFO", message)
	l.lastInfo = &message
	l.infoCount++
}

func (l *ConsoleLogger) Warn(message string) {
	print("WARN", message)
	l.lastWarn = &message
	l.warnCount++
}

func (l *ConsoleLogger) Error(message string) {
	print("ERROR", message)
	l.lastError = &message
	l.errCount++
}

/*
	the memory logger
*/
type LogEntry struct {
	message string
	logType int
}

type MemoryLogger struct {
	LogData
	messages *list.List
	maxSize  int
}

func NewMemoryLogger(maxSize int) *MemoryLogger {
	return new(MemoryLogger).init(maxSize)
}

func (l *MemoryLogger) init(maxSize int) *MemoryLogger {
	l.LogData.init()
	l.messages = list.New()
	l.maxSize = maxSize
	return l
}

func (l *MemoryLogger) addMessage(logType int, message string) {
	logEntry := new(LogEntry)
	logEntry.message = message
	logEntry.logType = logType
	l.messages.PushBack(logEntry)

	if l.messages.Len() > l.maxSize {
		l.messages.Remove(l.messages.Front())
	}
}

func (l *MemoryLogger) Info(message string) {
	l.infoCount++
	l.lastInfo = &message
	l.addMessage(info, message)
}

func (l *MemoryLogger) Warn(message string) {
	l.warnCount++
	l.lastWarn = &message
	l.addMessage(warn, message)
}

func (l *MemoryLogger) Error(message string) {
	l.errCount++
	l.lastError = &message
	l.addMessage(err, message)
}

func (l *MemoryLogger) Size() int {
	return l.messages.Len()
}

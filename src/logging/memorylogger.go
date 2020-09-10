package logging

import "container/list"

// the memory logger

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

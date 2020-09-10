package logging

import "fmt"

// the console logger

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

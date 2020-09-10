package logging

// the dummy logger

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

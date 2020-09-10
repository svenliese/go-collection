package logging

// base for all kind of loggers

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

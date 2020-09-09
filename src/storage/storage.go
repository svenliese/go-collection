package storage

import "strconv"

// base for all kind of storages

type IStorage interface {
	PutInt(name string, value int)
	GetInt() int

	PutBool(name string, value bool)
	GetBool() bool

	PutDouble(name string, value float64)
	GetDouble() float64

	Flush()
	Close()
}

const (
	NotSupported = "not supported"
	EndOfFile    = "end of file"
)

func stringToInt(value string) int {
	intValue, err := strconv.Atoi(value)
	if err != nil {
		panic(err)
	}
	return intValue
}

func stringToBool(value string) bool {
	boolValue, err := strconv.ParseBool(value)
	if err != nil {
		panic(err)
	}
	return boolValue
}

func stringToDouble(value string) float64 {
	floatValue, err := strconv.ParseFloat(value, 64)
	if err != nil {
		panic(err)
	}
	return floatValue
}

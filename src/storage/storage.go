package storage

import (
	"container/list"
	"strconv"
	"strings"
)

/**
base for all kind of storages
*/

type IStorage interface {
	PutInt(name string, value int)
	GetInt() int

	PutBool(name string, value bool)
	GetBool() bool

	PutDouble(name string, value float64)
	GetDouble() float64
}

/**
memory storage
*/

type MemoryStorage struct {
	lines *list.List
}

func NewMemoryStorage() *MemoryStorage {
	return new(MemoryStorage).init()
}

func (l *MemoryStorage) init() *MemoryStorage {
	l.lines = list.New()
	return l
}

func (l *MemoryStorage) putLine(name string, value string) {
	l.lines.PushBack(name + "=" + value)
}

func (l *MemoryStorage) getValue() string {
	line := l.lines.Front().Value.(string)
	l.lines.Remove(l.lines.Front())
	stringSlice := strings.Split(line, "=")
	return stringSlice[1]
}

func (l *MemoryStorage) PutInt(name string, value int) {
	l.putLine(name, strconv.Itoa(value))
}

func (l *MemoryStorage) GetInt() int {
	value := l.getValue()
	intValue, err := strconv.Atoi(value)
	if err != nil {
		panic(err)
	}
	return intValue
}

func (l *MemoryStorage) PutBool(name string, value bool) {
	l.putLine(name, strconv.FormatBool(value))
}

func (l *MemoryStorage) GetBool() bool {
	value := l.getValue()
	boolValue, err := strconv.ParseBool(value)
	if err != nil {
		panic(err)
	}
	return boolValue
}

func (l *MemoryStorage) PutDouble(name string, value float64) {
	l.putLine(name, strconv.FormatFloat(value, 'e', 16, 64))
}

func (l *MemoryStorage) GetDouble() float64 {
	value := l.getValue()
	floatValue, err := strconv.ParseFloat(value, 64)
	if err != nil {
		panic(err)
	}
	return floatValue
}

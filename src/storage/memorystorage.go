package storage

import (
	"container/list"
	"strconv"
	"strings"
)

// storage implementation to read/write to/from memory

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
	return stringToInt(value)
}

func (l *MemoryStorage) PutBool(name string, value bool) {
	l.putLine(name, strconv.FormatBool(value))
}

func (l *MemoryStorage) GetBool() bool {
	value := l.getValue()
	return stringToBool(value)
}

func (l *MemoryStorage) PutDouble(name string, value float64) {
	l.putLine(name, strconv.FormatFloat(value, 'e', 16, 64))
}

func (l *MemoryStorage) GetDouble() float64 {
	value := l.getValue()
	return stringToDouble(value)
}

func (l *MemoryStorage) Flush() {
	// not necessary
}

func (l *MemoryStorage) Close() {
	// not necessary
}

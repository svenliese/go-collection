package storage

import (
	"bufio"
	"os"
	"strings"
)

// storage implementation to read from an file

type FileStorageReader struct {
	file   *os.File
	reader *bufio.Scanner
}

func NewMFileStorageReader(name string) *FileStorageReader {
	return new(FileStorageReader).init(name)
}

func (l *FileStorageReader) init(name string) *FileStorageReader {
	file, err := os.Open(name)
	if err != nil {
		panic(err)
	}
	l.file = file

	l.reader = bufio.NewScanner(file)

	return l
}

func (l *FileStorageReader) getValue() string {
	if l.reader.Scan() {
		stringSlice := strings.Split(l.reader.Text(), "=")
		return stringSlice[1]
	}
	panic(EndOfFile)
}

func (l *FileStorageReader) PutInt(name string, value int) {
	panic(NotSupported)
}

func (l *FileStorageReader) GetInt() int {
	value := l.getValue()
	return stringToInt(value)
}

func (l *FileStorageReader) PutBool(name string, value bool) {
	panic(NotSupported)
}

func (l *FileStorageReader) GetBool() bool {
	value := l.getValue()
	return stringToBool(value)
}

func (l *FileStorageReader) PutDouble(name string, value float64) {
	panic(NotSupported)
}

func (l *FileStorageReader) GetDouble() float64 {
	value := l.getValue()
	return stringToDouble(value)
}

func (l *FileStorageReader) Flush() {
	// not necessary
}

func (l *FileStorageReader) Close() {
	err := l.file.Close()
	if err != nil {
		panic(err)
	}
}

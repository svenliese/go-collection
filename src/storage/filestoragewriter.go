package storage

import (
	"os"
	"strconv"
)

// storage implementation to write to an file

type FileStorageWriter struct {
	file         *os.File
	bytesWritten int
}

func NewMFileStorageWriter(name string) *FileStorageWriter {
	return new(FileStorageWriter).init(name)
}

func (l *FileStorageWriter) init(name string) *FileStorageWriter {
	file, err := os.Create(name)
	if err != nil {
		panic(err)
	}
	l.file = file
	l.bytesWritten = 0
	return l
}

func (l *FileStorageWriter) putLine(name string, value string) {
	bytesWritten, err := l.file.WriteString(name + "=" + value + "\n")
	if err != nil {
		panic(err)
	}
	l.bytesWritten += bytesWritten
}

func (l *FileStorageWriter) PutInt(name string, value int) {
	l.putLine(name, strconv.Itoa(value))
}

func (l *FileStorageWriter) GetInt() int {
	panic(NotSupported)
}

func (l *FileStorageWriter) PutBool(name string, value bool) {
	l.putLine(name, strconv.FormatBool(value))
}

func (l *FileStorageWriter) GetBool() bool {
	panic(NotSupported)
}

func (l *FileStorageWriter) PutDouble(name string, value float64) {
	l.putLine(name, strconv.FormatFloat(value, 'e', 16, 64))
}

func (l *FileStorageWriter) GetDouble() float64 {
	panic(NotSupported)
}

func (l *FileStorageWriter) Flush() {
	err := l.file.Sync()
	if err != nil {
		panic(err)
	}
}

func (l *FileStorageWriter) Close() {
	err := l.file.Close()
	if err != nil {
		panic(err)
	}
}

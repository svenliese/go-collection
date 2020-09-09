package storage

import (
	"os"
	"testing"
)

func compareInt(t *testing.T, expected int, value int) {
	if expected != value {
		t.Error("expect", expected, "but got", value)
	}
}

func compareBool(t *testing.T, expected bool, value bool) {
	if expected != value {
		t.Error("expect", expected, "but got", value)
	}
}

func compareDouble(t *testing.T, expected float64, value float64) {
	if expected != value {
		t.Error("expect", expected, "but got", value)
	}
}

func checkInt(t *testing.T, writer IStorage, reader IStorage) {

	writer.PutInt("test", 10)
	writer.PutInt("test", 20)

	var value = reader.GetInt()
	compareInt(t, 10, value)

	value = reader.GetInt()
	compareInt(t, 20, value)
}

func checkBool(t *testing.T, writer IStorage, reader IStorage) {

	writer.PutBool("test", true)
	writer.PutBool("test", false)

	var value = reader.GetBool()
	compareBool(t, true, value)

	value = reader.GetBool()
	compareBool(t, false, value)
}

func checkDouble(t *testing.T, writer IStorage, reader IStorage) {

	writer.PutDouble("test", -0.1)
	writer.PutDouble("test", 5.004)

	var value = reader.GetDouble()
	compareDouble(t, -0.1, value)

	value = reader.GetDouble()
	compareDouble(t, 5.004, value)
}

func TestMemoryStorage(t *testing.T) {
	storage := NewMemoryStorage()
	checkInt(t, storage, storage)
	checkBool(t, storage, storage)
	checkDouble(t, storage, storage)
}

func TestFileStorage(t *testing.T) {
	fileName := os.TempDir() + string(os.PathSeparator) + "test.txt"
	writer := NewMFileStorageWriter(fileName)
	reader := NewMFileStorageReader(fileName)
	checkInt(t, writer, reader)
	checkBool(t, writer, reader)
	checkDouble(t, writer, reader)
}

package storage

import "testing"

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

func checkInt(t *testing.T, storage IStorage) {

	storage.PutInt("test", 10)
	storage.PutInt("test", 20)

	var value = storage.GetInt()
	compareInt(t, 10, value)

	value = storage.GetInt()
	compareInt(t, 20, value)
}

func checkBool(t *testing.T, storage IStorage) {

	storage.PutBool("test", true)
	storage.PutBool("test", false)

	var value = storage.GetBool()
	compareBool(t, true, value)

	value = storage.GetBool()
	compareBool(t, false, value)
}

func checkDouble(t *testing.T, storage IStorage) {

	storage.PutDouble("test", -0.1)
	storage.PutDouble("test", 5.004)

	var value = storage.GetDouble()
	compareDouble(t, -0.1, value)

	value = storage.GetDouble()
	compareDouble(t, 5.004, value)
}

func TestMemoryStorage(t *testing.T) {
	storage := NewMemoryStorage()
	checkInt(t, storage)
	checkBool(t, storage)
	checkDouble(t, storage)
}

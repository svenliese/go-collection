package logging

import (
	"testing"
)

func compare(t *testing.T, expected *string, value *string) {
	if expected != value {
		if (expected == nil && value != nil) || (expected != nil && value == nil) || (*expected != *value) {
			t.Error("expect", *expected, "but got", *value)
		}
	}
}

func checkLogger(
	t *testing.T,
	log ILogger,
	expectedInfoCount int,
	expectedWarnCount int,
	expectedErrorCount int,
	expectedLastInfo *string,
	expectedLastWarn *string,
	expectedLastError *string) {

	var count = log.GetInfoCount()
	if count != expectedInfoCount {
		t.Error("expect", 0, "but got", count)
	}

	count = log.GetWarnCount()
	if count != expectedWarnCount {
		t.Error("expect", 0, "but got", count)
	}

	count = log.GetErrCount()
	if count != expectedErrorCount {
		t.Error("expect", 0, "but got", count)
	}

	compare(t, expectedLastInfo, log.GetLastInfo())
	compare(t, expectedLastWarn, log.GetLastWarn())
	compare(t, expectedLastError, log.GetLastError())
}

func testLogger(
	t *testing.T,
	log ILogger) {

	checkLogger(t, log, 0, 0, 0, nil, nil, nil)

	info := "info"
	log.Info(info)
	checkLogger(t, log, 1, 0, 0, &info, nil, nil)

	warn := "warn"
	log.Warn(warn)
	checkLogger(t, log, 1, 1, 0, &info, &warn, nil)

	err := "error"
	log.Error(err)
	checkLogger(t, log, 1, 1, 1, &info, &warn, &err)
}

func TestDummyLogger(t *testing.T) {
	var log = NewDummyLogger()
	testLogger(t, log)
}

func TestConsoleLogger(t *testing.T) {
	var log = NewConsoleLogger()
	testLogger(t, log)
}

func TestMemoryLogger(t *testing.T) {
	log := NewMemoryLogger(3)
	testLogger(t, log)

	log.Info("new")
	var count = log.Size()
	if count != 3 {
		t.Error("expect", 0, "but got", count)
	}
}

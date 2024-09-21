package test

import (
	log "plentylog/internal/domain"
	"testing"
)

type MockDriver struct {
	recievedLogs []log.Record
	returnError  bool
}

func (m *MockDriver) WriteLog(record *log.Record) error {
	m.recievedLogs = append(m.recievedLogs, *record)
	return nil
}

func TestLogger_LogSuccess(t *testing.T) {
	logger := log.NewLogger()
	mockDriver := &MockDriver{}
	logger.AddDriver(mockDriver)

	logger.Log(log.Info, "test message", nil, "")

	if len(mockDriver.recievedLogs) != 1 {
		t.Errorf("Expected 1 log, got %d", len(mockDriver.recievedLogs))
	}

	if mockDriver.recievedLogs[0].Message != "test message" {
		t.Errorf("Expected log message to be 'test message', got %s", mockDriver.recievedLogs[0].Message)
	}
}

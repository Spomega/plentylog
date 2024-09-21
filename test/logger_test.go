package test

import (
	log "plentylog/internal/domain"
	"testing"
)

type MockDriver struct {
	receivedLogs []log.Record
}

func (m *MockDriver) WriteLog(record *log.Record) error {
	m.receivedLogs = append(m.receivedLogs, *record)
	return nil
}

func TestLogger_LogSuccess(t *testing.T) {
	logger := log.NewLogger()
	mockDriver := &MockDriver{}
	logger.AddDriver(mockDriver)

	logger.Log(log.Info, "test message", nil, "")

	if len(mockDriver.receivedLogs) != 1 {
		t.Errorf("Expected 1 log, got %d", len(mockDriver.receivedLogs))
	}

	receivedLog := mockDriver.receivedLogs[0]

	if receivedLog.Message != "test message" {
		t.Errorf("Expected log message to be 'test message', got %s", mockDriver.receivedLogs[0].Message)
	}

	if receivedLog.Level != log.Info {
		t.Errorf("Expected log level Info, got %d", receivedLog.Level)
	}
}

package test

import (
	log "github.com/Spomega/plentylog/pkg/domain"
	"reflect"
	"testing"
)

type MockDriver struct {
	receivedLogs []log.Record
}

func (m *MockDriver) WriteLog(record *log.Record) error {
	m.receivedLogs = append(m.receivedLogs, *record)
	return nil
}

func TestLogger_Log(t *testing.T) {
	logger := log.NewLogger()
	mockDriver := &MockDriver{}
	logger.AddDriver(mockDriver)

	txnAttributes := map[string]string{
		"customerId": "123",
		"operation":  "purchase",
		"itemId":     "456",
	}

	logger.Log(log.Info, "test Log message", txnAttributes, "")

	if len(mockDriver.receivedLogs) != 1 {
		t.Errorf("Expected 1 log, got %d", len(mockDriver.receivedLogs))
	}

	receivedLog := mockDriver.receivedLogs[0]

	if receivedLog.Message != "test Log message" {
		t.Errorf("Expected log message to be 'test message', got %s", mockDriver.receivedLogs[0].Message)
	}

	if receivedLog.Level != log.Info {
		t.Errorf("Expected log level Info, got %d", receivedLog.Level)
	}

	expectedAttributes := map[string]string{
		"customerId": "123",
		"operation":  "purchase",
		"itemId":     "456",
	}

	if !reflect.DeepEqual(receivedLog.MetaData, expectedAttributes) {
		t.Errorf("Expected attributes %v, got %v", expectedAttributes, receivedLog.MetaData)
	}
}

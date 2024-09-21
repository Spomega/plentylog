package test

import (
	log "plentylog/internal/domain"
	"testing"
)

func TestTransactionLogger_log(t *testing.T) {
	t.Parallel()
	logger := log.NewLogger()
	mockDriver := &MockDriver{}
	logger.AddDriver(mockDriver)

	transactionID := "txn123"
	transactionLogger := log.NewTransactionLogger(logger, transactionID)
	attributes := map[string]string{"key": "value"}
	transactionLogger.Log(log.Info, "test message", attributes)

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

	if receivedLog.TransactionID != "txn123" {
		t.Errorf("Expected transaction ID '%s', got '%s'", transactionID, receivedLog.TransactionID)
	}
	if receivedLog.MetaData["key"] != "value" {
		t.Errorf("Expected tag 'key' to be 'value', got '%s'", receivedLog.MetaData["key"])
	}
}

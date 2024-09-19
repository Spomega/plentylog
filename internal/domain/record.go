package log

import (
	"time"
)

// Record represents a log record.
type Record struct {
	Timestamp     time.Time         `json:"timestamp"`
	Level         Level             `json:"level"`
	Message       string            `json:"message"`
	MetaData      map[string]string `json:"meta_data,omitempty"`
	TransactionID string            `json:"transaction_id,omitempty"`
}

func NewRecord(level Level, message string, metaData map[string]string, transactionID string) *Record {
	return &Record{
		Timestamp:     time.Now(),
		Level:         level,
		Message:       message,
		MetaData:      metaData,
		TransactionID: transactionID,
	}
}

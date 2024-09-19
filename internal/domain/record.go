package log

import (
	"time"
)

// Record represents a log record.
type Record struct {
	Timestamp     time.Time         `json:"timestamp"`
	Level         Level             `json:"level"`
	Message       string            `json:"message"`
	Tags          map[string]string `json:"tags,omitempty"`
	TransactionID string            `json:"transaction_id,omitempty"`
}

func NewRecord(level Level, message string, tags map[string]string, transactionID string) *Record {
	return &Record{
		Timestamp:     time.Now(),
		Level:         level,
		Message:       message,
		Tags:          tags,
		TransactionID: transactionID,
	}
}

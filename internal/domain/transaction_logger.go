package log

// TransactionLogger wraps around the Logger and manages transaction ID.
type TransactionLogger struct {
	logger        *Logger
	transactionID string
}

// NewTransactionLogger creates a logger that logs entries for a specific transaction.
func NewTransactionLogger(logger *Logger, transactionID string) *TransactionLogger {
	return &TransactionLogger{
		logger:        logger,
		transactionID: transactionID,
	}
}

// Log sends logs with a specific transaction ID.
func (t *TransactionLogger) Log(level Level, message string, tags map[string]string) {
	t.logger.Log(level, message, tags, t.transactionID)
}

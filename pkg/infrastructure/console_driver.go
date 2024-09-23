package infrastructure

import (
	"fmt"
	log "github.com/Spomega/plentylog/pkg/domain"
)

// ConsoleDriver represents a console driver.
type ConsoleDriver struct{}

// WriteLog writes a log record to the console.
func (c *ConsoleDriver) WriteLog(record *log.Record) error {
	fmt.Printf("[%s] %s: %s", record.Timestamp.Format("2006-01-02T15:04:05Z07:00"), record.Level, record.Message)

	if record.TransactionID != "" {
		fmt.Printf(" [TransactionID: %s]", record.TransactionID)

	}

	if len(record.MetaData) > 0 {
		fmt.Printf("[Attributes: ")
		for key, value := range record.MetaData {
			fmt.Printf(" %s:%s,", key, value)
		}
		fmt.Printf("]")
	}

	fmt.Println()

	return nil
}

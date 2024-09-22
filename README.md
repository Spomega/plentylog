# PlentyLog - A Flexible Logging Package in Go

**PlentyLog** is a customizable and extensible logging package written in Go. It supports multiple logging output drivers (CLI, file) and can handle structured logging with log levels, metadata (tags), and transaction-specific logging.

## Features

- **Multiple Log Levels**: Supports different log levels such as `Debug`,`Info`, `Warning`, and `Error`.
- **Customizable Drivers**: Easily extend logging output to different mediums, including CLI and file drivers.
- **Transaction Logging**: Group logs under a specific transaction ID for better tracking of requests or processes.
- **Extensible**: Built with extensibility in mind, allowing additional drivers (e.g., database, remote services) to be easily added.

## Installation

To use PlentyLog in your Go project, you can install it by running:

```bash
go get github.com/Spomega/plentylog

```

## Usage

### Basic Logger Setup
For basic Logging you can use the DefaultLogger which logs to the CLI,JSON file and log file out of the box.The files are created in the current working directory.

```go
package main

import (
	"fmt"
	"github.com/Spomega/plentylog/config"
	log "github.com/Spomega/plentylog/pkg/domain"
)

func main() {
	logger, err := config.GetDefaultLogger()
	if err != nil {
		fmt.Printf("Error initializing logger: %v\n", err)
		return
	}

	logger.Log(log.Info, "test Log message", map[string]string{"customerId": "123", "operation": "purchase", "itemId": "456"}, "")
	
	err = logger.CloseAll()

	if err != nil {
		fmt.Printf("Error closing logger: %v\n", err)
	}

}
```
Example CLI output will be as shown below.
```bash
[2024-09-23T00:18:10+02:00] [INFO] test Log message [Attributes: customerId:123, operation:purchase, itemId:456]
```

### Custom Logging
If you want to customize the logging output and chose  particular output driver or drivers(CLI,JSON file, Log file), you can create a new logger by passing the desired configuration.A JSON file with the format below.
It can just one driver or multiple drivers.

```json
{
  "drivers": [
    {
      "type": "cli"
    },
    {
      "type": "json",
      "filename": "logs.json"
    },
    {
      "type":     "logfile",
      "filename": "logs.log"
    }
  ]
}
```
- drivers: Specifies the drivers to use. Currently, cli , json and logfile drivers are supported.
Example Driver Output
CLI Driver: Outputs logs directly to the console.
File Driver: Logs are written in JSON format to the specified file.
With the config file above you can create a custom logger as shown below.
```go
package main

import (
	"fmt"
	"github.com/Spomega/plentylog/config"
	log "github.com/Spomega/plentylog/pkg/domain"
)

func main() {
	logger, err := config.GetLoggerWithConfig("./config/config.json")
	if err != nil {
		fmt.Printf("Error initializing logger: %v\n", err)
		return
	}

	logger.Log(log.Info, "test Log message", map[string]string{"customerId": "123", "operation": "purchase", "itemId": "456"}, "")

	err = logger.CloseAll()

	if err != nil {
		fmt.Printf("Error closing logger: %v\n", err)
		return
	}
}
```
Example  json log will be as shown below.
```json
{"timestamp":"2024-09-22T19:52:47.159832+02:00","level":"INFO","message":"test Log message","meta_data":{"customerId":"123","itemId":"456","operation":"purchase"}
```
### Transaction Logging
To group logs under a specific transaction ID, you a transaction logger as shown below.
```go
package main

import (
	"fmt"
	"github.com/Spomega/plentylog/config"
	log "github.com/Spomega/plentylog/pkg/domain"
)

func main() {
	logger, err := config.GetLoggerWithConfig("./config/config.json")
	if err != nil {
		fmt.Printf("Error initializing logger: %v\n", err)
		return
	}

	log.NewTransactionLogger(logger, "123").Log(log.Info, "test Log message", map[string]string{"customerId": "123", "operation": "purchase", "itemId": "456"})

	err = logger.CloseAll()

	if err != nil {
		fmt.Printf("Error closing logger: %v\n", err)
	}

}
```
Example CLI output will be as shown below.
```bash
[2024-09-23T00:18:10+02:00] [INFO] [TransactionID: 123] test Log message [Attributes: customerId:123, operation:purchase, itemId:456]
```
Note:
 Remember to close the logger after you are done with logging to ensure all logs are flushed to the output drivers.Especially when using the file drivers.











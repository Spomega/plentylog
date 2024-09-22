# PlentyLog - A Flexible Logging Package in Go

**PlentyLog** is a customizable and extensible logging package written in Go. It supports multiple logging drivers (CLI, file) and can handle structured logging with log levels, metadata (tags), and transaction-specific logging.

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

### Basic Logging
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

### Custom Logging
If you want to customize the logging output and chose  particular output driver or drivers(CLI,JSON file, Log file), you can create a new logger by passing the desired configuration.A JSON file with the format below.
i.e it can just one driver or multiple drivers.

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

```go
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




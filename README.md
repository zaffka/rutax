# Russian TaxID (TIN, INN) Validation Package

[![Go Reference](https://pkg.go.dev/badge/github.com/zaffka/rutax.svg)](https://pkg.go.dev/github.com/zaffka/rutax)  
[![Go Report Card](https://goreportcard.com/badge/github.com/zaffka/rutax)](https://goreportcard.com/report/github.com/zaffka/rutax)  
[![Tests](https://github.com/zaffka/rutax/actions/workflows/tests.yaml/badge.svg)](https://github.com/zaffka/rutax/actions/workflows/tests.yaml)

The `rutax` package provides functionality for validating and parsing Russian Tax Identification Numbers (INN - Идентификационный Номер Налогоплательщика, TIN - Tax Identification Number).

## Features

- Parsing and validation of TaxID for individuals and legal entities
- Checksum verification according to the official algorithm

## Installation

```bash
go get github.com/zaffka/rutax@latest
```

## Usage
#### `rutax.ParseID(innStr string) (rutax.ID, error)`

Main package function:
1. Verifies string format compliance with INN requirements
2. Validates checksums at key positions in the parsed string
3. Returns a structure with INN data

The returned `ID` structure contains:
- `Num` - validated INN(TIN) number
- `IsLegal` - flag indicating whether it belongs to a legal entity

Example:
```go
package main

import (
	"fmt"
	"github.com/zaffka/rutax"
)

func main() {
	// Parsing TaxID
	id, err := rutax.ParseID("7710140679")
	if err != nil {
		fmt.Println("Error:", err)
		
		return
	}

	fmt.Printf("INN: %s, Legal entity: %v\n", id.Num, id.IsLegal)
}
```

### Errors

The package returns the following error types:
- `ErrIDIncorrect` - INN format error (length, numeric characters)
- `ErrChecksumFailed` - checksum validation error

## License

MIT
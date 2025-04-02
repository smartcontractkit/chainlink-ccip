# Carpenter
A tool for parsing and analyzing log files.

## Build
```sh
~$ go build ./...
```

## Help

```
~$ ./carpenter -h
~$ go run . -h
```

## Run

```
~$ ./carpenter < log.log
~$ go run . < log.log
```

# Customization

Carpenter is designed for customization via 'modes'. By implementing a new mode you can
modify how to display data while carpenter does the heavy lifting around parsing and configuration.

## Example: the "basic" mode

Create a new package and file at `/internal/format/basic/basic.go` and implement a `formatter`:
```go
func basicFormatter(data *parse.Data) {
	fmt.Println(data.GetMessage())
}
```

In the same file create a factory function and register it with carpenter:
```go
func init() {
	// Register the basic formatter by name.
	format.Register("basic", basicFormatterFactory, "Print logs to the console with minimal processing.")
}

func basicFormatterFactory(options format.Options) format.Formatter {
	// This simple formatter is a pure function, more advanced formats
	// can implement the Formatter interface.
	return format.NewWrappedFormat(basicFormatter)
}
```

Finally, trigger the `init` function to run during program startup by adding a line to main.go:
```go
import (
	"context"
	"fmt"
	"os"

	// Register the formatters
	_ "github.com/smartcontractkit/chainlink-ccip/cmd/carpenter/internal/format/basic"
	_ "github.com/smartcontractkit/chainlink-ccip/cmd/carpenter/internal/format/fancy"
	_ "github.com/smartcontractkit/chainlink-ccip/cmd/carpenter/internal/format/summary"
)
```

You can now run carpenter using the new format:
```sh
~$ go run . --format basic < log.log
```

## Advanced customization

When carpenter is finished processing logs, it checks if the formatter implements the `io.Closer` and will call `Close()` if possible.

This can be used to aggregate logs rather than processing them line by line. You can see this in use by the `summary` formatter.

# Logger

A simple logging layer on top of Go's log pkg

## Installation

The typical `go get github.com/caglar10ur/gologger` will install gologger.

## Example usage

```go
package main

import "github.com/caglar10ur/gologger"

func main() {
    l := logger.New(nil)
    l.Infoln("...INFO LEVEL LOG...")
    l.Debugln("...DEBUG LEVEL LOG...")

    // Change logging level to Debug
    l.SetLogLevel(logger.Debug)

    // Enable tracing
    l.EnableTraceOutput()
    l.Debugln("...DEBUG LEVEL LOG...")
}
```

## Status

[![Build Status](https://secure.travis-ci.org/caglar10ur/gologger.png)](http://travis-ci.org/caglar10ur/gologger)

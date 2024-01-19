# SimpleWG

[![Go Reference](https://img.shields.io/badge/godoc-reference-%23007d9c.svg)](https://point-c.github.io/simplewg)

## Overview
SimpleWG is a Go package that provides a simplified wrapper around the standard `sync.WaitGroup`. It is designed to make working with goroutines more straightforward by automating the handling of `sync.WaitGroup.Add` and `sync.WaitGroup.Done`. This package is ideal for scenarios where you need to manage the lifecycle of multiple goroutines easily.

## Installation

To use SimpleWG in your Go project, install it using `go get`:

```bash
go get github.com/point-c/simplewg
```

## Usage

### Importing the Package

First, import the SimpleWG package into your Go file:

```go
import "github.com/point-c/simplewg"
```

### Creating a New WaitGroup

Create a new instance of `Wg`:

```go
var wg simplewg.Wg
```

### Running Goroutines

Use the `Go` method to run functions in separate goroutines. This method automatically handles the addition of goroutines to the waitgroup:

```go
wg.Go(func() {
    // Your goroutine logic here
})
```

### Waiting for Goroutines to Complete

To wait for all goroutines to complete, use the `Wait` method:

```go
wg.Wait()
```

### Wait for Goroutines to Complete with Timeout

o wait for goroutines to complete while also considering a timeout or other event, use the Done method in combination with a select statement.

```go
var wg simplewg.Wg
// Use Go to start goroutines...
ctx, cancel := context.WithTimeout(context.Background(), time.Minute)
defer cancel()
select {
case <-ctx.Done():
case <-wg.Done():
}
```

## Example

Here's a simple example demonstrating the usage of the SimpleWG package:

```go
package main

import (
    "fmt"
    "github.com/yourusername/simplewg"
)

func main() {
    var wg simplewg.Wg

    wg.Go(func() {
        fmt.Println("Hello from goroutine 1")
    })

    wg.Go(func() {
        fmt.Println("Hello from goroutine 2")
    })

    wg.Wait()
    // or
    select {
    case <-wg.Done():
    }
	
    fmt.Println("All goroutines completed")
}
```

## Testing

The package includes tests that demonstrate its functionality. Use Go's testing tools to run the tests:

```bash
go test
```

## Godocs

To regenerate godocs:

```bash
go generate -tags docs ./...
```
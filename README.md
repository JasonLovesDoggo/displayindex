# Go Display Index

[![CI Status](https://github.com/jasonlovesdoggo/displayindex/workflows/Test/badge.svg)](https://github.com/jasonlovesdoggo/displayindex/actions)
[![Go Report Card](https://goreportcard.com/badge/github.com/jasonlovesdoggo/displayindex)](https://goreportcard.com/report/github.com/jasonlovesdoggo/displayindex)
[![Go Reference](https://pkg.go.dev/badge/github.com/jasonlovesdoggo/displayindex.svg)](https://pkg.go.dev/github.com/jasonlovesdoggo/displayindex)

A cross-platform Go package to detect which display contains the cursor.

## Features
- Supports Windows, macOS, and Linux
- Simple API

## Installation
```bash
go get github.com/jasonlovesdoggo/displayindex
```

## Usage
```go
package main

import (
	"fmt"
	"github.com/jasonlovesdoggo/displayindex"
	"github.com/kbinani/screenshot"
)

func main() {
	index, err := displayindex.CurrentDisplayIndex()
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	fmt.Printf("Cursor is on display %d\n", index)
	img, err := screenshot.CaptureDisplay(index)  /* Example how you can use the index*/
    if err != nil {
        fmt.Println("Error:", err)
        return
    }
	// Do something with the image
}
```

## Requirements
- **Linux**: Requires X11 development libraries
  ```bash
  sudo apt-get install libx11-dev
  ```
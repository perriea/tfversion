package main

import (
  "fmt"
  "runtime"
)

func ShowVersion()  {
    fmt.Println("Version: 0.0.2")
    fmt.Printf("Family OS: %s\n", runtime.GOOS)
    fmt.Printf("Arch processor: %s\n", runtime.GOARCH)
}

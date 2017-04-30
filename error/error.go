package tferror

import (
  "errors"
  "fmt"

  "github.com/fatih/color"
)

var (
    // Errors
    info    *color.Color
    good    *color.Color
    warn    *color.Color
    fatal   *color.Color
    err     error
)

func init()  {
    info  = color.New(color.FgBlue, color.Bold)
    good  = color.New(color.FgGreen, color.Bold)
    warn  = color.New(color.FgYellow, color.Bold)
    fatal = color.New(color.FgRed, color.Bold)
}

func Run(level int, message string)  {

    err = errors.New(message)
    if err != nil {
        switch level {
          case 0:
              info.Println(err)
          case 1:
              good.Println(err)
          case 2:
              warn.Println(err)
          case 3:
              fatal.Println(err)
          default:
              fmt.Println(err)
        }
    }
}

func Panic(err error)  {

    if err != nil {
        panic(err)
    }
}

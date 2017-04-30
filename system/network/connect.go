package tfnetwork

import (
    "net"
    "time"
)

var (
    hostname string
    seconds  int
    err      error
)

func init()  {

    hostname = "releases.hashicorp.com:80"
    seconds  = 5
}

func Run() bool {

    timeOut := time.Duration(seconds) * time.Second
    _, err := net.DialTimeout("tcp", hostname, timeOut)

    if err != nil {
        return false
    }

    return true
 }

package tfnetwork

import (
	"fmt"
	"net"
	"time"
)

var (
	resp net.Conn
	err  error
)

// Run : Lauch test internet connection
func Run(hostname string, timeout int, verbose bool) bool {

	if hostname == "" {
		hostname = "releases.hashicorp.com:80"
	}
	if timeout == 0 {
		timeout = 3
	}

	calcTimeOut := time.Duration(timeout) * time.Second
	resp, err = net.DialTimeout("tcp", hostname, calcTimeOut)

	if verbose {
		fmt.Printf("Request %s\nRemoteAddr: %s\nLocalAddr: %s\n\n", hostname, resp.RemoteAddr(), resp.LocalAddr())
	}

	if err != nil {
		return false
	}

	return true
}

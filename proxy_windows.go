package main

import (
	"fmt"
	"net"
	"strings"
	//	"syscall"
)

var _ = fmt.Println

func isErrConnReset(err error) bool {
	if ne, ok := err.(net.Error); ok {
		return ne.Temporary() && ne.Timeout()
	}
	return false
}

func isDNSError(err error) bool {
	/*
		fmt.Printf("calling isDNSError for err type: %v %s\n",
			reflect.TypeOf(err), err.Error())
	*/
	// DNS error are not of type DNSError on Windows, so I used this ugly
	// hack.
	errMsg := err.Error()
	return strings.Contains(errMsg, "No such host") ||
		strings.Contains(errMsg, "GetAddrInfoW") ||
		strings.Contains(errMsg, "dial tcp")
}

func isErrOpWrite(err error) bool {
	ne, ok := err.(*net.OpError)
	if !ok {
		return false
	}
	return ne.Op == "WSASend"
}

func isErrOpRead(err error) bool {
	ne, ok := err.(*net.OpError)
	if !ok {
		return false
	}
	return ne.Op == "WSARecv"
}

func isErrTooManyOpenFd(err error) bool {
	// TODO implement this.
	return false
}

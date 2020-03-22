package main

import (
	"fmt"
	"net"
)

func pingRespondClientHandler(daction string, ddata string, conn net.Conn) {
	fmt.Println(ddata)
}
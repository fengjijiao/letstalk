package main

import (
	"fmt"
	"net"
)

func pingRespondServerHandler(daction string, ddata string, conn net.Conn) {
	fmt.Println("sent \n")
	transmitData(1, daction, ddata, conn)
}
package main

import (
    "fmt"
    "net"
)

func runAsServer() {
    fmt.Println("Starting the server ...")
    // 创建 listener
    listener, err := net.Listen("tcp", conf.LocalListenAddress)
    if err != nil {
        fmt.Println("Error listening", err.Error())
        return //终止程序
    }
    // 监听并接受来自客户端的连接
    for {
        conn, err := listener.Accept()
        if err != nil {
            fmt.Println("Error accepting", err.Error())
            return // 终止程序
        }
        go receiveDataForServer(conn)
    }
}

func receiveDataForServer(conn net.Conn) {
    for {
        buf := make([]byte, 512)
        len, err := conn.Read(buf)
        if err != nil {
            fmt.Println("Error reading", err.Error())
            return // 终止程序
        }
        fmt.Printf("Received data: %v\n", string(buf[:len]))
        switchDTypeForServer(string(buf[:len]), conn)
    }
}

// 分析数据包格式
func switchDTypeForServer(data string,conn net.Conn) {
    dtype, daction, ddata := parseData(data)
    if(dtype == 1) {
        switchFuncForServer(daction, ddata, conn)
    }else {
        fmt.Println("An error occurred while analyse data format.")
    }
}

// 分流功能
func switchFuncForServer(daction string, ddata string, conn net.Conn) {
    if daction == "ping" {
        pingRespondServerHandler(daction, ddata, conn)
    }
}
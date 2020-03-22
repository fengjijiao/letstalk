package main

import (
    "bufio"
    "fmt"
    "net"
    "os"
    "strings"
)

func runAsClient() {
    //打开连接:
    conn, err := net.Dial("tcp", conf.RemoteServerAddress)
    if err != nil {
        //由于目标计算机积极拒绝而无法创建连接
        fmt.Println("Error dialing", err.Error())
        return // 终止程序
    }

    go receiveDataForClient(conn)

    inputReader := bufio.NewReader(os.Stdin)
    fmt.Println("First, what is your name?")
    clientName, _ := inputReader.ReadString('\n')
    // fmt.Printf("CLIENTNAME %s", clientName)
    trimmedClient := strings.Trim(clientName, "\r\n") // Windows 平台下用 "\r\n"，Linux平台下使用 "\n"
    // 给服务器发送信息直到程序退出：
    for {
        fmt.Println("What to send to the server? Type Q to quit.")
        input, _ := inputReader.ReadString('\n')
        trimmedInput := strings.Trim(input, "\r\n")
        // fmt.Printf("input:--s%--", input)
        // fmt.Printf("trimmedInput:--s%--", trimmedInput)
        if trimmedInput == "Q" {
            return
        }
        transmitData(1, "ping", trimmedClient + " says: " + trimmedInput, conn)
    }
}

func receiveDataForClient(conn net.Conn) {
    for {
        reply := make([]byte, 1024)
        _, err := conn.Read(reply)
        if err != nil {
            fmt.Println("Read to server failed:", err.Error())
            return
        }
        switchDTypeForClient(string(reply), conn)
    }
}

// 分析数据包格式
func switchDTypeForClient(data string,conn net.Conn) {
    dtype, daction, ddata := parseData(data)
    if(dtype == 1) {
        switchFuncForClient(daction, ddata, conn)
    }else {
        fmt.Println("An error occurred while analyse data format.")
    }
}

// 分流功能
func switchFuncForClient(daction string, ddata string, conn net.Conn) {
    if daction == "ping" {
        pingRespondClientHandler(daction, ddata, conn)
    }
}
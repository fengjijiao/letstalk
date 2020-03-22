package main

import (
	"fmt"
	"net"
	"strconv"
	"strings"
)

// 解析数据
func parseData(data string) (int64, string, string) {
    res := strings.Split(data, "|")
    //fmt.Printf("%v\n", res)
    if len(res) < 3 {
        return 0, "error", "received data was wrong."
    }
    dtype, err := strconv.ParseInt(res[0], 10, 64)
    if err != nil {
            return 0, "error", "received data was wrong."
    }
    var daction string = res[1]
    var ddata string = res[2]
    return dtype, daction, ddata
}

func checkError(error error, info string) {
    if error != nil {
        panic("ERROR: " + info + " " + error.Error()) // terminate program
    }
}

func transmitData(dtype int, daction string, ddata string, conn net.Conn) {
    data := generateDataPacket(dtype, daction, ddata)
    _, err := conn.Write([]byte(data))
    if err != nil {
        fmt.Println("An error occurred while transmitDataForServer.")
        return
    }
}

// 生成数据包
func generateDataPacket(dtype int, daction string, ddata string) string {
    return strconv.Itoa(dtype) + "|" + daction + "|" + ddata
}

// 关闭指定来自客户端的连接
func closeClientConnection(conn net.Conn) {
    conn.Close()
}
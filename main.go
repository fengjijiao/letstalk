package main

import (
	"fmt"
	"github.com/nadoo/conflag"
)

// conflag
var conf struct {
    isServer bool
    isClient bool
    LocalListenAddress string
    RemoteServerAddress string
    isGenKey bool
}

func init() {
	flag := conflag.New()
    flag.BoolVar(&conf.isServer,"s",false,"Whether it is a server")
    flag.BoolVar(&conf.isClient,"c",false,"Whether it is a client")
    flag.StringVar(&conf.LocalListenAddress,"l","0.0.0.0:9000","Local Listen Address")
    flag.StringVar(&conf.RemoteServerAddress,"r","127.0.0.1:9000","Remote Server Address")
    flag.BoolVar(&conf.isGenKey,"g",false,"Whether generate keys")
    flag.Parse()
}

func main() {
	if conf.isServer && conf.isClient {
        fmt.Println("It is not allowed to run server and client at the same time!")
    }else if conf.isServer && !conf.isClient {
        runAsServer()
    }else if !conf.isServer && conf.isClient {
        runAsClient()
    }else if conf.isGenKey {
        //
    }else {
        fmt.Println("No operating mode specified!")
    }
}
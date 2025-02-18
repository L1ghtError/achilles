package main

import (
	"ahls_srvi/internal/infrastructure"
	"flag"
)

const (
	DefaultAppaAddr       = "127.0.0.1:8081"
	DefaultAhlcServerAddr = "127.0.0.1:8080"
)

var (
	appAddrFlag  = flag.String("appAddr", DefaultAppaAddr, `Address of this application, default is: `+DefaultAppaAddr)
	ahlcAddrFlag = flag.String("ahlcAddr", DefaultAhlcServerAddr, `Address of ahlc application, default is: `+DefaultAhlcServerAddr)
)

func main() {
	flag.Parse()
	jsonServerAddr := ""

	infrastructure.RunL(*appAddrFlag, *ahlcAddrFlag, jsonServerAddr)
}

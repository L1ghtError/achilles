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
	helpFlag     = flag.Bool("help", false, `Prints "help" message`)
	appAddrFlag  = flag.String("appAddr", DefaultAppaAddr, `Address of this application`)
	ahlcAddrFlag = flag.String("ahlcAddr", DefaultAhlcServerAddr, `Address of ahlc application`)
)

func main() {
	flag.Parse()
	if *helpFlag {
		flag.PrintDefaults()
		return
	}
	jsonServerAddr := ""

	infrastructure.RunL(*appAddrFlag, *ahlcAddrFlag, jsonServerAddr)
}

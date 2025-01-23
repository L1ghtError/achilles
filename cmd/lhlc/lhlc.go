package main

import (
	"ahls_srvi/internal/infrastructure"
	"os"
)

const DEFAULT_APP_ADDR = "localhost:8081"
const DEFAULT_AHLC_ADDR = "http://localhost:8080"
const DEFAULT_JSON_SERVER_ADDR = "http://localhost:3000"

func main() {
	appAddr := DEFAULT_APP_ADDR
	ahlcAddr := DEFAULT_AHLC_ADDR
	jsonServerAddr := DEFAULT_JSON_SERVER_ADDR
	if len(os.Args) > 2 {
		appAddr = os.Args[1]
	}
	if len(os.Args) > 3 {
		ahlcAddr = os.Args[2]
	}
	if len(os.Args) > 4 {
		jsonServerAddr = os.Args[3]
	}
	infrastructure.RunL(appAddr, ahlcAddr, jsonServerAddr)
}

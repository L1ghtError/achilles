package main

import (
	"ahls_srvi/internal/infrastructure"
	"ahls_srvi/internal/storage"
	"log"
	"os"
	"strconv"
)

const DEFAULT_APP_ADDR = "localhost:8080"
const DEFAULT_JSON_SERVER_ADDR = "http://localhost:3000"
const DEFAULT_CACHE_UNINX_TTL = "12"

func main() {
	appAddr := DEFAULT_APP_ADDR
	jsonServerAddr := DEFAULT_JSON_SERVER_ADDR
	cacheUnixTTLStr := DEFAULT_CACHE_UNINX_TTL

	if len(os.Args) > 2 {
		appAddr = os.Args[1]
	}
	if len(os.Args) > 3 {
		jsonServerAddr = os.Args[2]
	}
	if len(os.Args) > 4 {
		cacheUnixTTLStr = os.Args[3]
	}
	// Validate CACHE_UNIX_TTL
	cacheUnixTTL, err := strconv.Atoi(cacheUnixTTLStr)
	if err != nil {
		log.Fatalf("Invalid CACHE_UNIX_TTL value: %v", err)
	}

	db := storage.JsonServerDrv{Hostname: jsonServerAddr}
	cache := storage.NewACache(int64(cacheUnixTTL), &db)

	infrastructure.Run(appAddr, &cache)
}

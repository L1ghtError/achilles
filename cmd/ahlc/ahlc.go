package main

import (
	"ahls_srvi/internal/infrastructure"
	"ahls_srvi/internal/storage"
	"flag"
	"log"
	"reflect"
	"strings"
)

// Constants for database types
const (
	DBTypeCassandra  = "cassandra"
	DBTypeJSONServer = "json-server"
)

// Constants for database addresses
const (
	DefaultCassandraAddr  = "127.0.0.1:9042"
	DefaultJSONServerAddr = "http://localhost:3000"
)

// Default configuration values
const (
	DefaultAppAddr      = "localhost:8080"
	DefaultDBType       = DBTypeCassandra
	DefaultDBServerAddr = DefaultCassandraAddr
	DefaultCacheUnixTTL = 12
)

// Command-line flags
var (
	helpFlag         = flag.Bool("help", false, `Prints "help" message`)
	appAddrFlag      = flag.String("appAddr", DefaultAppAddr, `Address of this application`)
	dbServerTypeFlag = flag.String("dbServerType", DefaultDBType, `DB that stores data, it can be either "cassandra" or "json-server"`)
	dbServerAddrFlag = flag.String("dbServerAddr", DefaultDBServerAddr, `DB address, for exmple: "localhost:3000"`)
	cacheTTL         = flag.Int64("cacheTTL", DefaultCacheUnixTTL, `TTL (time-to-live) for server cache in seconds`)
)

func main() {
	flag.Parse()
	if *helpFlag {
		flag.PrintDefaults()
		return
	}
	var db storage.ICustodian
	if *dbServerTypeFlag == DBTypeJSONServer {
		if !strings.Contains(*dbServerAddrFlag, "http://") &&
			!strings.Contains(*dbServerAddrFlag, "https://") {
			log.Fatalf("%s should contain eather http or https", *dbServerAddrFlag)
		}
		db = &storage.JsonServerDrv{Hostname: *dbServerAddrFlag}

	} else if *dbServerTypeFlag == DBTypeCassandra {
		x := storage.CassandraServerOps{Hostnames: []string{*dbServerAddrFlag}, Keyspace: "ahlc"}
		var dbClose func()
		db, dbClose = storage.NewCassandraServerDrv(x)
		if reflect.ValueOf(db).IsNil() {
			log.Fatalf("Cannot connect to %s %s", *dbServerTypeFlag, *dbServerAddrFlag)
		}
		defer dbClose()

	} else {
		log.Fatalf("%s is not supported type of db server", *dbServerTypeFlag)
	}
	cache := storage.NewACache(*cacheTTL, db)

	infrastructure.Run(*appAddrFlag, &cache)
}

package infrastructure

import (
	"ahls_srvi/internal/handler"
	"ahls_srvi/internal/middleware"
	"ahls_srvi/internal/queries"
	"log"
	"net/http"
)

func RunL(appAddr, ahlcAddr, dbAddr string) {
	router := http.NewServeMux()

	cli := queries.LClient{Hostname: ahlcAddr, DBaddr: dbAddr}
	hndl := handler.AdHandler{C: &cli}

	router.HandleFunc("POST /stitching.m3u8", hndl.StichM3U8)
	stack := middleware.CreateStack(middleware.Logging)

	lhls := &http.Server{
		Addr:         appAddr,
		ReadTimeout:  DEFAULT_TIMEOUT,
		WriteTimeout: DEFAULT_TIMEOUT,
		Handler:      stack(router),
	}
	gracefulShutdown(lhls)
	log.Println("Server listening on", lhls.Addr)
	if err := lhls.ListenAndServe(); err != http.ErrServerClosed {
		log.Fatalf("HTTP server ListenAndServe: %v", err)
	}
}

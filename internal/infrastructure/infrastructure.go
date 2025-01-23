package infrastructure

import (
	"context"
	"log"
	"net/http"
	"os"
	"os/signal"
	"time"
)

const DEFAULT_TIMEOUT time.Duration = time.Second * 10

func gracefulShutdown(srv *http.Server) {
	go func() {
		sigint := make(chan os.Signal, 1)
		signal.Notify(sigint, os.Interrupt)
		<-sigint
		if err := srv.Shutdown(context.Background()); err != nil {
			log.Printf("HTTP server Shutdown: %v", err)
		}
	}()

}

package infrastructure

import (
	"ahls_srvi/internal/handler"
	"ahls_srvi/internal/middleware"
	"ahls_srvi/internal/storage"
	"context"
	"log"
	"net/http"
	"os"
	"os/signal"
	"time"
)

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
func Run(appAddr string, c storage.ICustodian) {
	router := http.NewServeMux()

	ah := handler.AppHandler{Cache: c}
	router.HandleFunc("/auction", ah.GetAppropriateContent)

	stack := middleware.CreateStack(middleware.Logging)
	ahls := &http.Server{
		Addr:         appAddr,
		ReadTimeout:  10 * time.Second,
		WriteTimeout: 10 * time.Second,
		Handler:      stack(router),
	}
	gracefulShutdown(ahls)

	log.Println("Server listening on", ahls.Addr)
	if err := ahls.ListenAndServe(); err != http.ErrServerClosed {
		log.Fatalf("HTTP server ListenAndServe: %v", err)
	}
}

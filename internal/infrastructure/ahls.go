package infrastructure

import (
	"ahls_srvi/internal/handler"
	"ahls_srvi/internal/middleware"
	"ahls_srvi/internal/storage"
	"log"
	"net/http"
)

func Run(appAddr string, c storage.ICustodian) {
	router := http.NewServeMux()

	ah := handler.AppHandler{Cache: c}
	router.HandleFunc("/auction", ah.GetAppropriateContent)

	stack := middleware.CreateStack(middleware.Logging)
	ahls := &http.Server{
		Addr:         appAddr,
		ReadTimeout:  DEFAULT_TIMEOUT,
		WriteTimeout: DEFAULT_TIMEOUT,
		Handler:      stack(router),
	}
	gracefulShutdown(ahls)

	log.Println("Server listening on", ahls.Addr)
	if err := ahls.ListenAndServe(); err != http.ErrServerClosed {
		log.Fatalf("HTTP server ListenAndServe: %v", err)
	}
}

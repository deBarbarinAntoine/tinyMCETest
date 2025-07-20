package main

import (
	"context"
	"fmt"
	"log"
	"log/slog"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"
)

var logger = slog.New(slog.NewTextHandler(os.Stdout, &slog.HandlerOptions{Level: slog.LevelDebug}))

func main() {
	port := 3030
	
	mux := router()
	
	server := &http.Server{
		Addr:    fmt.Sprintf(":%d", port),
		Handler: mux,
	}
	
	sig := make(chan os.Signal, 1)
	signal.Notify(sig, os.Interrupt, syscall.SIGTERM, syscall.SIGQUIT)
	
	go func() {
		<-sig
		
		ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
		defer cancel()
		
		logger.Info("Shutting down server...")
		
		err := server.Shutdown(ctx)
		if err != nil {
			log.Fatal(err)
		}
	}()
	
	logger.Info(fmt.Sprintf("Server listening on port %d", port))
	
	log.Fatal(server.ListenAndServe())
}

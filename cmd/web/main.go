package main

import (
	"context"
	"fmt"
	"io/fs"
	"log"
	"log/slog"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"
	
	"github.com/alexedwards/flow"
	"tinyMCETest/ui"
)

var logger = slog.New(slog.NewTextHandler(os.Stdout, &slog.HandlerOptions{Level: slog.LevelDebug}))

func main() {
	port := 3030
	mux := flow.New()
	
	staticFS, err := fs.Sub(ui.NodeModules, "assets/js/node_modules")
	if err != nil {
		logger.Error(err.Error())
		os.Exit(1)
	}
	mux.Handle("/static/...", http.StripPrefix("/static/", http.FileServer(http.FS(staticFS))), http.MethodGet)
	
	uploadsFile, err := os.OpenRoot("uploads")
	if err != nil {
		logger.Error(err.Error())
		os.Exit(1)
	}
	uploadFS, err := fs.Sub(uploadsFile.FS(), ".")
	if err != nil {
		logger.Error(err.Error())
		os.Exit(1)
	}
	mux.Handle("/uploads/...", http.StripPrefix("/uploads/", http.FileServer(http.FS(uploadFS))), http.MethodGet)
	
	mux.HandleFunc("/", home, http.MethodGet)
	mux.HandleFunc("/upload", upload, http.MethodPost)
	mux.HandleFunc("/save", save, http.MethodPost)
	
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
	// file, err := uploadFS.Open("GreatRuler_titlescreen.jpg")
	// if err != nil {
	// 	logger.Error(err.Error())
	// }
	// logger.Debug(fmt.Sprintf("file: %+v", file))
	log.Fatal(server.ListenAndServe())
}

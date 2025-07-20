package main

import (
	"io/fs"
	"net/http"
	"os"
	
	"github.com/alexedwards/flow"
	"tinyMCETest/ui"
)

func router() *flow.Mux {
	
	mux := flow.New()
	
	jsFS, err := fs.Sub(ui.NodeModules, "assets/node_modules")
	if err != nil {
		logger.Error(err.Error())
		os.Exit(1)
	}
	mux.Handle("/static/js/...", http.StripPrefix("/static/js/", http.FileServer(http.FS(jsFS))), http.MethodGet)
	
	cssFS, err := fs.Sub(ui.CSS, "assets/css")
	if err != nil {
		logger.Error(err.Error())
		os.Exit(1)
	}
	mux.Handle("/static/css/...", http.StripPrefix("/static/css/", http.FileServer(http.FS(cssFS))), http.MethodGet)
	
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
	mux.HandleFunc("/menu", menu, http.MethodGet)
	mux.HandleFunc("/upload", upload, http.MethodPost)
	mux.HandleFunc("/save", save, http.MethodPost)
	
	return mux
}

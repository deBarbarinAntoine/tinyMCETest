package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"html/template"
	"io"
	"mime/multipart"
	"net/http"
	"os"
	"path"
	"path/filepath"
	"strings"
	"time"
	
	"tinyMCETest/ui"
)

func render(w http.ResponseWriter, blockName string, data any) {
	
	// creating a bytes Buffer
	buf := new(bytes.Buffer)
	
	tmpl, err := template.ParseFS(ui.Templates, "templates/*.tmpl")
	if err != nil {
		logger.Error(err.Error())
		return
	}
	
	// executing the template in the buffer to catch any possible parsing error,
	// so that the user doesn't see a half-empty page
	err = tmpl.ExecuteTemplate(w, blockName, data)
	if err != nil {
		logger.Error(err.Error())
		return
	}
	
	buf.WriteTo(w)
}

func ajaxResponse(w http.ResponseWriter, status int, data envelope) {
	
	// initializing the data if necessary
	if data == nil {
		data = envelope{"status": status}
	}
	
	jsonResponse(w, status, data)
}

func ajaxError(w http.ResponseWriter, err error) {
	
	logger.Error(err.Error())
	
	jsonResponse(w, http.StatusInternalServerError, envelope{"error": "internal server error"})
}

func jsonResponse(w http.ResponseWriter, status int, data envelope) {
	
	// marshalling the resData
	jsonData, err := json.Marshal(data)
	if err != nil {
		logger.Error(err.Error())
		return
	}
	
	// setting the Content-Type header to JSON
	w.Header().Set("Content-Type", "application/json")
	
	// setting the Status response
	w.WriteHeader(status)
	
	// send the response with the JSON data
	_, err = w.Write(jsonData)
	if err != nil {
		logger.Error(err.Error())
	}
}

// Upload get a file and put it in the uploads directory
func Upload(file multipart.File, header *multipart.FileHeader) (string, string, error) {
	
	// extracting the file name and extension
	_, filename := path.Split(header.Filename)
	ext := path.Ext(filename)
	filename = strings.TrimSuffix(filename, ext)
	
	// creating the destination file
	filename = filepath.Join("uploads", fmt.Sprint(filename, "_", time.Now().Format("2006-01-02T15:04"), ext))
	dst, err := os.Create(filename)
	if err != nil {
		return "", "", fmt.Errorf("error creating file: %w", err)
	}
	defer dst.Close()
	
	// upload the file to destination path
	nbBytes, err := io.Copy(dst, file)
	if err != nil {
		return "", "", fmt.Errorf("error copying file: %w", err)
	}
	
	// return the message
	return fmt.Sprintf("%d bytes copied to %s", nbBytes, filename), filename, nil
}

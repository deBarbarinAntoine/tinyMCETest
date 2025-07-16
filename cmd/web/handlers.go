package main

import (
	"fmt"
	"net/http"
	"os"
	"time"
)

func home(w http.ResponseWriter, r *http.Request) {
	
	logger.Info(fmt.Sprintf("%s %s", r.Method, r.URL.Path))
	
	render(w, "home")
}

func upload(w http.ResponseWriter, r *http.Request) {
	
	logger.Info(fmt.Sprintf("%s %s", r.Method, r.URL.Path))
	
	// getting the file from the form
	file, header, err := r.FormFile("file")
	if err != nil {
		ajaxError(w, err)
		return
	}
	defer file.Close()
	
	// uploading the file
	msg, filePath, err := Upload(file, header)
	if err != nil {
		ajaxError(w, err)
		return
	}
	
	// DEBUG
	logger.Debug(msg)
	
	// respond with a validation message and the image name
	ajaxResponse(w, http.StatusOK, envelope{
		"message":  msg,
		"location": fmt.Sprintf("/%s", filePath),
	})
}

func save(w http.ResponseWriter, r *http.Request) {
	logger.Info(fmt.Sprintf("%s %s", r.Method, r.URL.Path))
	
	var content, author, date, termsAndConditions string
	
	err := r.ParseForm()
	if err != nil {
		logger.Error(err.Error())
		ajaxError(w, err)
		return
	}
	if r.Form.Has("content") {
		content = r.FormValue("content")
		logger.Info(fmt.Sprintf("content: %s", content))
	} else {
		logger.Error("no content found!")
	}
	if r.Form.Has("author") {
		author = r.FormValue("author")
		logger.Info(fmt.Sprintf("author: %s", author))
	} else {
		logger.Error("no author found!")
	}
	if r.Form.Has("date") {
		date = r.FormValue("date")
		logger.Info(fmt.Sprintf("date: %s", date))
	} else {
		logger.Error("no date found!")
	}
	if r.Form.Has("terms-and-conditions") {
		termsAndConditions = r.FormValue("terms-and-conditions")
		logger.Info(fmt.Sprintf("terms-and-conditions: %s", termsAndConditions))
	} else {
		logger.Error("no terms-and-conditions found!")
	}
	
	err = os.WriteFile(fmt.Sprintf("uploads/post-%s.html", time.Now().Format(time.RFC3339)), []byte(content), 0644)
	if err != nil {
		logger.Error(err.Error())
		ajaxError(w, err)
		return
	}
	
	ajaxResponse(w, http.StatusOK, envelope{"message": "save successfully!"})
}

package main

import (
	"fmt"
	"net/http"
	"os"
	"time"
)

func home(w http.ResponseWriter, r *http.Request) {
	
	logger.Info(fmt.Sprintf("%s %s", r.Method, r.URL.Path))
	
	render(w, "home", nil)
}

func menu(w http.ResponseWriter, r *http.Request) {
	
	logger.Info(fmt.Sprintf("%s %s", r.Method, r.URL.Path))
	
	var data = MenuData{
		Title: "Projects",
		H1:    "My personal projects",
		Article: Article{
			Title:       "MangaThorg",
			ImageURL:    "/uploads/MangaThorg-EndingMaker.png",
			Description: "A Manga website done in my first year using the API MangaDex",
			CreatedAt:   time.Now().Format("02 Jan 2006"),
			UpdatedAt:   time.Now().Format("02 Jan 2006"),
		},
		Articles: []Article{
			{
				Title:     "The Great Ruler",
				ImageURL:  "/uploads/GreatRuler_titlescreen.jpg",
				CreatedAt: time.Now().Format("02 Jan 2006"),
				UpdatedAt: time.Now().Format("02 Jan 2006"),
			},
			{
				Title:     "Link and the Guardians",
				ImageURL:  "/uploads/zelda_linkarcguardian_2025-07-19T12:46.jpg",
				CreatedAt: time.Now().Format("02 Jan 2006"),
				UpdatedAt: time.Now().Format("02 Jan 2006"),
			},
			{
				Title:     "Zelda - Breath of the Wild",
				ImageURL:  "/uploads/zelda_titlescreenswordlandscape_2025-07-19T12:49.jpg",
				CreatedAt: time.Now().Format("02 Jan 2006"),
				UpdatedAt: time.Now().Format("02 Jan 2006"),
			},
			{
				Title:     "Link with an arc",
				ImageURL:  "/uploads/zelda_linkarc2_2025-07-17T00:09.jpg",
				CreatedAt: time.Now().Format("02 Jan 2006"),
				UpdatedAt: time.Now().Format("02 Jan 2006"),
			},
		},
	}
	
	render(w, "menu", data)
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
		"location": fmt.Sprintf("%s", filePath),
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

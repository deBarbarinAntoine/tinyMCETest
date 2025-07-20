package main

// envelope data type for JSON responses
type envelope map[string]any

type Article struct {
	Title       string
	ImageURL    string
	Description string
	CreatedAt   string
	UpdatedAt   string
}

type MenuData struct {
	Title    string
	H1       string
	Article  Article
	Articles []Article
}

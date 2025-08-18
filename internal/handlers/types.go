package handlers

import "github.com/olemart1n/nub/internal/db"

type Notification struct {
	Error   bool
	Message string
}

type ContextKey string

const UserIDKey ContextKey = "userID"

const ImageUrsKey ContextKey = "imageURLs"

type TemplateDataIndex struct {
	IsLoggedIn bool
	UserID     string
	Title      string
	Page       string // Identifies which template to include
}

type TemplateDataPost struct {
	Index  TemplateDataIndex
	Images []db.Image
	Post   db.Post
}

type TemplateDataUpload struct {
	Index     TemplateDataIndex
	countries []string
}

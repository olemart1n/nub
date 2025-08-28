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
	Query      string
}

type TemplateDataPost struct {
	Index  TemplateDataIndex
	Images []db.Image
	Post   db.Post
}

type TemplateDataUpload struct {
	Index     TemplateDataIndex
	Countries []string
}

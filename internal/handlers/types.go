package handlers

type Notification struct {
	Error   bool
	Message string
}

type ContextKey string

const UserIDKey ContextKey = "userID"

const ImageUrsKey ContextKey = "imageURLs"

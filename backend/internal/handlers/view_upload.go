package handlers

import (
	"fmt"
	"html/template"
	"net/http"

	contextkeys "github.com/olemart1n/nub/internal/handlers/context-keys"
	"github.com/olemart1n/nub/utils"
)

func ViewUpload(tpl *template.Template) http.HandlerFunc {

	return func(w http.ResponseWriter, r *http.Request) {
		var data TemplateDataUpload

		data.Index.Title = "Create a post"
		data.Countries = utils.Countries
		data.Index.UserID = r.Context().Value(contextkeys.UserIDKey).(string)
		data.Index.IsLoggedIn = r.Context().Value(contextkeys.IsLoggedInKey).(bool)

		err := tpl.ExecuteTemplate(w, "upload.html", data)
		if err != nil {
			fmt.Print(err)
			http.Error(w, "error when executing post.html", http.StatusInternalServerError)
		}

	}
}

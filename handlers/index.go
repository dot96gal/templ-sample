package handlers

import (
	"net/http"

	"github.com/dot96gal/templ-sample/pages"
)

func IndexHandler(w http.ResponseWriter, r *http.Request) {
	page := pages.IndexPage()
	err := page.Render(r.Context(), w)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

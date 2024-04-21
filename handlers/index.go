package handlers

import (
	"fmt"
	"net/http"

	"github.com/dot96gal/templ-sample/pages"
	"github.com/dot96gal/templ-sample/storage"
)

type IndexHandler struct {
	state          *storage.State
	handlePostPath string
}

func NewIndexHandler(state *storage.State, handlePostPath string) IndexHandler {
	return IndexHandler{
		state:          state,
		handlePostPath: handlePostPath,
	}
}

func (h *IndexHandler) Get(w http.ResponseWriter, r *http.Request) {
	page := pages.IndexPage(h.state.Count(), h.handlePostPath)
	err := page.Render(r.Context(), w)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

func (h *IndexHandler) Post(w http.ResponseWriter, r *http.Request) {
	err := r.ParseForm()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}

	operation := r.PostForm.Get("operation")

	switch operation {
	case "up":
		h.state.CountUp()
	case "down":
		h.state.CountDown()
	}

	fmt.Fprintf(w, "%d", h.state.Count())
}

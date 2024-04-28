package handlers

import (
	"net/http"

	"github.com/dot96gal/templ-sample/components"
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

	req := components.ParseSimpleCounterRequest(r.PostForm)
	switch req.Operation {
	case components.SIMPLE_COUNTER_OPERATION_UP:
		h.state.CountUp()
	case components.SIMPLE_COUNTER_OPERATION_DOWN:
		h.state.CountDown()
	}

	component := components.SimpleCounterIndicator(h.state.Count())
	err = component.Render(r.Context(), w)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

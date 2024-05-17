package handlers

import (
	"fmt"
	"net/http"

	"github.com/dot96gal/templ-sample/components"
	"github.com/dot96gal/templ-sample/pages"
	"github.com/dot96gal/templ-sample/storage"
)

type IndexHandler struct {
	mux   *http.ServeMux
	path  string
	state *storage.State
}

func NewIndexHandler(path string, state *storage.State) IndexHandler {
	mux := http.NewServeMux()

	h := IndexHandler{
		mux:   mux,
		path:  path,
		state: state,
	}

	mux.HandleFunc(fmt.Sprintf("GET %s", h.path), h.handleGet)
	mux.HandleFunc(fmt.Sprintf("POST %s", h.path), h.handlePost)

	return h
}

func (h *IndexHandler) Pattern() string {
	return h.path
}

func (h *IndexHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	h.mux.ServeHTTP(w, r)
}

func (h *IndexHandler) handleGet(w http.ResponseWriter, r *http.Request) {
	props := pages.IndexPageProps{
		Count:          h.state.Count(),
		HandlePostPath: h.path,
	}
	page := pages.IndexPage(props)
	err := page.Render(r.Context(), w)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

func (h *IndexHandler) handlePost(w http.ResponseWriter, r *http.Request) {
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

	props := components.SimpleCounterIndicatorProps{
		Count: h.state.Count(),
	}
	component := components.SimpleCounterIndicator(props)
	err = component.Render(r.Context(), w)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

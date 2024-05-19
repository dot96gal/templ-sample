package handlers

import (
	"fmt"
	"net/http"

	"github.com/dot96gal/templ-sample/components"
	"github.com/dot96gal/templ-sample/pages"
	"github.com/dot96gal/templ-sample/storage"
)

type IndexHandler struct {
	path  string
	state *storage.State
}

func NewIndexHandler(path string, state *storage.State) IndexHandler {
	return IndexHandler{
		path:  path,
		state: state,
	}
}

func (h *IndexHandler) Register(mux *http.ServeMux) {
	mux.HandleFunc(fmt.Sprintf("GET %s", h.path), h.GetIndexPage)
	mux.HandleFunc(fmt.Sprintf("POST %s", h.path), h.PostSimpleCounter)
}

func (h *IndexHandler) GetIndexPage(w http.ResponseWriter, r *http.Request) {
	props := pages.IndexPageProps{
		Count:                 h.state.Count(),
		EndpointSimpleCounter: h.path,
	}

	page := pages.IndexPage(props)
	err := page.Render(r.Context(), w)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

func (h *IndexHandler) PostSimpleCounter(w http.ResponseWriter, r *http.Request) {
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

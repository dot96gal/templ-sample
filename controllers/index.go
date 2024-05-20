package controllers

import (
	"fmt"
	"net/http"

	"github.com/dot96gal/templ-sample/components"
	"github.com/dot96gal/templ-sample/pages"
	"github.com/dot96gal/templ-sample/storage"
)

type IndexController struct {
	path  string
	state *storage.State
}

func NewIndexController(path string, state *storage.State) IndexController {
	return IndexController{
		path:  path,
		state: state,
	}
}

func (c *IndexController) Register(mux *http.ServeMux) {
	mux.HandleFunc(fmt.Sprintf("GET %s", c.path), c.GetIndexPage)
	mux.HandleFunc(fmt.Sprintf("POST %s", c.path), c.PostSimpleCounter)
}

func (c *IndexController) GetIndexPage(w http.ResponseWriter, r *http.Request) {
	props := pages.IndexPageProps{
		Count:                 c.state.Count(),
		EndpointSimpleCounter: c.path,
	}

	page := pages.IndexPage(props)
	err := page.Render(r.Context(), w)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

func (c *IndexController) PostSimpleCounter(w http.ResponseWriter, r *http.Request) {
	err := r.ParseForm()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}

	req := components.ParseSimpleCounterRequest(r.PostForm)
	switch req.Operation {
	case components.SIMPLE_COUNTER_OPERATION_UP:
		c.state.CountUp()
	case components.SIMPLE_COUNTER_OPERATION_DOWN:
		c.state.CountDown()
	}

	props := components.SimpleCounterIndicatorProps{
		Count: c.state.Count(),
	}

	component := components.SimpleCounterIndicator(props)
	err = component.Render(r.Context(), w)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

package handlers

import (
	"bytes"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/PuerkitoBio/goquery"
	"github.com/dot96gal/templ-sample/components"
	"github.com/dot96gal/templ-sample/storage"
)

func TestIndexHandler_Get(t *testing.T) {
	state := storage.NewState()
	handlePostPath := "/"
	indexHandler := NewIndexHandler(state, handlePostPath)

	w := httptest.NewRecorder()
	r := httptest.NewRequest(http.MethodGet, "/", nil)

	indexHandler.Get(w, r)
	doc, err := goquery.NewDocumentFromReader(w.Result().Body)
	if err != nil {
		t.Fatalf("failed to read template: %v", err)
	}

	expected := "index page content"
	if actual := doc.Text(); !strings.Contains(actual, expected) {
		t.Errorf("expected %q, got %q", expected, actual)
	}
}

func TestIndexHandler_Post(t *testing.T) {
	state := storage.NewState()
	handlePostPath := "/"
	indexHandler := NewIndexHandler(state, handlePostPath)

	w := httptest.NewRecorder()
	r := httptest.NewRequest(
		http.MethodPost,
		"/",
		bytes.NewBuffer(
			[]byte(
				components.NewSimpleCounterRequestJSON(components.SIMPLE_COUNTER_OPERATION_UP)),
		),
	)

	indexHandler.Post(w, r)

	doc, err := goquery.NewDocumentFromReader(w.Result().Body)
	if err != nil {
		t.Fatalf("failed to read template: %v", err)
	}

	expected := "1"
	if actual := doc.Text(); actual != expected {
		t.Errorf("expected %q, got %q", expected, actual)
	}
}

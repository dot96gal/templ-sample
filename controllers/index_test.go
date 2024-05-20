package controllers

import (
	"net/http"
	"net/http/httptest"
	"net/url"
	"strings"
	"testing"

	"github.com/PuerkitoBio/goquery"
	"github.com/dot96gal/templ-sample/components"
	"github.com/dot96gal/templ-sample/storage"
)

func TestIndexController_GetIndexPage(t *testing.T) {
	path := "/"
	state := storage.NewState()
	controller := NewIndexController(path, state)

	w := httptest.NewRecorder()
	r := httptest.NewRequest(http.MethodGet, path, nil)

	controller.GetIndexPage(w, r)

	doc, err := goquery.NewDocumentFromReader(w.Result().Body)
	if err != nil {
		t.Fatalf("failed to read template: %v", err)
	}
	expected := "index page content"
	if actual := doc.Text(); !strings.Contains(actual, expected) {
		t.Errorf("expected %q, got %q", expected, actual)
	}
}

func TestIndexController_PostSimpleCounter_CountUp(t *testing.T) {
	path := "/"
	state := storage.NewState()
	controller := NewIndexController(path, state)

	form := url.Values{}
	form.Add(components.SIMPLE_COUNTER_OPERATION_KEY, string(components.SIMPLE_COUNTER_OPERATION_UP))
	body := strings.NewReader(form.Encode())

	w := httptest.NewRecorder()
	r := httptest.NewRequest(http.MethodPost, path, body)
	r.Header.Set("Content-Type", "application/x-www-form-urlencoded")

	controller.PostSimpleCounter(w, r)

	doc, err := goquery.NewDocumentFromReader(w.Result().Body)
	if err != nil {
		t.Fatalf("failed to read template: %v", err)
	}

	expected := "2"
	if actual := doc.Find(`[class^="simpleCounterCountIndicatorStyle"]`).Text(); actual != expected {
		t.Errorf("expected %q, got %q", expected, actual)
	}
}

func TestIndexController_PostSimpleCounter_CountDown(t *testing.T) {
	path := "/"
	state := storage.NewState()
	controller := NewIndexController(path, state)

	form := url.Values{}
	form.Add(components.SIMPLE_COUNTER_OPERATION_KEY, string(components.SIMPLE_COUNTER_OPERATION_DOWN))
	body := strings.NewReader(form.Encode())

	w := httptest.NewRecorder()
	r := httptest.NewRequest(http.MethodPost, path, body)
	r.Header.Set("Content-Type", "application/x-www-form-urlencoded")

	controller.PostSimpleCounter(w, r)

	doc, err := goquery.NewDocumentFromReader(w.Result().Body)
	if err != nil {
		t.Fatalf("failed to read template: %v", err)
	}

	expected := "0"
	if actual := doc.Find(`[class^="simpleCounterCountIndicatorStyle"]`).Text(); actual != expected {
		t.Errorf("expected %q, got %q", expected, actual)
	}
}

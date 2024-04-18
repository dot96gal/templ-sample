package components

import (
	"context"
	"io"
	"testing"

	"github.com/PuerkitoBio/goquery"
)

func TestHeader(t *testing.T) {
	r, w := io.Pipe()

	go func() {
		_ = Header("index page title").Render(context.Background(), w)
		_ = w.Close()
	}()

	doc, err := goquery.NewDocumentFromReader(r)
	if err != nil {
		t.Fatalf("failed to read template: %v", err)
	}

	component := doc.Find(`[data-testid="header-component"]`)
	if component.Length() == 0 {
		t.Error("expected data-testid attribute to be rendered, but it wasn't")
	}

	expected := "index page title"
	if actual := component.Find("h1").Text(); actual != expected {
		t.Errorf("expected %q, got %q", expected, actual)
	}
}

package pages

import (
	"context"
	"io"
	"testing"

	"github.com/PuerkitoBio/goquery"
)

func TestHeader(t *testing.T) {
	r, w := io.Pipe()

	go func() {
		_ = IndexPage(0, "/").Render(context.Background(), w)
		_ = w.Close()
	}()

	doc, err := goquery.NewDocumentFromReader(r)
	if err != nil {
		t.Fatalf("failed to read template: %v", err)
	}

	if doc.Find(`[data-testid="header-component"]`).Length() == 0 {
		t.Error("expected data-testid attribute to be rendered, but it wasn't")
	}

	component := doc.Find(`[data-testid="main-component"]`)
	if component.Length() == 0 {
		t.Error("expected data-testid attribute to be rendered, but it wasn't")
	}

	expected := "index page content"
	if actual := component.Find("h2").Text(); actual != expected {
		t.Errorf("expected %q, got %q", expected, actual)
	}

	if doc.Find(`[data-testid="footer-component"]`).Length() == 0 {
		t.Error("expected data-testid attribute to be rendered, but it wasn't")
	}
}

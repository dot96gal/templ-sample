package components

import (
	"context"
	"io"
	"testing"

	"github.com/PuerkitoBio/goquery"
)

func TestFooter(t *testing.T) {
	r, w := io.Pipe()

	go func() {
		_ = Footer("2024 dot96gal All Rights Reserved.").Render(context.Background(), w)
		_ = w.Close()
	}()

	doc, err := goquery.NewDocumentFromReader(r)
	if err != nil {
		t.Fatalf("failed to read template: %v", err)
	}

	component := doc.Find(`[data-testid="footer-component"]`)
	if component.Length() == 0 {
		t.Error("expected data-testid attribute to be rendered, but it wasn't")
	}

	expected := "© 2024 dot96gal All Rights Reserved."
	if actual := component.Text(); actual != expected {
		t.Errorf("expected %q, got %q", expected, actual)
	}
}

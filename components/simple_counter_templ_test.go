package components

import (
	"context"
	"io"
	"strings"
	"testing"

	"github.com/PuerkitoBio/goquery"
)

func TestSimpleCounter(t *testing.T) {
	r, w := io.Pipe()

	go func() {
		_ = SimpleCounter(0, "/").Render(context.Background(), w)
		_ = w.Close()
	}()

	doc, err := goquery.NewDocumentFromReader(r)
	if err != nil {
		t.Fatalf("failed to read template: %v", err)
	}

	component := doc.Find(`[data-testid="simple-counter-component"]`)
	if component.Length() == 0 {
		t.Error("expected data-testid attribute to be rendered, but it wasn't")
	}

	expected := ""
	if actual := component.Text(); !strings.Contains(actual, expected) {
		t.Errorf("expected %q, got %q", expected, actual)
	}
}

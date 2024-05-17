package components

import (
	"context"
	"fmt"
	"io"
	"testing"

	"github.com/PuerkitoBio/goquery"
	"github.com/a-h/templ"
)

func TestSimpleCounter(t *testing.T) {
	dataTestid := "simple-counter"

	r, w := io.Pipe()

	go func() {
		props := SimpleCounterProps{
			Count:      0,
			Endpoint:   "/",
			Attributes: templ.Attributes{"data-testid": dataTestid},
		}
		_ = SimpleCounter(props).Render(context.Background(), w)
		_ = w.Close()
	}()

	doc, err := goquery.NewDocumentFromReader(r)
	if err != nil {
		t.Fatalf("failed to read template: %v", err)
	}

	component := doc.Find(fmt.Sprintf(`[data-testid="%s"]`, dataTestid))
	if component.Length() == 0 {
		t.Error("expected data-testid attribute to be rendered, but it wasn't")
	}

	expected := "0"
	if actual := doc.Find(`[class^="simpleCounterCountIndicatorStyle"]`).Text(); actual != expected {
		t.Errorf("expected %q, got %q", expected, actual)
	}
}

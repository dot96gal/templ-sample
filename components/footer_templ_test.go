package components

import (
	"context"
	"fmt"
	"io"
	"testing"

	"github.com/PuerkitoBio/goquery"
	"github.com/a-h/templ"
)

func TestFooter(t *testing.T) {
	dataTestid := "footer"

	r, w := io.Pipe()

	go func() {
		props := FooterProps{
			Copyright:  "2024 dot96gal All Rights Reserved.",
			Attributes: templ.Attributes{"data-testid": dataTestid},
		}
		_ = Footer(props).Render(context.Background(), w)
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

	expected := "© 2024 dot96gal All Rights Reserved."
	if actual := component.Text(); actual != expected {
		t.Errorf("expected %q, got %q", expected, actual)
	}
}

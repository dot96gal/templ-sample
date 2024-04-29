package components

import (
	"context"
	"io"
	"strings"
	"testing"

	"github.com/PuerkitoBio/goquery"
	"github.com/a-h/templ"
)

func TestIconLink(t *testing.T) {
	r, w := io.Pipe()

	go func() {
		props := IconLinkProps{
			Icon: IconHome(),
			Text: "home",
			Link: "/",
			CSSClass: []templ.KeyValue[templ.CSSClass, bool]{
				templ.KV(sideMenuItemSelectedStyle(), true),
			},
			HoverStyleClass: sideMenuItemHoverStyleClass,
			HoverStyle:      sideMenuItemHoverStyle(sideMenuItemHoverStyleClass),
		}
		_ = IconLink(props).Render(context.Background(), w)
		_ = w.Close()
	}()

	doc, err := goquery.NewDocumentFromReader(r)
	if err != nil {
		t.Fatalf("failed to read template: %v", err)
	}

	component := doc.Find(`[data-testid="icon-link-component"]`)
	if component.Length() == 0 {
		t.Error("expected data-testid attribute to be rendered, but it wasn't")
	}

	expected := "home"
	if actual := component.Text(); !strings.Contains(actual, expected) {
		t.Errorf("expected %q, got %q", expected, actual)
	}
}

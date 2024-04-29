package components

import (
	"context"
	"fmt"
	"io"
	"strings"
	"testing"

	"github.com/PuerkitoBio/goquery"
	"github.com/a-h/templ"
)

func TestIconLink(t *testing.T) {
	dataTestid := "icon-link"

	r, w := io.Pipe()

	go func() {
		props := IconLinkProps{
			Icon:            IconHome(),
			Text:            "home",
			Link:            "/",
			SelectedStyle:   templ.KV(sideMenuItemSelectedStyle(), true),
			HoverStyleClass: sideMenuItemHoverStyleClass,
			HoverStyle:      sideMenuItemHoverStyle(sideMenuItemHoverStyleClass),
			Attributes:      templ.Attributes{"data-testid": dataTestid},
		}
		_ = IconLink(props).Render(context.Background(), w)
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

	expected := "home"
	if actual := component.Text(); !strings.Contains(actual, expected) {
		t.Errorf("expected %q, got %q", expected, actual)
	}
}

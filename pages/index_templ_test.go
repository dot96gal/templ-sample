package pages

import (
	"context"
	"fmt"
	"io"
	"testing"

	"github.com/PuerkitoBio/goquery"
	"github.com/a-h/templ"
)

func TestHeader(t *testing.T) {
	dataTestidHeader := "index-page-header"
	dataTestidMain := "index-page-main"
	dataTestidSideMenu := "index-page-side-menu"
	dataTestidFooter := "index-page-footer"

	r, w := io.Pipe()

	go func() {
		props := IndexPageProps{
			Count:          0,
			HandlePostPath: "/",
			Attributes: templ.Attributes{
				"data-testid-header":    dataTestidHeader,
				"data-testid-side-menu": dataTestidSideMenu,
				"data-testid-main":      dataTestidMain,
				"data-testid-footer":    dataTestidFooter,
			},
		}
		_ = IndexPage(props).Render(context.Background(), w)
		_ = w.Close()
	}()

	doc, err := goquery.NewDocumentFromReader(r)
	if err != nil {
		t.Fatalf("failed to read template: %v", err)
	}

	if doc.Find(fmt.Sprintf(`[data-testid="%s"]`, dataTestidHeader)).Length() == 0 {
		t.Error("expected data-testid attribute to be rendered, but it wasn't")
	}

	if doc.Find(fmt.Sprintf(`[data-testid="%s"]`, dataTestidSideMenu)).Length() == 0 {
		t.Error("expected data-testid attribute to be rendered, but it wasn't")
	}

	if doc.Find(fmt.Sprintf(`[data-testid="%s"]`, dataTestidMain)).Length() == 0 {
		t.Error("expected data-testid attribute to be rendered, but it wasn't")
	}

	if doc.Find(fmt.Sprintf(`[data-testid="%s"]`, dataTestidFooter)).Length() == 0 {
		t.Error("expected data-testid attribute to be rendered, but it wasn't")
	}
}

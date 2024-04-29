package components

import (
	"context"
	"io"
	"strings"
	"testing"

	"github.com/PuerkitoBio/goquery"
	"github.com/a-h/templ"
)

func TestSideMenu(t *testing.T) {
	r, w := io.Pipe()

	go func() {
		props := SideMenuProps{
			Items: []templ.Component{
				SideMenuItemHome(SideMenuItemHomeProps{Link: "/", IsSelected: true}),
				SideMenuItemChart(SideMenuItemChartProps{Link: "/", IsSelected: false}),
				SideMenuItemTrend(SideMenuItemTrendProps{Link: "/", IsSelected: false}),
			},
		}
		_ = SideMenu(props).Render(context.Background(), w)
		_ = w.Close()
	}()

	doc, err := goquery.NewDocumentFromReader(r)
	if err != nil {
		t.Fatalf("failed to read template: %v", err)
	}

	component := doc.Find(`[data-testid="side-menu-component"]`)
	if component.Length() == 0 {
		t.Error("expected data-testid attribute to be rendered, but it wasn't")
	}

	expected := "home"
	if actual := component.Text(); !strings.Contains(actual, expected) {
		t.Errorf("expected %q, got %q", expected, actual)
	}
}

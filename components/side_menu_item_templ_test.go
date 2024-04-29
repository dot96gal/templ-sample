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

func TestSideMenuItemHome(t *testing.T) {
	dataTestid := "side-menu-item-home"

	r, w := io.Pipe()

	go func() {
		props := SideMenuItemHomeProps{
			Link:       "/",
			IsSelected: true,
			Attributes: templ.Attributes{"data-testid": dataTestid},
		}
		_ = SideMenuItemHome(props).Render(context.Background(), w)
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

func TestSideMenuItemChart(t *testing.T) {
	dataTestid := "side-menu-item-chart"

	r, w := io.Pipe()

	go func() {
		props := SideMenuItemChartProps{
			Link:       "/",
			IsSelected: true,
			Attributes: templ.Attributes{"data-testid": dataTestid},
		}
		_ = SideMenuItemChart(props).Render(context.Background(), w)
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

	expected := "chart"
	if actual := component.Text(); !strings.Contains(actual, expected) {
		t.Errorf("expected %q, got %q", expected, actual)
	}
}

func TestSideMenuItemTrend(t *testing.T) {
	dataTestid := "side-menu-item-trend"

	r, w := io.Pipe()

	go func() {
		props := SideMenuItemTrendProps{
			Link:       "/",
			IsSelected: true,
			Attributes: templ.Attributes{"data-testid": dataTestid},
		}
		_ = SideMenuItemTrend(props).Render(context.Background(), w)
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

	expected := "trend"
	if actual := component.Text(); !strings.Contains(actual, expected) {
		t.Errorf("expected %q, got %q", expected, actual)
	}
}

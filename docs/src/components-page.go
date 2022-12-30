package main

import (
	"github.com/maxence-charriere/go-app/v9/pkg/analytics"
	"github.com/maxence-charriere/go-app/v9/pkg/app"
)

type ComponentsPage struct {
	app.Compo
}

func NewComponentsPage() *ComponentsPage {
	return &ComponentsPage{}
}

func (p *ComponentsPage) OnPreRender(ctx app.Context) {
	p.initPage(ctx)
}

func (p *ComponentsPage) OnNav(ctx app.Context) {
	p.initPage(ctx)
}

func (p *ComponentsPage) initPage(ctx app.Context) {
	ctx.Page().SetTitle("Building Components: Customizable, Independent, and Reusable UI Elements")
	ctx.Page().SetDescription("Documentation about building customizable, independent, and reusable UI elements.")
	analytics.Page("components", nil)
}

func (p *ComponentsPage) Render() app.UI {
	return NewPage().
		Title("Components").
		Icon(gridSVG).
		Index(
			NewIndexLink().Title("What is a Component?"),
			NewIndexLink().Title("Create"),
			NewIndexLink().Title("Customize Look"),
			NewIndexLink().Title("Fields"),
			NewIndexLink().Title("    Exported vs Unexported"),
			NewIndexLink().Title("    How chose between Exported and Unexported?"),
			NewIndexLink().Title("Lifecycle Events"),
			NewIndexLink().Title("    PreRender"),
			NewIndexLink().Title("    Mount"),
			NewIndexLink().Title("    Nav"),
			NewIndexLink().Title("    Dismount"),
			NewIndexLink().Title("    Reference"),
			NewIndexLink().Title("Updates"),
			NewIndexLink().Title("    Manually Trigger an Update"),

			app.Div().Class("separator"),

			NewIndexLink().Title("Next"),
		).
		Content(
			NewRemoteMarkdownDoc().Src("/web/documents/components.md"),
		)
}

package main

import (
	"github.com/maxence-charriere/go-app/v9/pkg/analytics"
	"github.com/maxence-charriere/go-app/v9/pkg/app"
)

type ArchitecturePage struct {
	app.Compo
}

func NewArchitecturePage() *ArchitecturePage {
	return &ArchitecturePage{}
}

func (p *ArchitecturePage) OnPreRender(ctx app.Context) {
	p.initPage(ctx)
}

func (p *ArchitecturePage) OnNav(ctx app.Context) {
	p.initPage(ctx)
}

func (p *ArchitecturePage) initPage(ctx app.Context) {
	ctx.Page().SetTitle("Understanding go-app Architecture")
	ctx.Page().SetDescription("Documentation about how go-app parts are working together to form a Progressive Web App (PWA).")
	analytics.Page("architecture", nil)
}

func (p *ArchitecturePage) Render() app.UI {
	return NewPage().
		Title("Architecture").
		Icon(fileTreeSVG).
		Index(
			NewIndexLink().Title("Overview"),
			NewIndexLink().Title("Web Browser"),
			NewIndexLink().Title("Server"),
			NewIndexLink().Title("HTML Pages"),
			NewIndexLink().Title("Package Resources"),
			NewIndexLink().Title("app.wasm"),
			NewIndexLink().Title("Static Resources"),

			app.Div().Class("separator"),

			NewIndexLink().Title("Next"),
		).
		Content(
			NewRemoteMarkdownDoc().Src("/web/documents/architecture.md"),
		)
}

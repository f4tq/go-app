package main

import (
	"github.com/maxence-charriere/go-app/v9/pkg/analytics"
	"github.com/maxence-charriere/go-app/v9/pkg/app"
)

type LifecyclePage struct {
	app.Compo
}

func NewLifecyclePage() *LifecyclePage {
	return &LifecyclePage{}
}

func (p *LifecyclePage) OnPreRender(ctx app.Context) {
	p.initPage(ctx)
}

func (p *LifecyclePage) OnNav(ctx app.Context) {
	p.initPage(ctx)
}

func (p *LifecyclePage) initPage(ctx app.Context) {
	ctx.Page().SetTitle("App Lifecycle and Updates")
	ctx.Page().SetDescription("Documentation that describes how a web browser installs and updates a go-app Progressive Web App.")
	analytics.Page("lifecycle", nil)
}

func (p *LifecyclePage) Render() app.UI {
	return NewPage().
		Title("Lifecycle and Updates").
		Icon(arrowSVG).
		Index(
			NewIndexLink().Title("Lifecycle Overview"),
			NewIndexLink().Title("    First loading"),
			NewIndexLink().Title("    Recurrent loadings"),
			NewIndexLink().Title("    Loading after an app update"),
			NewIndexLink().Title("Listen for App Updates"),

			app.Div().Class("separator"),

			NewIndexLink().Title("Next"),
		).
		Content(
			NewRemoteMarkdownDoc().Src("/web/documents/lifecycle.md"),
		)
}

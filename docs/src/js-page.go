package main

import (
	"github.com/maxence-charriere/go-app/v9/pkg/analytics"
	"github.com/maxence-charriere/go-app/v9/pkg/app"
)

type JsPage struct {
	app.Compo
}

func NewJSPage() *JsPage {
	return &JsPage{}
}

func (p *JsPage) OnPreRender(ctx app.Context) {
	p.initPage(ctx)
}

func (p *JsPage) OnNav(ctx app.Context) {
	p.initPage(ctx)
}

func (p *JsPage) initPage(ctx app.Context) {
	ctx.Page().SetTitle("JavaScript Interoperability")
	ctx.Page().SetDescription("Documentation about how to call JavaScript from Go or Go from JavaScript.")
	analytics.Page("js", nil)
}

func (p *JsPage) Render() app.UI {
	return NewPage().
		Title("JavaScript Interoperability").
		Icon(jsSVG).
		Index(
			NewIndexLink().Title("Intro"),
			NewIndexLink().Title("Include JS files"),
			NewIndexLink().Title("    Page's scope"),
			NewIndexLink().Title("    Inlined in Components"),
			NewIndexLink().Title("Using window global object"),
			NewIndexLink().Title("    Get element by ID"),
			NewIndexLink().Title("    Create JS object"),
			NewIndexLink().Title("Cancel an event"),
			NewIndexLink().Title("Get input value"),

			app.Div().Class("separator"),

			NewIndexLink().Title("Next"),
		).
		Content(
			NewRemoteMarkdownDoc().Src("/web/documents/js.md"),
		)
}

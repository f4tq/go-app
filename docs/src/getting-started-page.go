package main

import (
	"github.com/maxence-charriere/go-app/v9/pkg/analytics"
	"github.com/maxence-charriere/go-app/v9/pkg/app"
)

type GettingStartedPage struct {
	app.Compo
}

func NewGettingStartedPage() *GettingStartedPage {
	return &GettingStartedPage{}
}

func (p *GettingStartedPage) OnPreRender(ctx app.Context) {
	p.initPage(ctx)
}

func (p *GettingStartedPage) OnNav(ctx app.Context) {
	p.initPage(ctx)
}

func (p *GettingStartedPage) initPage(ctx app.Context) {
	ctx.Page().SetTitle("Start building a PWA with Go and WASM")
	ctx.Page().SetDescription("Documentation that shows how to start building a Progressive Web App (PWA) with Go (Golang) and WebAssembly (WASM).")
	analytics.Page("getting-started", nil)
}

func (p *GettingStartedPage) Render() app.UI {
	return NewPage().
		Title("Getting Started").
		Icon(rocketSVG).
		Index(
			NewIndexLink().Title("Intro"),
			NewIndexLink().Title("Prerequisite"),
			NewIndexLink().Title("Install"),
			NewIndexLink().Title("Code"),
			NewIndexLink().Title("    Hello component"),
			NewIndexLink().Title("    Main"),
			NewIndexLink().Title("Build and Run"),
			NewIndexLink().Title("    Build the Client"),
			NewIndexLink().Title("    Build the Server"),
			NewIndexLink().Title("    Run the App"),
			NewIndexLink().Title("    Use a Makefile"),

			app.Div().Class("separator"),

			NewIndexLink().Title("Next"),
		).
		Content(
			NewRemoteMarkdownDoc().Src("/web/documents/getting-started.md"),
		)
}

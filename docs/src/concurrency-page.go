package main

import (
	"github.com/maxence-charriere/go-app/v9/pkg/analytics"
	"github.com/maxence-charriere/go-app/v9/pkg/app"
)

type ConcurrencyPage struct {
	app.Compo
}

func NewConcurrencyPage() *ConcurrencyPage {
	return &ConcurrencyPage{}
}

func (p *ConcurrencyPage) OnPreRender(ctx app.Context) {
	p.initPage(ctx)
}

func (p *ConcurrencyPage) OnNav(ctx app.Context) {
	p.initPage(ctx)
}

func (p *ConcurrencyPage) initPage(ctx app.Context) {
	ctx.Page().SetTitle("Building a Concurrency-Safe PWA")
	ctx.Page().SetDescription("Documentation about how to build a concurrency-safe and reactive progressive web app (PWA).")
	analytics.Page("concurrency", nil)
}

func (p *ConcurrencyPage) Render() app.UI {
	return NewPage().
		Title("Concurrency").
		Icon(concurrencySVG).
		Index(
			NewIndexLink().Title("Intro"),
			NewIndexLink().Title("UI Goroutine"),
			NewIndexLink().Title("Async"),
			NewIndexLink().Title("Dispatch"),
			NewIndexLink().Title("Defer"),

			app.Div().Class("separator"),

			NewIndexLink().Title("Next"),
		).
		Content(
			NewRemoteMarkdownDoc().Src("/web/documents/concurrency.md"),
		)
}

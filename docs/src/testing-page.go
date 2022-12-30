package main

import (
	"github.com/maxence-charriere/go-app/v9/pkg/analytics"
	"github.com/maxence-charriere/go-app/v9/pkg/app"
)

type TestingPage struct {
	app.Compo
}

func NewTestingPage() *TestingPage {
	return &TestingPage{}
}

func (p *TestingPage) OnPreRender(ctx app.Context) {
	p.initPage(ctx)
}

func (p *TestingPage) OnNav(ctx app.Context) {
	p.initPage(ctx)
}

func (p *TestingPage) initPage(ctx app.Context) {
	ctx.Page().SetTitle("Testing Components")
	ctx.Page().SetDescription("Documentation about how to unit test components created with go-app.")
	analytics.Page("testing", nil)
}

func (p *TestingPage) Render() app.UI {
	return NewPage().
		Title("Testing").
		Icon(testSVG).
		Index(
			NewIndexLink().Title("Intro"),
			NewIndexLink().Title("Component server prerendering"),
			NewIndexLink().Title("Component client lifecycle"),
			NewIndexLink().Title("Asynchronous operations"),
			NewIndexLink().Title("UI elements"),

			app.Div().Class("separator"),

			NewIndexLink().Title("Next"),
		).
		Content(
			NewRemoteMarkdownDoc().Src("/web/documents/testing.md"),
		)
}

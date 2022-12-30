package main

import (
	"github.com/maxence-charriere/go-app/v9/pkg/analytics"
	"github.com/maxence-charriere/go-app/v9/pkg/app"
)

type StatesPage struct {
	app.Compo
}

func NewStatesPage() *StatesPage {
	return &StatesPage{}
}

func (p *StatesPage) OnPreRender(ctx app.Context) {
	p.initPage(ctx)
}

func (p *StatesPage) OnNav(ctx app.Context) {
	p.initPage(ctx)
}

func (p *StatesPage) initPage(ctx app.Context) {
	ctx.Page().SetTitle("State Management")
	ctx.Page().SetDescription("Documentation about how to set and observe states.")
	analytics.Page("states", nil)
}

func (p *StatesPage) Render() app.UI {
	return NewPage().
		Title("State Management").
		Icon(stateSVG).
		Index(
			NewIndexLink().Title("What is a state?"),
			NewIndexLink().Title("Set"),
			NewIndexLink().Title("    Options"),
			NewIndexLink().Title("Observe"),
			NewIndexLink().Title("    Conditional Observation"),
			NewIndexLink().Title("    Additional Instructions"),
			NewIndexLink().Title("Get"),

			app.Div().Class("separator"),

			NewIndexLink().Title("Next"),
		).
		Content(
			NewRemoteMarkdownDoc().Src("/web/documents/states.md"),
		)
}

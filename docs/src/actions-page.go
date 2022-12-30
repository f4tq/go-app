package main

import (
	"github.com/maxence-charriere/go-app/v9/pkg/analytics"
	"github.com/maxence-charriere/go-app/v9/pkg/app"
)

type ActionPage struct {
	app.Compo
}

func NewActionPage() *ActionPage {
	return &ActionPage{}
}

func (p *ActionPage) OnPreRender(ctx app.Context) {
	p.initPage(ctx)
}

func (p *ActionPage) OnNav(ctx app.Context) {
	p.initPage(ctx)
}

func (p *ActionPage) initPage(ctx app.Context) {
	ctx.Page().SetTitle("Creating and Listening to Actions")
	ctx.Page().SetDescription("Documentation about how to create and listen to actions.")
	analytics.Page("actions", nil)
}

func (p *ActionPage) Render() app.UI {
	return NewPage().
		Title("Actions").
		Icon(actionSVG).
		Index(
			NewIndexLink().Title("What is an Action?"),
			NewIndexLink().Title("Create"),
			NewIndexLink().Title("Handling"),
			NewIndexLink().Title("    Global Level"),
			NewIndexLink().Title("    Component Level"),

			app.Div().Class("separator"),

			NewIndexLink().Title("Next"),
		).
		Content(
			NewRemoteMarkdownDoc().Src("/web/documents/actions.md"),
		)
}

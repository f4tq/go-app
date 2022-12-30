package main

import (
	"github.com/maxence-charriere/go-app/v9/pkg/analytics"
	"github.com/maxence-charriere/go-app/v9/pkg/app"
)

type RoutingPage struct {
	app.Compo
}

func NewRoutingPage() *RoutingPage {
	return &RoutingPage{}
}

func (p *RoutingPage) OnPreRender(ctx app.Context) {
	p.initPage(ctx)
}

func (p *RoutingPage) OnNav(ctx app.Context) {
	p.initPage(ctx)
}

func (p *RoutingPage) initPage(ctx app.Context) {
	ctx.Page().SetTitle("Routing URL paths to Components")
	ctx.Page().SetDescription("Documentation about how to associate URL paths to go-app components.")
	analytics.Page("routing", nil)
}

func (p *RoutingPage) Render() app.UI {
	return NewPage().
		Title("Routing").
		Icon(routeSVG).
		Index(
			NewIndexLink().Title("Intro"),
			NewIndexLink().Title("Define a route"),
			NewIndexLink().Title("    Simple route"),
			NewIndexLink().Title("    Route with regular expression"),
			NewIndexLink().Title("How it works?"),
			NewIndexLink().Title("Detect navigation"),

			app.Div().Class("separator"),

			NewIndexLink().Title("Next"),
		).
		Content(
			NewRemoteMarkdownDoc().Src("/web/documents/routing.md"),
		)
}

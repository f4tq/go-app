package main

import (
	"github.com/maxence-charriere/go-app/v9/pkg/analytics"
	"github.com/maxence-charriere/go-app/v9/pkg/app"
	"github.com/maxence-charriere/go-app/v9/pkg/ui"
)

type HomePage struct {
	app.Compo
}

func NewHomePage() *HomePage {
	return &HomePage{}
}

func (p *HomePage) OnPreRender(ctx app.Context) {
	p.initPage(ctx)
}

func (p *HomePage) OnNav(ctx app.Context) {
	p.initPage(ctx)
}

func (p *HomePage) initPage(ctx app.Context) {
	ctx.Page().SetTitle(defaultTitle)
	ctx.Page().SetDescription(defaultDescription)
	analytics.Page("home", nil)
}

func (p *HomePage) Render() app.UI {
	return NewPage().
		Title("go-app").
		Icon("https://storage.googleapis.com/murlok-github/icon-192.png").
		Index(
			NewIndexLink().Title("What is go-app?"),
			NewIndexLink().Title("Updates"),
			NewIndexLink().Title("Declarative Syntax"),
			NewIndexLink().Title("Standard HTTP Server"),
			NewIndexLink().Title("Other features"),
			NewIndexLink().Title("Built With go-app"),

			app.Div().Class("separator"),

			NewIndexLink().Title("Next"),
		).
		Content(
			ui.Flow().
				StretchItems().
				Spacing(84).
				Content(
					NewRemoteMarkdownDoc().
						Class("fill").
						Src("/web/documents/what-is-go-app.md"),
					NewRemoteMarkdownDoc().
						Class("fill").
						Class("updates").
						Src("/web/documents/updates.md"),
				),

			app.Div().Class("separator"),

			NewRemoteMarkdownDoc().Src("/web/documents/home.md"),

			app.Div().Class("separator"),

			NewBuiltWithGoapp().ID("built-with-go-app"),

			app.Div().Class("separator"),

			NewRemoteMarkdownDoc().Src("/web/documents/home-next.md"),
		)
}

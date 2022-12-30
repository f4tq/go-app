package main

import (
	"github.com/maxence-charriere/go-app/v9/pkg/analytics"
	"github.com/maxence-charriere/go-app/v9/pkg/app"
)

type MigratePage struct {
	app.Compo
}

func NewMigratePage() *MigratePage {
	return &MigratePage{}
}

func (p *MigratePage) OnPreRender(ctx app.Context) {
	p.initPage(ctx)
}

func (p *MigratePage) OnNav(ctx app.Context) {
	p.initPage(ctx)
}

func (p *MigratePage) initPage(ctx app.Context) {
	ctx.Page().SetTitle("Migrate Codebase From go-app v8 To v9")
	ctx.Page().SetDescription("Documentation about what changed between go-app v8 and v9.")
	analytics.Page("migrate", nil)
}

func (p *MigratePage) Render() app.UI {
	return NewPage().
		Title("Migrate From v8 to v9").
		Icon(swapSVG).
		Index(
			NewIndexLink().Title("Intro"),
			NewIndexLink().Title("Changes"),
			NewIndexLink().Title("    General"),
			NewIndexLink().Title("    Components"),
			NewIndexLink().Title("    Context"),
			NewIndexLink().Title("API Design Decisions"),

			app.Div().Class("separator"),
		).
		Content(
			NewRemoteMarkdownDoc().Src("/web/documents/migrate.md"),
		)
}

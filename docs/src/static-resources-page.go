package main

import (
	"github.com/maxence-charriere/go-app/v9/pkg/analytics"
	"github.com/maxence-charriere/go-app/v9/pkg/app"
)

type StaticResourcesPage struct {
	app.Compo
}

func NewStaticResourcePage() *StaticResourcesPage {
	return &StaticResourcesPage{}
}

func (p *StaticResourcesPage) OnPreRender(ctx app.Context) {
	p.initPage(ctx)
}

func (p *StaticResourcesPage) OnNav(ctx app.Context) {
	p.initPage(ctx)
}

func (p *StaticResourcesPage) initPage(ctx app.Context) {
	ctx.Page().SetTitle("Images and Static Resources")
	ctx.Page().SetDescription("Documentation about how to deal with images and other static resources.")
	analytics.Page("static-resources", nil)
}

func (p *StaticResourcesPage) Render() app.UI {
	return NewPage().
		Title("Images and Static Resources").
		Icon(imgFolderSVG).
		Index(
			NewIndexLink().Title("Intro"),
			NewIndexLink().Title("Access static resources"),
			NewIndexLink().Title("    In Handler"),
			NewIndexLink().Title("    In components"),
			NewIndexLink().Title("Setup Custom Web directory"),
			NewIndexLink().Title("    Setup local web directory"),
			NewIndexLink().Title("    Setup remote web directory"),

			app.Div().Class("separator"),

			NewIndexLink().Title("Next"),
		).
		Content(
			NewRemoteMarkdownDoc().Src("/web/documents/static-resources.md"),
		)
}

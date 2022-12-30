package main

import (
	"github.com/maxence-charriere/go-app/v9/pkg/analytics"
	"github.com/maxence-charriere/go-app/v9/pkg/app"
)

type InstallPage struct {
	app.Compo
}

func NewInstallPage() *InstallPage {
	return &InstallPage{}
}

func (p *InstallPage) OnPreRender(ctx app.Context) {
	p.initPage(ctx)
}

func (p *InstallPage) OnNav(ctx app.Context) {
	p.initPage(ctx)
}

func (p *InstallPage) initPage(ctx app.Context) {
	ctx.Page().SetTitle("Handling App Install")
	ctx.Page().SetDescription("Documentation about how to install an app created with go-app.")
	analytics.Page("install", nil)
}

func (p *InstallPage) Render() app.UI {
	return NewPage().
		Title("Install").
		Icon(downloadSVG).
		Index(
			NewIndexLink().Title("Intro"),
			NewIndexLink().Title("Desktop"),
			NewIndexLink().Title("IOS"),
			NewIndexLink().Title("Android"),
			NewIndexLink().Title("Programmatically"),
			NewIndexLink().Title("    Detect Install Support"),
			NewIndexLink().Title("    Display Install Popup"),

			app.Div().Class("separator"),

			NewIndexLink().Title("Next"),
		).
		Content(
			NewRemoteMarkdownDoc().Src("/web/documents/install.md"),
		)
}

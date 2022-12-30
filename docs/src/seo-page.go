package main

import (
	"github.com/maxence-charriere/go-app/v9/pkg/analytics"
	"github.com/maxence-charriere/go-app/v9/pkg/app"
)

type SeoPage struct {
	app.Compo
}

func newSEOPage() *SeoPage {
	return &SeoPage{}
}

func (p *SeoPage) OnPreRender(ctx app.Context) {
	p.initPage(ctx)
}

func (p *SeoPage) OnNav(ctx app.Context) {
	p.initPage(ctx)
}

func (p *SeoPage) initPage(ctx app.Context) {
	ctx.Page().SetTitle("Building SEO-friendly PWA")
	ctx.Page().SetDescription("Documentation about how to make a Progressive Web App indexable by search engines with go-app package.")
	analytics.Page("seo", nil)
}

func (p *SeoPage) Render() app.UI {
	return NewPage().
		Title("SEO").
		Icon(seoSVG).
		Index(
			NewIndexLink().Title("Intro"),
			NewIndexLink().Title("Prerendering"),
			NewIndexLink().Title("    Customizing prerendering"),
			NewIndexLink().Title("    Customizing page metadata"),
			NewIndexLink().Title("    Caching"),

			app.Div().Class("separator"),

			NewIndexLink().Title("Next"),
		).
		Content(
			NewRemoteMarkdownDoc().Src("/web/documents/seo.md"),
		)
}

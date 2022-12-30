package main

import (
	"github.com/maxence-charriere/go-app/v9/pkg/analytics"
	"github.com/maxence-charriere/go-app/v9/pkg/app"
)

type GithubDeployPage struct {
	app.Compo
}

func NewGithubDeployPage() *GithubDeployPage {
	return &GithubDeployPage{}
}

func (p *GithubDeployPage) OnPreRender(ctx app.Context) {
	p.initPage(ctx)
}

func (p *GithubDeployPage) OnNav(ctx app.Context) {
	p.initPage(ctx)
}

func (p *GithubDeployPage) initPage(ctx app.Context) {
	ctx.Page().SetTitle("Deploy a PWA on GitHub Pages")
	ctx.Page().SetDescription("Documentation about how to deploy a PWA created with go-app on GitHub Pages.")
	analytics.Page("github-deploy", nil)
}

func (p *GithubDeployPage) Render() app.UI {
	return NewPage().
		Title("Deploy on GitHub Pages").
		Icon(githubSVG).
		Index(
			NewIndexLink().Title("Intro"),
			NewIndexLink().Title("Generate a Static Website"),
			NewIndexLink().Title("Deployment"),
			NewIndexLink().Title("    Domainless Repository"),

			app.Div().Class("separator"),
		).
		Content(
			NewRemoteMarkdownDoc().Src("/web/documents/github-deploy.md"),
		)
}

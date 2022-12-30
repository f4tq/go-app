package main

import (
	"github.com/maxence-charriere/go-app/v9/pkg/app"
	"github.com/maxence-charriere/go-app/v9/pkg/ui"
)

type Menu struct {
	app.Compo

	Iclass string

	appInstallable bool
}

func NewMenu() *Menu {
	return &Menu{}
}

func (m *Menu) Class(v string) *Menu {
	m.Iclass = app.AppendClass(m.Iclass, v)
	return m
}

func (m *Menu) OnNav(ctx app.Context) {
	m.appInstallable = ctx.IsAppInstallable()
}

func (m *Menu) OnAppInstallChange(ctx app.Context) {
	m.appInstallable = ctx.IsAppInstallable()
}

func (m *Menu) Render() app.UI {
	linkClass := "link heading fit unselectable"

	isFocus := func(path string) string {
		if app.Window().URL().Path == path {
			return "focus"
		}
		return ""
	}

	return ui.Scroll().
		Class("menu").
		Class(m.Iclass).
		HeaderHeight(headerHeight).
		Header(
			ui.Stack().
				Class("fill").
				Middle().
				Content(
					app.Header().Body(
						app.A().
							Class("heading").
							Class("app-title").
							Href("/").
							Text("Go-App"),
					),
				),
		).
		Content(
			app.Nav().Body(
				app.Div().Class("separator"),

				ui.Link().
					Class(linkClass).
					Icon(homeSVG).
					Label("Home").
					Href("/").
					Class(isFocus("/")),
				ui.Link().
					Class(linkClass).
					Icon(rocketSVG).
					Label("Getting Started").
					Href("/getting-started").
					Class(isFocus("/getting-started")),
				ui.Link().
					Class(linkClass).
					Icon(fileTreeSVG).
					Label("Architecture").
					Href("/architecture").
					Class(isFocus("/architecture")),
				ui.Link().
					Class(linkClass).
					Icon(golangSVG).
					Label("Reference").
					Href("/reference").
					Class(isFocus("/reference")),

				app.Div().Class("separator"),

				ui.Link().
					Class(linkClass).
					Icon(gridSVG).
					Label("Components").
					Href("/components").
					Class(isFocus("/components")),
				ui.Link().
					Class(linkClass).
					Icon(keyboardSVG).
					Label("Declarative Syntax").
					Href("/declarative-syntax").
					Class(isFocus("/declarative-syntax")),
				ui.Link().
					Class(linkClass).
					Icon(routeSVG).
					Label("Routing").
					Href("/routing").
					Class(isFocus("/routing")),
				ui.Link().
					Class(linkClass).
					Icon(imgFolderSVG).
					Label("Images and Static Resources").
					Href("/static-resources").
					Class(isFocus("/static-resources")),
				ui.Link().
					Class(linkClass).
					Icon(jsSVG).
					Label("JavaScript Interoperability").
					Href("/js").
					Class(isFocus("/js")),
				ui.Link().
					Class(linkClass).
					Icon(concurrencySVG).
					Label("Concurrency").
					Href("/concurrency").
					Class(isFocus("/concurrency")),
				ui.Link().
					Class(linkClass).
					Icon(seoSVG).
					Label("SEO").
					Href("/seo").
					Class(isFocus("/seo")),
				ui.Link().
					Class(linkClass).
					Icon(arrowSVG).
					Label("Lifecycle and Updates").
					Href("/lifecycle").
					Class(isFocus("/lifecycle")),
				ui.Link().
					Class(linkClass).
					Icon(downloadSVG).
					Label("Install").
					Href("/install").
					Class(isFocus("/install")),
				ui.Link().
					Class(linkClass).
					Icon(testSVG).
					Label("Testing").
					Href("/testing").
					Class(isFocus("/testing")),
				ui.Link().
					Class(linkClass).
					Icon(actionSVG).
					Label("Actions").
					Href("/actions").
					Class(isFocus("/actions")),
				ui.Link().
					Class(linkClass).
					Icon(stateSVG).
					Label("State Management").
					Href("/states").
					Class(isFocus("/states")),
				ui.Link().
					Class(linkClass).
					Icon(bellSVG).
					Label("Notifications").
					Href("/notifications").
					Class(isFocus("/notifications")),

				app.Div().Class("separator"),

				ui.Link().
					Class(linkClass).
					Icon(swapSVG).
					Label("Migrate From v8 to v9").
					Href("/migrate").
					Class(isFocus("/migrate")),
				ui.Link().
					Class(linkClass).
					Icon(githubSVG).
					Label("Deploy on GitHub Pages").
					Href("/github-deploy").
					Class(isFocus("/github-deploy")),

				app.Div().Class("separator"),

				ui.Link().
					Class(linkClass).
					Icon(twitterSVG).
					Label("Twitter").
					Href(twitterURL),
				ui.Link().
					Class(linkClass).
					Icon(githubSVG).
					Label("GitHub").
					Href(githubURL),
				ui.Link().
					Class(linkClass).
					Icon(opensourceSVG).
					Label("Open Collective").
					Href(openCollectiveURL),

				app.Div().Class("separator"),

				app.If(m.appInstallable,
					ui.Link().
						Class(linkClass).
						Icon(downloadSVG).
						Label("Install").
						OnClick(m.installApp),
				),
				ui.Link().
					Class(linkClass).
					Icon(userLockSVG).
					Label("Privacy Policy").
					Href("/privacy-policy").
					Class(isFocus("/privacy-policy")),

				app.Div().Class("separator"),
			),
		)
}

func (m *Menu) installApp(ctx app.Context, e app.Event) {
	ctx.NewAction(installApp)
}

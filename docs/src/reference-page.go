package main

import (
	"github.com/maxence-charriere/go-app/v9/pkg/analytics"
	"github.com/maxence-charriere/go-app/v9/pkg/app"
)

type ReferencePage struct {
	app.Compo
}

func newReferencePage() *ReferencePage {
	return &ReferencePage{}
}

func (p *ReferencePage) OnPreRender(ctx app.Context) {
	p.initPage(ctx)
}

func (p *ReferencePage) OnNav(ctx app.Context) {
	p.initPage(ctx)
}

func (p *ReferencePage) initPage(ctx app.Context) {
	ctx.Page().SetTitle("Reference for building PWA with Go and WASM")
	ctx.Page().SetDescription("Go-app API reference for building Progressive Web Apps (PWA) with Go (Golang) and WebAssembly (WASM).")
	analytics.Page("reference", nil)
}

func (p *ReferencePage) Render() app.UI {
	return NewPage().
		Title("Reference").
		Icon(golangSVG).
		Index(
			app.A().
				Class("index-link").
				Class(fragmentFocus("pkg-overview")).
				Href("#pkg-overview").
				Text("Overview"),
			NewReferenceContent().
				Class("reference-index").
				Index(true),
			app.Div().Class("separator"),
		).
		Content(
			NewReferenceContent().Class("reference"),
		)
}

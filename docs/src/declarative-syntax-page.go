package main

import (
	"github.com/maxence-charriere/go-app/v9/pkg/analytics"
	"github.com/maxence-charriere/go-app/v9/pkg/app"
)

type DeclarativeSyntaxPage struct {
	app.Compo
}

func NewDeclarativeSyntaxPage() *DeclarativeSyntaxPage {
	return &DeclarativeSyntaxPage{}
}

func (p *DeclarativeSyntaxPage) OnPreRender(ctx app.Context) {
	p.initPage(ctx)
}

func (p *DeclarativeSyntaxPage) OnNav(ctx app.Context) {
	p.initPage(ctx)
}

func (p *DeclarativeSyntaxPage) initPage(ctx app.Context) {
	ctx.Page().SetTitle("Customize Components with go-app Declarative Syntax")
	ctx.Page().SetDescription("Documentation about how to customize components with go-app declarative syntax.")
	analytics.Page("declarative-syntax", nil)
}

func (p *DeclarativeSyntaxPage) Render() app.UI {
	return NewPage().
		Title("Declarative Syntax").
		Icon(keyboardSVG).
		Index(
			NewIndexLink().Title("Intro"),
			NewIndexLink().Title("HTML Elements"),
			NewIndexLink().Title("    Create"),
			NewIndexLink().Title("    Standard Elements"),
			NewIndexLink().Title("    Self Closing Elements"),
			NewIndexLink().Title("    Attributes"),
			NewIndexLink().Title("    Style"),
			NewIndexLink().Title("    Event handlers"),
			NewIndexLink().Title("Raw elements"),
			NewIndexLink().Title("Nested Components"),
			NewIndexLink().Title("Condition"),
			NewIndexLink().Title("    If"),
			NewIndexLink().Title("    ElseIf"),
			NewIndexLink().Title("    Else"),
			NewIndexLink().Title("Range"),
			NewIndexLink().Title("    Slice"),
			NewIndexLink().Title("    Map"),
			NewIndexLink().Title("Form helpers"),
			NewIndexLink().Title("    ValueTo"),

			app.Div().Class("separator"),

			NewIndexLink().Title("Next"),
		).
		Content(
			NewRemoteMarkdownDoc().Src("/web/documents/declarative-syntax.md"),
		)
}

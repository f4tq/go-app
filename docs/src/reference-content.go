package main

import (
	"github.com/maxence-charriere/go-app/v9/pkg/app"
	"github.com/maxence-charriere/go-app/v9/pkg/ui"
)

type ReferenceContent struct {
	app.Compo

	Iid    string
	Iclass string
	Iindex bool

	content         HtmlContent
	currentFragment string
}

func NewReferenceContent() *ReferenceContent {
	return &ReferenceContent{}
}

func (c *ReferenceContent) ID(v string) *ReferenceContent {
	c.Iid = v
	return c
}

func (c *ReferenceContent) Class(v string) *ReferenceContent {
	c.Iclass = app.AppendClass(c.Iclass, v)
	return c
}

func (c *ReferenceContent) Index(v bool) *ReferenceContent {
	c.Iindex = v
	return c
}

func (c *ReferenceContent) OnPreRender(ctx app.Context) {
	c.load(ctx)
}

func (c *ReferenceContent) OnMount(ctx app.Context) {
	c.load(ctx)
}

func (c *ReferenceContent) OnNav(ctx app.Context) {
	c.handleFragment(ctx)
}

func (c *ReferenceContent) load(ctx app.Context) {
	ctx.ObserveState(referenceState).
		OnChange(func() {
			ctx.Defer(c.handleFragment)
			ctx.Defer(c.scrollTo)
		}).
		Value(&c.content)

	ctx.NewAction(getReference)
}

func (c *ReferenceContent) Render() app.UI {
	loaderSize := 60
	loaderSpacing := 18
	if c.Iindex {
		loaderSize = 30
		loaderSpacing = 9
	}

	return app.Section().
		ID(c.Iid).
		Class(c.Iclass).
		Body(
			ui.Loader().
				Class("separator").
				Class("fill").
				Class("heading").
				Loading(c.content.Status == loading).
				Err(c.content.Err).
				Size(loaderSize).
				Spacing(loaderSpacing),

			app.If(!c.Iindex && c.content.Content != "",
				app.Raw(c.content.Content),
			).ElseIf(c.Iindex && c.content.Index != "",
				app.Raw(c.content.Index),
			),
			app.Div().Text(c.content.Err),
		)
}

func (c *ReferenceContent) handleFragment(ctx app.Context) {
	if !c.Iindex {
		return
	}
	if c.currentFragment != "" {
		c.unfocusCurrentIndex(ctx)
	}
	c.focusCurrentIndex(ctx)
}

func (c *ReferenceContent) unfocusCurrentIndex(ctx app.Context) {
	link := app.Window().GetElementByID(refLinkID(c.currentFragment))
	if !link.Truthy() {
		return
	}
	link.Set("className", "")
}

func (c *ReferenceContent) focusCurrentIndex(ctx app.Context) {
	fragment := ctx.Page().URL().Fragment
	link := app.Window().GetElementByID(refLinkID(fragment))
	if !link.Truthy() {
		return
	}
	link.Set("className", "focus")
	c.currentFragment = fragment
}

func (c *ReferenceContent) scrollTo(ctx app.Context) {
	id := ctx.Page().URL().Fragment
	if c.Iindex {
		id = refLinkID(id)
	}
	ctx.ScrollTo(id)
}

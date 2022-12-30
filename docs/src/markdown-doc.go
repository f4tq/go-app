package main

import (
	"fmt"
	"path/filepath"

	"github.com/gomarkdown/markdown"
	"github.com/gomarkdown/markdown/parser"
	"github.com/maxence-charriere/go-app/v9/pkg/app"
	"github.com/maxence-charriere/go-app/v9/pkg/ui"
)

type MarkdownDoc struct {
	app.Compo

	Iid    string
	Iclass string
	Imd    string
}

func NewMarkdownDoc() *MarkdownDoc {
	return &MarkdownDoc{}
}

func (d *MarkdownDoc) ID(v string) *MarkdownDoc {
	d.Iid = v
	return d
}

func (d *MarkdownDoc) Class(v string) *MarkdownDoc {
	d.Iclass = app.AppendClass(d.Iclass, v)
	return d
}

func (d *MarkdownDoc) MD(v string) *MarkdownDoc {
	d.Imd = fmt.Sprintf(`<div class="markdown">%s</div>`, parseMarkdown([]byte(v)))
	return d
}

func (d *MarkdownDoc) OnMount(ctx app.Context) {
	ctx.Defer(d.highlightCode)
}

func (d *MarkdownDoc) OnUpdate(ctx app.Context) {
	ctx.Defer(d.highlightCode)
}

func (d *MarkdownDoc) Render() app.UI {
	return app.Div().
		ID(d.Iid).
		Class(d.Iclass).
		Body(
			app.Raw(d.Imd),
		)
}

func (d *MarkdownDoc) highlightCode(ctx app.Context) {
	app.Window().Get("Prism").Call("highlightAll")
}

func parseMarkdown(md []byte) []byte {
	extensions := parser.CommonExtensions | parser.AutoHeadingIDs
	parser := parser.NewWithExtensions(extensions)
	return markdown.ToHTML(md, parser, nil)
}

type RemoteMarkdownDoc struct {
	app.Compo

	Iid    string
	Iclass string
	Isrc   string

	md markdownContent
}

func NewRemoteMarkdownDoc() *RemoteMarkdownDoc {
	return &RemoteMarkdownDoc{}
}

func (d *RemoteMarkdownDoc) ID(v string) *RemoteMarkdownDoc {
	d.Iid = v
	return d
}

func (d *RemoteMarkdownDoc) Class(v string) *RemoteMarkdownDoc {
	d.Iclass = app.AppendClass(d.Iclass, v)
	return d
}

func (d *RemoteMarkdownDoc) Src(v string) *RemoteMarkdownDoc {
	d.Isrc = v
	return d
}

func (d *RemoteMarkdownDoc) OnMount(ctx app.Context) {
	d.load(ctx)
}

func (d *RemoteMarkdownDoc) OnUpdate(ctx app.Context) {
	d.load(ctx)
}

func (d *RemoteMarkdownDoc) load(ctx app.Context) {
	src := d.Isrc
	ctx.ObserveState(markdownState(src)).
		While(func() bool {
			return src == d.Isrc
		}).
		OnChange(func() {
			ctx.Defer(scrollTo)
		}).
		Value(&d.md)

	ctx.NewAction(getMarkdown, app.T("path", d.Isrc))
}

func (d *RemoteMarkdownDoc) Render() app.UI {
	return app.Div().
		ID(d.Iid).
		Class(d.Iclass).
		Body(
			ui.Loader().
				Class("heading").
				Class("fill").
				Loading(d.md.Status == loading).
				Err(d.md.Err).
				Label(fmt.Sprintf("Loading %s...", filepath.Base(d.Isrc))),
			app.If(d.md.Status == loaded,
				NewMarkdownDoc().
					Class("fill").
					MD(d.md.Data),
			).Else(),
		)
}

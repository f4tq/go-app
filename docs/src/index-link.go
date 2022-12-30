package main

import (
	"strings"

	"github.com/maxence-charriere/go-app/v9/pkg/app"
)

type IndexLink struct {
	app.Compo

	Iclass string
	Ititle string
	Ihref  string
}

func NewIndexLink() *IndexLink {
	return &IndexLink{}
}

func (l *IndexLink) Class(v string) *IndexLink {
	l.Iclass = app.AppendClass(l.Iclass, v)
	return l
}

func (l *IndexLink) Title(v string) *IndexLink {
	l.Ititle = v
	return l
}

func (l *IndexLink) Href(v string) *IndexLink {
	l.Ihref = v
	return l
}

func (l *IndexLink) OnNav(ctx app.Context) {}

func (l *IndexLink) Render() app.UI {
	fragment := titleToFragment(l.Ititle)

	href := l.Ihref
	if href == "" {
		href = "#" + fragment
	}

	return app.A().
		Class("index-link").
		Class(l.Iclass).
		Class(fragmentFocus(fragment)).
		Href(href).
		Text(l.Ititle).
		Title(l.Ititle)
}

func titleToFragment(v string) string {
	v = strings.TrimSpace(v)
	v = strings.ToLower(v)
	v = strings.ReplaceAll(v, " ", "-")
	v = strings.ReplaceAll(v, ".", "-")
	v = strings.ReplaceAll(v, "?", "")
	return v
}

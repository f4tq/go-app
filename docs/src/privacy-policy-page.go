package main

import (
	"github.com/maxence-charriere/go-app/v9/pkg/analytics"
	"github.com/maxence-charriere/go-app/v9/pkg/app"
)

type PrivacyPolicyPage struct {
	app.Compo
}

func NewPrivacyPolicyPage() *PrivacyPolicyPage {
	return &PrivacyPolicyPage{}
}

func (p *PrivacyPolicyPage) OnPreRender(ctx app.Context) {
	p.initPage(ctx)
}

func (p *PrivacyPolicyPage) OnNav(ctx app.Context) {
	p.initPage(ctx)
}

func (p *PrivacyPolicyPage) initPage(ctx app.Context) {
	ctx.Page().SetTitle("Privacy Policy")
	ctx.Page().SetDescription("go-app documentation privacy policy.")
	analytics.Page("privacy-policy", nil)
}

func (p *PrivacyPolicyPage) Render() app.UI {
	return NewPage().
		Title("Privacy Policy").
		Icon(userLockSVG).
		Index(
			NewIndexLink().Title("Intro"),
			NewIndexLink().Title("Personal Data"),
			NewIndexLink().Title("Log Data"),
			NewIndexLink().Title("Cookies"),
			NewIndexLink().Title("Service Providers"),
			NewIndexLink().Title("Links to Other Sites"),
			NewIndexLink().Title("Changes to this Privacy Policy"),

			app.Div().Class("separator"),
		).
		Content(
			NewRemoteMarkdownDoc().Src("/web/documents/privacy-policy.md"),
		)
}

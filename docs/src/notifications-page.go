package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"math/rand"
	"net/http"

	"github.com/maxence-charriere/go-app/v9/pkg/analytics"
	"github.com/maxence-charriere/go-app/v9/pkg/app"
	"github.com/maxence-charriere/go-app/v9/pkg/errors"
)

type NotificationsPage struct {
	app.Compo

	notificationPermission app.NotificationPermission
}

func NewNotificationsPage() *NotificationsPage {
	return &NotificationsPage{}
}

func (p *NotificationsPage) OnPreRender(ctx app.Context) {
	p.initPage(ctx)
}

func (p *NotificationsPage) OnMount(ctx app.Context) {
	p.notificationPermission = ctx.Notifications().Permission()
	p.registerSubscription(ctx)
}

func (p *NotificationsPage) OnNav(ctx app.Context) {
	p.initPage(ctx)
}

func (p *NotificationsPage) initPage(ctx app.Context) {
	ctx.Page().SetTitle("Receive And Display Notifications")
	ctx.Page().SetDescription("Documentation about how receive and display notifications.")
	analytics.Page("notifications", nil)
}

func (p *NotificationsPage) Render() app.UI {
	requestEnabled := ""
	if p.notificationPermission != app.NotificationDefault {
		requestEnabled = "disabled"
	}

	testEnabled := "disabled"
	if p.notificationPermission == app.NotificationGranted {
		testEnabled = ""
	}

	return NewPage().
		Title("Notifications").
		Icon(bellSVG).
		Index(
			NewIndexLink().Title("Enable Notifications"),
			NewIndexLink().Title("    Current Permission"),
			NewIndexLink().Title("    Request Permission"),
			NewIndexLink().Title("    Display Local Notifications"),
			NewIndexLink().Title("    Example"),

			app.Div().Class("separator"),

			NewIndexLink().Title("Push Notifications"),
			NewIndexLink().Title("    Getting Notification Subscription"),
			NewIndexLink().Title("    Registering Notification Subscription"),
			NewIndexLink().Title("    Sending Push Notification"),

			app.Div().Class("separator"),

			NewIndexLink().Title("Next"),
		).
		Content(
			NewRemoteMarkdownDoc().Src("/web/documents/notifications.md"),

			app.P().Body(
				app.Button().
					Class("button").
					Class(requestEnabled).
					Text("Enable Notifications").
					OnClick(p.enableNotifications),
				app.Button().
					Class("button").
					Class(testEnabled).
					Text("Test Notification").
					OnClick(p.testNotification),
			),

			app.Div().Class("separator"),

			NewRemoteMarkdownDoc().Src("/web/documents/notifications-push.md"),
		)
}

func (p *NotificationsPage) enableNotifications(ctx app.Context, e app.Event) {
	p.notificationPermission = ctx.Notifications().RequestPermission()
	p.registerSubscription(ctx)
}

func (p *NotificationsPage) testNotification(ctx app.Context, e app.Event) {
	n := rand.Intn(43)

	ctx.Notifications().New(app.Notification{
		Title: fmt.Sprintln("go-app test", n),
		Body:  fmt.Sprintln("Test notification for go-app number", n),
		Path:  "/notifications#example",
	})
}

func (p *NotificationsPage) registerSubscription(ctx app.Context) {
	if p.notificationPermission != app.NotificationGranted {
		return
	}

	sub, err := ctx.Notifications().Subscribe(app.Getenv("VAPID_PUBLIC_KEY"))
	if err != nil {
		app.Log(err)
		return
	}

	ctx.Async(func() {
		var body bytes.Buffer
		if err := json.NewEncoder(&body).Encode(sub); err != nil {
			app.Log(errors.New("encoding notification subscription failed").Wrap(err))
			return
		}

		res, err := http.Post("/test/notifications/register", "application/json", &body)
		if err != nil {
			app.Log(errors.New("registering notification subscription failed").Wrap(err))
			return
		}
		defer res.Body.Close()
	})
}

package footer

import (
	"github.com/maxence-charriere/go-app/v10/pkg/app"
)

type Main struct {
	app.Compo
}

func (f *Main) Render() app.UI {
	return app.Div().Body(
		app.Div().Styles(map[string]string{
			"display":         "flex",
			"justify-content": "space-around",
			"padding":         "20px",
		}).Body(
			&github{},
			&button{},
		),
		&created{},
	)
}

type github struct {
	app.Compo
}

func (g *github) Render() app.UI {
	return app.Img().Src("/web/assets/img/octocat.jpg")
}

type button struct {
	app.Compo
}

func (b *button) Render() app.UI {
	return app.A().Class("git-button").Href("#").Text("Go to github.com").OnClick(b.onClick)
}

func (b *button) onClick(ctx app.Context, e app.Event) {
	ctx.Navigate("https://github.com/ekudinov")
}

type created struct {
	app.Compo
}

func (c *created) Render() app.UI {
	return app.Div().Styles(map[string]string{
		"display":         "flex",
		"flex-direction":  "row",
		"justify-content": "center",
		"padding":         "10px",
	}).Body(
		app.Div().Body(
			app.Span().Styles(map[string]string{
				"font-size":  "14px",
				"font-color": "grey",
				"font-style": "italic",
			}).Text("This PWA v2 was created based on:"),
			app.A().Styles(map[string]string{
				"font-size": "14px",
				"padding":   "0 5px",
			}).Text("go-app").Href("https://go-app.dev"),
		),
		app.Div().Styles(map[string]string{
			"font-size":   "12px",
			"font-wiight": "bold",
			"padding":     "5px 10px",
		}).Body(
			app.Text("c2022-2025"),
		),
	)
}

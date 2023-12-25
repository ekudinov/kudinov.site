package content

import (
	"github.com/maxence-charriere/go-app/v9/pkg/app"
)

var isRu = false

var messages = map[string]map[bool]string{
	"str1": {
		true:  "Доброе время суток. В этом мире много технологий и стеков.",
		false: "Good time of day. There are many technologies and stacks in this world.",
	},
	"str2": {
		true:  "Чтобы познать все их нехватит даже несколько человеческих жизней.",
		false: "To know all of them, even a few human lives are not enough.",
	},
	"str3": {
		true:  "Для того чтобы не утонуть в этом океане возможностей, я решил сконцентрироваться на Golang стэке",
		false: "In order not to drown in this ocean of opportunity, I decided to concentrate on Golang stack",
	},
	"stack": {
		true:  "Текущий используемый стек:",
		false: "Current used stack:",
	},
	"desire": {
		true:  "Желание использовать:",
		false: "Desire to try:",
	},
	"hobby": {
		true:  "Делаю в свободное время для души:",
		false: "I do it in my free time for the soul:",
	},
}

type Main struct {
	app.Compo
}

func (c *Main) Render() app.UI {
	stack := &stack{}
	desire := &desire{}
	description := &description{}
	hobby := &hobby{}
	return app.Div().Body(
		&image{},
		app.Div().Styles(map[string]string{
			"display":        "flex",
			"flex-direction": "row",
		}).Body(
			app.Div().Styles(map[string]string{
				"width": "80%",
			}).Body(description),
			app.Div().Styles(map[string]string{
				"width": "20%",
			}).Body(
				&button{
					desire:      desire,
					stack:       stack,
					description: description,
					hobby:       hobby,
				}),
		),
		stack,
		desire,
		hobby,
	)
}

type image struct {
	app.Compo
}

func (i *image) Render() app.UI {
	return app.Div().Styles(map[string]string{
		"padding-bottom": "10px",
	}).Body(app.Img().Styles(map[string]string{
		"width": "100%",
	}).Src("/web/assets/img/black.jpg"))
}

type description struct {
	app.Compo
}

func (d *description) Render() app.UI {
	return app.Div().Styles(map[string]string{
		"padding":   "10px",
		"font-size": "14px",
	}).Body(
		app.Div().Styles(map[string]string{
			"padding-bottom": "5px",
		}).Body(app.If(isRu,
			app.Text(messages["str1"][isRu])).Else(app.Text(messages["str1"][isRu]))),
		app.Div().Styles(map[string]string{
			"padding-bottom": "5px",
		}).Body(app.If(isRu,
			app.Text(messages["str2"][isRu])).Else(app.Text(messages["str2"][isRu]))),
		app.Div().Styles(map[string]string{
			"padding-bottom": "5px",
		}).Body(app.If(isRu,
			app.Text(messages["str3"][isRu])).Else(app.Text(messages["str3"][isRu]))),
	)
}

type button struct {
	app.Compo
	desire      *desire
	stack       *stack
	description *description
	hobby       *hobby
}

func (b *button) Render() app.UI {
	return app.Div().Styles(map[string]string{
		"padding": "10px 0",
	}).Body(
		app.If(isRu,
			app.A().Class("lang-button").Href("#").Text("EN").OnClick(b.onClick)).Else(
			app.A().Class("lang-button").Href("#").Text("RU").OnClick(b.onClick)),
	)
}

func (b *button) onClick(ctx app.Context, e app.Event) {
	isRu = !isRu
	b.desire.Update()
	b.stack.Update()
	b.description.Update()
	b.hobby.Update()

}

type stack struct {
	app.Compo
}

func (s *stack) Render() app.UI {
	return app.Div().Styles(map[string]string{
		"padding":   "5px 10px 0px",
		"font-size": "14px",
	}).Body(
		app.Div().Styles(map[string]string{
			"margin":      "0px 5px 10px 5px",
			"font-weight": "bold",
		}).Body(app.If(isRu,
			app.Text(messages["stack"][isRu])).Else(app.Text(messages["stack"][isRu]))),
		app.Div().Styles(map[string]string{
			"font-style": "italic",
		}).Text("Go, Mysql, Postgresql, Cockrouch, Rabbitmq, Redis, Nginx, Docker, Docker-compose, Gitlab"),
	)
}

type desire struct {
	app.Compo
}

func (d *desire) Render() app.UI {
	return app.Div().Styles(map[string]string{
		"padding":   "5px 10px 0px",
		"font-size": "14px",
	}).Body(
		app.Div().Styles(map[string]string{
			"margin":      "10px 5px",
			"font-weight": "bold",
		}).Body(app.If(isRu,
			app.Text(messages["desire"][isRu])).Else(app.Text(messages["desire"][isRu]))),
		app.Div().Text("Microservices, grpc, data processing, spa&pwa(golang->wasm), highload projects"),
	)
}

type hobby struct {
	app.Compo
}

func (f *hobby) Render() app.UI {
	return app.Div().Styles(map[string]string{
		"padding":   "5px 10px",
		"font-size": "14px",
	}).Body(
		app.Div().Styles(map[string]string{
			"margin":      "10px 5px",
			"font-weight": "bold",
		}).Body(app.If(isRu,
			app.Text(messages["hobby"][isRu])).Else(app.Text(messages["hobby"][isRu]))),
		app.Div().Text("App framework for Ubuntu Touch && Tiling window manager like dwm. (All on golang)"),
	)

}

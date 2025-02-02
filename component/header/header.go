package header

import (
	"github.com/maxence-charriere/go-app/v10/pkg/app"
)

type Main struct {
	app.Compo
}

func (h *Main) Render() app.UI {
	return app.Div().Styles(map[string]string{
		"display":        "flex",
		"flex-direction": "row",
	}).Body(
		app.Div().Styles(map[string]string{
			"padding": "20px",
		}).Body(&avatar{}),
		app.Div().Styles(map[string]string{
			"padding": "20px",
		}).Body(&name{}),
	)
}

type avatar struct {
	app.Compo
}

func (a *avatar) Render() app.UI {
	return app.Img().Styles(map[string]string{
		"object-fit":    "cover",
		"border-radius": "50%",
		"width":         "50px",
		"height":        "50px",
	}).Src("/web/assets/img/avatar.jpg")
}

type name struct {
	app.Compo
}

func (n *name) Render() app.UI {
	return app.Div().Body(
		app.Div().Styles(map[string]string{
			"font-weight": "bold",
			"font-size":   "20px",
		}).Text("Evgeniy Kudinov"),
		app.Div().Styles(map[string]string{
			"padding-top": "10px",
			"font-size":   "14px",
			"color":       "#0000008A",
		}).Body().Text("Backend developer"),
	)
}

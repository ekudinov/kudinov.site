package component

import (
	"github.com/ekudinov/kudinov.site/component/content"
	"github.com/ekudinov/kudinov.site/component/footer"
	"github.com/ekudinov/kudinov.site/component/header"
	"github.com/maxence-charriere/go-app/v10/pkg/app"
)

type MainPage struct {
	app.Compo
	height int
	width  int
}

func (mp *MainPage) Render() app.UI {
	return app.Div().Styles(map[string]string{
		"display":         "flex",
		"justify-content": "center",
	}).Body(
		app.Div().Styles(map[string]string{
			"margin":         "30px",
			"border":         "1px solid #d6bcd6",
			"display":        "flex",
			"flex-direction": "column",
			"max-width":      "600px",
		}).Body(
			&header.Main{},
			&content.Main{},
			&footer.Main{},
		),
	)
}

func (mp *MainPage) OnResize(ctx app.Context) {
	mp.width, mp.height = ctx.Page().Size()
}

func (mp *MainPage) OnMount(ctx app.Context) {
	mp.width, mp.height = ctx.Page().Size()
}

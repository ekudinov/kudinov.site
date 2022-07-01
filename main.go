package main

import (
	"log"
	"net/http"

	"github.com/ekudinov/kudinov.site/component"
	"github.com/maxence-charriere/go-app/v9/pkg/app"
)

func main() {
	app.Route("/", &component.MainPage{})

	app.RunWhenOnBrowser()

	http.Handle("/", &app.Handler{
		Name:        "Evgeniy Kudinov personal page",
		Description: "This is a personal web page created by Evgeniy Kudinov",
		Icon: app.Icon{
			Default: "/web/assets/favicon.ico",
		},
		Styles: []string{
			"/web/assets/css/style.css",
			"/web/assets/fonts/roboto/stylesheet.css",
		},
		Title: "Hello to all!",
	})

	if err := http.ListenAndServe(":8000", nil); err != nil {
		log.Fatal(err)
	}
}

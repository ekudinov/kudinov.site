package main

import (
	"log"
	"net/http"
	"os"

	"github.com/ekudinov/kudinov.site/component"
	"github.com/maxence-charriere/go-app/v10/pkg/app"
)

func main() {
	app.Route("/", func() app.Composer { return &component.MainPage{} })

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

	port := os.Getenv("PORT")
	if port == "" {
		port = "8000"
	}

	log.Printf("Listening on port %s", port)

	if err := http.ListenAndServe(":"+port, nil); err != nil {
		log.Fatal(err)
	}
}

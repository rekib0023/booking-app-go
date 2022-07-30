package main

import (
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/alexedwards/scs/v2"
	"github.com/rekib0023/booking-app-go/pkg/config"
	"github.com/rekib0023/booking-app-go/pkg/render"
)

const port = ":8080"

var app config.AppConfig
var session *scs.SessionManager

func main() {

	app.Debug = true

	session = scs.New()
	session.Lifetime = 24 * time.Hour
	session.Cookie.Persist = true
	session.Cookie.Secure = !app.Debug
	session.Cookie.SameSite = http.SameSiteLaxMode

	app.Session = session

	tc, err := render.CreateTemplateCache()
	if err != nil {
		log.Fatal("Cannot create template cache")
	}
	app.TemplateCache = tc
	app.UseCache = false

	render.NewTemplate(&app)

	fmt.Printf("Listening on port %s\n", port)
	srv := &http.Server{
		Addr:    port,
		Handler: routes(&app),
	}

	err = srv.ListenAndServe()
	log.Fatal(err)
}

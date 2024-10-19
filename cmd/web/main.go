package main

import (
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/alexedwards/scs/v2"
	"github.com/juanchuletas/booking_app/config"
	"github.com/juanchuletas/booking_app/pkg/handlers"
	"github.com/juanchuletas/booking_app/pkg/render"
)

const port string = ":8080"

var app config.AppConfig
var session *scs.SessionManager

func main() {

	//Change when production
	app.InProduction = false

	session = scs.New()
	session.Lifetime = 24 * time.Hour //session for 24 hours
	session.Cookie.Persist = true     //remains when the page is closed
	session.Cookie.SameSite = http.SameSiteDefaultMode
	session.Cookie.Secure = app.InProduction //just for localhost

	app.Session = session

	var err error
	app.TemplateCache, err = render.CreateTemplateCaches()
	if err != nil {

		log.Fatal("cannot create template")

	}
	app.UseCache = false
	render.BuildTemplate(&app)

	repo := handlers.CreateRepo(&app)

	handlers.CreateHandlers(repo)

	// http.HandleFunc("/", handlers.Repo.Home)
	// http.HandleFunc("/about", handlers.Repo.About)

	fmt.Println("Listening on port ", port)
	//_ = http.ListenAndServe(port, nil)

	srv := &http.Server{
		Addr:    port,
		Handler: routes(&app),
	}

	err = srv.ListenAndServe()
	log.Fatal(err)
}

package main

import (
	"fmt"
	"github.com/alexedwards/scs/v2"
	"github.com/marufnwu/go-web/pkg/config"
	"github.com/marufnwu/go-web/pkg/handlers"
	"github.com/marufnwu/go-web/pkg/render"
	"log"
	"net/http"
	"time"
)

var App config.AppConfig
var session *scs.SessionManager

func main() {
	App.InProduction = false

	session = scs.New()
	session.Lifetime = 24 * time.Hour
	session.Cookie.Persist = true
	session.Cookie.SameSite = http.SameSiteLaxMode
	session.Cookie.Secure = App.InProduction

	App.Session = session

	tc, err := render.CreateTemplateCache()
	if err != nil {
		log.Fatal(err)
	}

	App.TemplateCache = tc
	App.UseCache = false

	repo := handlers.NewRepo(&App)
	handlers.NewHandler(repo)

	render.NewTemplates(&App)

	//http.HandleFunc("/", repo.Home)
	//http.HandleFunc("/about", repo.About)

	fmt.Println("Starting server at port: 8080")
	//err = http.ListenAndServe(":8080", nil)
	srv := &http.Server{
		Addr:    ":8080",
		Handler: routes(&App),
	}

	err = srv.ListenAndServe()
	log.Fatal(err)
}

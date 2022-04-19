package render

import (
	"bytes"
	"fmt"
	"github.com/marufnwu/go-bookings-website/pkg/config"
	"github.com/marufnwu/go-bookings-website/pkg/models"
	"html/template"
	"log"
	"net/http"
	"path/filepath"
)

var functions = template.FuncMap{}
var app *config.AppConfig

func NewTemplates(a *config.AppConfig) {
	app = a
}

func AddDefaultData(td *models.TemplateData) *models.TemplateData {
	return td
}

func RenderTemplate(w http.ResponseWriter, tmpl string, td *models.TemplateData) {
	var tc map[string]*template.Template
	if app.UseCache {
		tc = app.TemplateCache
	} else {
		tc, _ = CreateTemplateCache()
	}

	t, ok := tc[tmpl]

	if !ok {
		log.Fatal("template not found")
	}

	buf := new(bytes.Buffer)

	td = AddDefaultData(td)

	err := t.Execute(buf, td)
	if err != nil {
		return
	}

	_, err = buf.WriteTo(w)
	if err != nil {
		return
	}
}

func CreateTemplateCache() (map[string]*template.Template, error) {
	myCache := map[string]*template.Template{}

	pages, err := filepath.Glob("G:\\Go\\Go web development\\templates\\*.page.tmpl")

	if err != nil {
		return myCache, err
	}

	for _, page := range pages {
		name := filepath.Base(page)

		fmt.Println("Page is currently: ", page)

		ts, err := template.New(name).Funcs(functions).ParseFiles(page)

		if err != nil {
			return myCache, err

		}

		matches, err := filepath.Glob("G:\\Go\\Go web development\\templates\\*.layout.tmpl")
		if err != nil {
			return myCache, err
		}

		fmt.Println("LayoutFile", matches)

		if len(matches) > 0 {
			ts, err = ts.ParseGlob("G:\\Go\\Go web development\\templates\\*.layout.tmpl")
			if err != nil {
				return myCache, err
			}

			fmt.Println("Template", ts)
		}
		myCache[name] = ts

	}

	return myCache, nil
}

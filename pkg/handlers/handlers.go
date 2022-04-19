package handlers

import (
	"github.com/marufnwu/go-bookings-website/pkg/config"
	"github.com/marufnwu/go-bookings-website/pkg/models"
	"github.com/marufnwu/go-bookings-website/pkg/render"
	"net/http"
)

var Repo *Repository

type Repository struct {
	App *config.AppConfig
}

func NewRepo(a *config.AppConfig) *Repository {
	return &Repository{
		App: a,
	}
}

func NewHandler(r *Repository) {
	Repo = r
}

func (r *Repository) About(writer http.ResponseWriter, request *http.Request) {
	stringMap := make(map[string]string)
	stringMap["test"] = "I am Maruf"

	remoteIp := r.App.Session.GetString(request.Context(), "remote_ip")
	stringMap["remote_ip"] = remoteIp

	render.RenderTemplate(writer, "about.page.tmpl", &models.TemplateData{
		StringMap: stringMap,
	})

}

func (r *Repository) Home(writer http.ResponseWriter, request *http.Request) {
	remoteIp := request.RemoteAddr
	r.App.Session.Put(request.Context(), "remote_ip", remoteIp)
	render.RenderTemplate(writer, "home.page.tmpl", &models.TemplateData{})
}

package handlers

import (
	"fmt"
	"github.com/marufnwu/go-bookings-website/internal/config"
	"github.com/marufnwu/go-bookings-website/internal/forms"
	"github.com/marufnwu/go-bookings-website/internal/models"
	"github.com/marufnwu/go-bookings-website/internal/render"
	"log"
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

func (repo *Repository) About(writer http.ResponseWriter, request *http.Request) {
	stringMap := make(map[string]string)
	stringMap["test"] = "I am Maruf"

	remoteIp := repo.App.Session.GetString(request.Context(), "remote_ip")
	stringMap["remote_ip"] = remoteIp

	render.RenderTemplate(writer, request, "about.page.tmpl", &models.TemplateData{
		StringMap: stringMap,
	})

}

func (repo *Repository) Home(writer http.ResponseWriter, request *http.Request) {
	render.RenderTemplate(writer, request, "home.page.tmpl", &models.TemplateData{})
}

func (repo *Repository) Contact(writer http.ResponseWriter, request *http.Request) {
	render.RenderTemplate(writer, request, "home.page.tmpl", &models.TemplateData{})
}

func (repo *Repository) Generals(writer http.ResponseWriter, request *http.Request) {
	render.RenderTemplate(writer, request, "generals.page.tmpl", &models.TemplateData{})
}

func (repo *Repository) Majors(writer http.ResponseWriter, request *http.Request) {
	render.RenderTemplate(writer, request, "majors.page.tmpl", &models.TemplateData{})
}
func (repo *Repository) Availability(writer http.ResponseWriter, request *http.Request) {
	render.RenderTemplate(writer, request, "search-avaiability.page.tmpl", &models.TemplateData{})
}
func (repo *Repository) PostAvailability(writer http.ResponseWriter, request *http.Request) {

	start := request.Form.Get("start_date")
	end := request.Form.Get("end_date")

	_, err := fmt.Fprintf(writer, "Start date is %s and end date is %s", start, end)
	if err != nil {
		return
	}
}

func (repo *Repository) Reservation(writer http.ResponseWriter, request *http.Request) {
	render.RenderTemplate(writer, request, "make-reservation.page.tmpl", &models.TemplateData{
		Form: forms.New(nil),
	})
}

func (repo *Repository) PostReservation(writer http.ResponseWriter, r *http.Request) {
	err := r.ParseForm()
	if err != nil {
		log.Println(err)
	}

	reservation := models.Reservation{
		FirstName: r.Form.Get("first_name"),
		LastName:  r.Form.Get("last_name"),
		Phone:     r.Form.Get("phone"),
		Email:     r.Form.Get("email"),
	}

	form := forms.New(r.PostForm)

	form.Has("first_name", r)
	if !form.Valid() {
		data := make(map[string]interface{})
		data["reservation"] = reservation

		render.RenderTemplate(writer, r, "make-reservation.page.tmpl", &models.TemplateData{
			Form: form,
			Data: data,
		})

		return

	}
}

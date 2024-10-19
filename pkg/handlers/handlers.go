package handlers

import (
	"errors"
	"fmt"
	"net/http"

	"github.com/juanchuletas/booking_app/config"
	"github.com/juanchuletas/booking_app/pkg/models"
	"github.com/juanchuletas/booking_app/pkg/render"
)

type Repository struct {
	App *config.AppConfig
}

var Repo *Repository

func CreateRepo(inApp *config.AppConfig) *Repository {

	return &Repository{
		App: inApp,
	}

}

func CreateHandlers(r *Repository) {
	Repo = r
}

func (r_reciver *Repository) Home(w http.ResponseWriter, req *http.Request) {
	remoteIP := req.RemoteAddr

	r_reciver.App.Session.Put(req.Context(), "remote_ip", remoteIP)

	render.RenderTemplate(w, "home.page.html", &models.TemplateData{})
}

func (r_reciver *Repository) About(w http.ResponseWriter, req *http.Request) {

	stringMap := map[string]string{}

	stringMap["Test"] = "Testing a data"

	remoteIP := r_reciver.App.Session.GetString(req.Context(), "remote_ip")

	stringMap["remote_ip"] = remoteIP
	render.RenderTemplate(w, "about.page.html", &models.TemplateData{
		StringMap: stringMap,
	})
}
func Divide(w http.ResponseWriter, req *http.Request) {

	var a = 13.0
	var b = 10.0

	f, err := dividevalues(a, b)
	if err != nil {
		fmt.Fprintf(w, "cannot divide by zero  : %f, %f\n", a, b)
		return
	}
	fmt.Fprintf(w, "The division %f / %f is : %f\n", a, b, f)

}
func dividevalues(x, y float64) (float64, error) {

	if y == 0.0 {
		err := errors.New("cannot divide by zero")
		return 0.0, err
	}
	result := x / y

	return result, nil

}

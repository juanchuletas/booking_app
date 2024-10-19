package render

import (
	"bytes"
	"fmt"
	"log"
	"net/http"
	"path/filepath"
	"text/template"

	"github.com/juanchuletas/booking_app/config"
	"github.com/juanchuletas/booking_app/pkg/models"
)

const pathTemplates = "../../templates/*.page.html"
const pathLayout = "../../templates/*layout.html"

var app *config.AppConfig

func BuildTemplate(inApp *config.AppConfig) {

	app = inApp
}

func RenderTemplate(w http.ResponseWriter, in_tmpl string, inData *models.TemplateData) {
	//get the template cache from the app config

	//Developer Mode!! decide if you need to use cache
	var tempCache map[string]*template.Template
	if app.UseCache {
		tempCache = app.TemplateCache
	} else {
		tempCache, _ = CreateTemplateCaches()
	}
	log.Println("Getting the template from the app config")

	//get requested template from cache
	tmpl, ok := tempCache[in_tmpl] //tmpl is the template
	if !ok {
		log.Fatal("Could not get template from template cache")
	}
	buffer := new(bytes.Buffer) //holds bytes
	err := tmpl.Execute(buffer, inData)
	if err != nil {
		log.Println(err)
	}
	//render the template:
	_, err = buffer.WriteTo(w)
	if err != nil {
		log.Println(err)
	}

}

var templateCache = make(map[string]*template.Template)

func RenderTemplateOld(w http.ResponseWriter, in_tmpl string) {

	var tmpl *template.Template
	var err error

	//check if we have the template on the cache
	_, inMap := templateCache[in_tmpl]

	if !inMap { //If we do not have the rtemplate
		log.Println("creating template and adding to cache")
		err = createTemplateCache(in_tmpl)
		if err != nil {
			log.Println(err)
		}
	} else {
		//We have the template in the cache

		log.Println("Using cached template")
	}

	tmpl = templateCache[in_tmpl]

	err = tmpl.Execute(w, nil)
	if err != nil {
		log.Println(err)
	}

}

func createTemplateCache(in_tmpl string) error {
	templates := []string{

		fmt.Sprintf("../../templates/%s", in_tmpl),
		"../../templates/base.layout.html",
	}
	//Parse the te plate

	tmpl, err := template.ParseFiles(templates...)
	if err != nil {

		return err
	}
	//add template to cache
	templateCache[in_tmpl] = tmpl

	return nil
}
func CreateTemplateCaches() (map[string]*template.Template, error) {

	myCache := map[string]*template.Template{}
	//get all the files with the end page.html
	pages, err := filepath.Glob(pathTemplates)
	if err != nil {
		return myCache, err
	}
	//range throug all the files ending with *page.html
	for _, item := range pages {
		name := filepath.Base(item)
		templateSet, err := template.New(name).ParseFiles(item)
		if err != nil {
			return myCache, err
		}
		match, err := filepath.Glob(pathLayout)
		if err != nil {
			return myCache, err
		}
		if len(match) > 0 { //Means that we have a layout
			templateSet, err = templateSet.ParseGlob(pathLayout)
			if err != nil {
				return myCache, err
			}
		}

		myCache[name] = templateSet
	}

	return myCache, nil
}

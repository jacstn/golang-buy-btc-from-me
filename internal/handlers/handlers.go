package handlers

import (
	"fmt"
	"html/template"
	"net/http"

	"github.com/jacstn/golang-buy-btc-from-me/config"
	"github.com/jacstn/golang-buy-btc-from-me/internal/forms"
	"github.com/jacstn/golang-buy-btc-from-me/internal/models"
)

var app *config.AppConfig

func NewHandlers(c *config.AppConfig) {
	app = c
}

func GetOmisePublicKey(w http.ResponseWriter, r *http.Request) {

}

func Home(w http.ResponseWriter, r *http.Request) {
	data := make(map[string]interface{})

	renderTemplate(w, "new-url", &models.TemplateData{
		Form: forms.New(nil),
		Data: data,
	})
}

func renderTemplate(w http.ResponseWriter, templateName string, data *models.TemplateData) {
	parsedTemplate, _ := template.ParseFiles("./templates/"+templateName+".go.tmpl", "./templates/base.layout.go.tmpl")

	err := parsedTemplate.Execute(w, data)
	if err != nil {
		fmt.Fprintf(w, "Error handling template page!!", err)
	}
}

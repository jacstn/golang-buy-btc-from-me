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
	w.Header().Set("Content-Type", "application/json")
	retVal := fmt.Sprintf("{\"OmisePKey\":\"%s\"}", app.OmisePublicKey)
	fmt.Fprint(w, retVal)
}

func Home(w http.ResponseWriter, r *http.Request) {
	data := make(map[string]interface{})

	renderTemplate(w, "home", &models.TemplateData{
		Form: forms.New(nil),
		Data: data,
	})
}

func renderTemplate(w http.ResponseWriter, templateName string, data *models.TemplateData) {
	parsedTemplate, _ := template.ParseFiles("./templates/"+templateName+".tmpl", "./templates/base.layout.tmpl")

	err := parsedTemplate.Execute(w, data)
	if err != nil {
		fmt.Fprintf(w, "Error handling template page!!", err)
	}
}

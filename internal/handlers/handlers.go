package handlers

import (
	"encoding/json"
	"fmt"
	"html/template"
	"net/http"
	"strconv"

	"github.com/jacstn/golang-buy-btc-from-me/config"
	"github.com/jacstn/golang-buy-btc-from-me/internal/ext"
	"github.com/jacstn/golang-buy-btc-from-me/internal/forms"
	"github.com/jacstn/golang-buy-btc-from-me/internal/models"
	"github.com/justinas/nosurf"
)

var app *config.AppConfig

func NewHandlers(c *config.AppConfig) {
	app = c
}

func CreateOrder(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	r.ParseForm()
	form := forms.New(r.PostForm)
	form.ValidAddress("btc_addr", r)
	form.Has("btc_addr", r)
	form.ValidAmunt("btc_amount", r)
	form.Has("btc_amount", r)
	form.ValidAmunt("btc_price", r)
	form.Has("btc_price", r)

	if form.Valid() {
		fmt.Println("form valid")
		btcAmountFl, _ := strconv.ParseFloat(r.Form.Get("btc_amount"), 64)
		btcAmount := uint64(btcAmountFl) * app.BTCDecimals
		btcPrice, _ := strconv.ParseFloat(r.Form.Get("btc_price"), 64)

		usdAmount := btcPrice * btcAmountFl * 100

		o := models.Order{
			BTCAmount: btcAmount,
			Address:   r.Form.Get("btc_addr"),
			USDAmount: uint64(usdAmount),
		}

		err := models.NewOrder(app.DB, &o)
		if err != nil {
			fmt.Fprint(w, "{\"status\":\"err\", \"error while saving model\"}")
			return
		}
		fmt.Fprint(w, "{\"status\":\"ok\"}")
		return
	} else {
		formErrors, _ := json.Marshal(form.Errors)

		fmt.Fprintf(w, "{\"status\":\"err\", \"errors\": %s}", string(formErrors))
	}

}

func GetBTCPrice(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	fmt.Fprintf(w, "{\"btc_price\":%s, \"sell_margin\":%.2f}", ext.GetBTCPrice(), app.SellMargin)
}

func UpdateOrderStatus(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "{\"status\":\"ok\"}")
}

func Home(w http.ResponseWriter, r *http.Request) {
	data := make(map[string]interface{})
	data["omise_key"] = app.OmisePublicKey
	data["csrf_token"] = nosurf.Token(r)

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

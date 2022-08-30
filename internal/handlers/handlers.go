package handlers

import (
	"encoding/json"
	"fmt"
	"html/template"
	"io/ioutil"
	"log"
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
		btcAmount, _ := strconv.ParseFloat(r.Form.Get("btc_amount"), 64)
		satoshiAmount := uint64(btcAmount) * app.BTCDecimals

		btcPrice, _ := strconv.ParseFloat(r.Form.Get("btc_price"), 64)

		usdAmount := btcPrice * btcAmount * 100

		o := models.Order{
			SatoshiAmount: satoshiAmount,
			Address:       r.Form.Get("btc_addr"),
			USDAmount:     uint64(usdAmount),
		}

		id, err := models.NewOrder(app.DB, &o)
		fmt.Println(o)
		if err != nil {
			fmt.Fprint(w, "{\"status\":\"err\", \"error while saving model\"}")
			return
		}
		fmt.Fprintf(w, "{\"status\":\"ok\", \"order_id\":%d}", id)
		return
	} else {
		fmt.Println("form invalid")
		formErrors, _ := json.Marshal(form.Errors)

		fmt.Fprintf(w, "{\"status\":\"err\", \"errors\": %s}", string(formErrors))
	}

}

func Charge(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	d := struct {
		Amount     float32 `json:"amount"`
		OrderId    int64   `json:"orderId"`
		OmiseToken string  `json:"omiseToken"`
	}{}

	body, err := ioutil.ReadAll(r.Body)

	if err != nil {
		log.Println("error while read all io util")
	}

	err = json.Unmarshal(body, &d)
	if err != nil {
		fmt.Println(err)
		log.Println("error unmarshaling json")
	}

	amount := int64(d.Amount * 100)
	o := models.GetOrderById(app.DB, d.OrderId)
	_, err = ext.CreateOmiseCharge(&d.OmiseToken, &amount, &app.OmisePublicKey, &app.OmisePrivateKey)

	if err != nil {
		fmt.Fprintf(w, "{\"status\":\"err\", \"err\": \"%s\"}", err)
		return
	}

	o.Status = "PAYM_COMPL"
	o.Save(app.DB)
	fmt.Fprint(w, "{\"status\":\"ok\"}")
}

func GetBTCBalance(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	balance := ext.GetBtcBalance()
	fmt.Fprintf(w, "{\"btc_balance\":%.8f}", balance)
}

func GetBTCPrice(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	btcPrice := ext.GetBTCPrice()
	if btcPrice == "0" {
		fmt.Fprint(w, "{\"err\":\"unable to get current bitcoin price\"")
	}
	fmt.Fprintf(w, "{\"btc_price\":%s, \"sell_margin\":%.2f}", btcPrice, app.SellMargin)
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

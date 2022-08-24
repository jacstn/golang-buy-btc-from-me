package models

import "github.com/jacstn/golang-buy-btc-from-me/internal/forms"

type TemplateData struct {
	Data map[string]interface{}
	Form *forms.Form
}

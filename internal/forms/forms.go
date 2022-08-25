package forms

import (
	"net/http"
	"net/url"
	"strconv"

	"github.com/jacstn/golang-buy-btc-from-me/internal/ext"
)

type Form struct {
	url.Values
	Errors errors
}

func New(data url.Values) *Form {
	return &Form{
		data,
		errors(map[string][]string{}),
	}
}

func (f *Form) ValidAddress(field string, r *http.Request) bool {
	if ext.IsValidBTCAddress(r.Form.Get(field)) != true {
		f.Errors.Add(field, "Bitcoin Address is invalid")
		return false
	}

	return true
}

func (f *Form) Has(field string, r *http.Request) bool {
	if r.Form.Get(field) == "" {
		f.Errors.Add(field, "This field cannot be empty")
		return false
	}

	return true
}

func (f *Form) ValidAmunt(field string, r *http.Request) bool {
	amount := r.Form.Get(field)
	if _, err := strconv.ParseFloat(amount, 64); err != nil {
		f.Errors.Add(field, "Not a valid amount")
		return false
	}
	return true
}

func (f *Form) Valid() bool {
	return len(f.Errors) == 0
}

package ext

import (
	"errors"
	"log"

	"github.com/omise/omise-go"
	"github.com/omise/omise-go/operations"
)

func CreateOmiseCharge(omiseToken *string, amount *int64, pKey *string, sKey *string) (string, error) {
	client, e := omise.NewClient(*pKey, *sKey)

	if e != nil {
		log.Println(e)
		return "", errors.New("cannot create omise client object")
	}
	charge, createCharge := &omise.Charge{}, &operations.CreateCharge{
		Amount:   *amount,
		Currency: "usd",
		Card:     *omiseToken,
	}

	log.Println(charge, createCharge)

	if e = client.Do(charge, createCharge); e != nil {
		log.Println(e)
		return "", errors.New("cannot create charge")
	}

	log.Println(e)

	log.Printf("charge: %s  amount: %s %d\n", charge.ID, charge.Currency, charge.Amount)
	return charge.ID, nil
}

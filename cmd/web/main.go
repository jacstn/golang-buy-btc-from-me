package main

import (
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/alexedwards/scs/v2"
	"github.com/jacstn/golang-buy-btc-from-me/config"
	"github.com/jacstn/golang-buy-btc-from-me/internal/database"
	"github.com/jacstn/golang-buy-btc-from-me/internal/handlers"
)

func main() {
	err := run()
	if err != nil {
		panic("error while initializing application")
	}

	handlers.NewHandlers(&app)
	fmt.Println("Starting application", portNumber)

	srv := &http.Server{
		Addr:    portNumber,
		Handler: routes(),
	}

	err = srv.ListenAndServe()
	app.DB.Close()
	if err != nil {
		log.Fatal("Cannot start server")
	}
}

const portNumber = ":3000"

var app = config.AppConfig{
	Production: false,
}

func run() error {
	session := scs.New()
	session.Lifetime = 24 * time.Hour
	session.Cookie.Persist = true
	session.Cookie.Secure = app.Production

	db := database.Connect()
	app.DB = db
	var err error

	if err != nil {
		fmt.Println("error while Reading Char Array from file")
		return err
	}

	app.Session = session

	return nil
}

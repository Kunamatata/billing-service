package main

import (
	"log"
	"net/http"
	"os"

	"github.com/stripe/stripe-go"
)

type application struct {
}

func run() {

	app := application{}
	if os.Getenv("ENV") == "production" {
		stripe.Key = os.Getenv("STRIPE_KEY")
	} else {
		stripe.Key = os.Getenv("TEST_STRIPE_KEY")
	}

	srv := http.Server{
		Addr:    ":8080",
		Handler: app.routes(),
	}

	log.Println("Server listening on", srv.Addr)
	err := srv.ListenAndServe()
	log.Fatal(err)
}

func main() {
	run()
}

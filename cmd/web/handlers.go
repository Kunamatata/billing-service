package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"runtime/debug"

	"stripe-billing-service/pkg/models"

	"github.com/stripe/stripe-go"
	"github.com/stripe/stripe-go/charge"
)

func (app *application) createCharge(w http.ResponseWriter, r *http.Request) {
	if r.Method != "POST" {
		http.Error(w, http.StatusText(http.StatusMethodNotAllowed), http.StatusMethodNotAllowed)
		return
	}

	var stripeCharge models.StripeCharge
	json.NewDecoder(r.Body).Decode(&stripeCharge)

	_, err := charge.New(&stripe.ChargeParams{
		Amount:       stripe.Int64(stripeCharge.Amount),
		Currency:     stripe.String(string(stripe.CurrencyUSD)),
		Description:  stripe.String("my first test charge"),
		ReceiptEmail: stripe.String(stripeCharge.ReceiptEmail),
		Source:       &stripe.SourceParams{Token: stripe.String("tok_visa")},
	})

	if err != nil {
		log.Println(fmt.Sprintf("%s\n%s", err.Error(), debug.Stack()))
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
	}

	w.WriteHeader(http.StatusOK)
}

package models

// StripeCharge incoming data for Stripe API
type StripeCharge struct {
	Amount       int64  `json:"amount"`
	ReceiptEmail string `json:"receiptEmail"`
}

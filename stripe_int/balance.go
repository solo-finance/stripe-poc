package stripe_int

import (
	"github.com/stripe/stripe-go/v76"
	"github.com/stripe/stripe-go/v76/balance"
	"log"
)

func (s *Stripe) getBalance() interface{} {
	params := &stripe.BalanceParams{}
	data, err := balance.Get(params)
	if err != nil {
		log.Println(err)
	}
	var amount int64
	for _, val := range data.Available {
		if val.Currency == "usd" {
			amount = val.Amount
		}
	}
	res := map[string]interface{}{
		"amount":            amount,
		"iso_currency_code": "USD",
	}
	return res
}

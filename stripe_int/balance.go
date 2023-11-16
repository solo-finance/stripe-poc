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
	return data
}

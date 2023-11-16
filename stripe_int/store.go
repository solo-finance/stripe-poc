package stripe_int

import (
	"github.com/stripe/stripe-go/v76"
	"github.com/stripe/stripe-go/v76/payout"
)

func (s *Stripe) getStore() interface{} {
	params := &stripe.PayoutListParams{}
	list := payout.List(params).PayoutList()
	return list
}

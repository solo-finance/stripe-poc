package stripe_int

import (
	"github.com/stripe/stripe-go/v76"
	"github.com/stripe/stripe-go/v76/payout"
)

func (s *Stripe) getPayouts() interface{} {
	params := &stripe.PayoutListParams{}
	list := payout.List(params).PayoutList()
	return list
}

func (s *Stripe) getPayout() interface{} {
	params := &stripe.PayoutParams{}
	data, _ := payout.Get("po_1N3coX2eZvKYlo2CrFLvNX0K", params)
	return data
}

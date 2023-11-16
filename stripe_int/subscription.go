package stripe_int

import (
	"github.com/stripe/stripe-go/v76"
	"github.com/stripe/stripe-go/v76/subscription"
	"log"
)

func (s *Stripe) getSubscriptions() interface{} {
	params := &stripe.SubscriptionListParams{}
	return subscription.List(params).SubscriptionList()
}

func (s *Stripe) getSubscription() interface{} {
	params := &stripe.SubscriptionParams{}
	data, err := subscription.Get("sub_1OD4qO2eZvKYlo2CIWOQU25g", params)
	if err != nil {
		log.Println(err)
	}
	return data
}

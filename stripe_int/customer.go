package stripe_int

import (
	"github.com/stripe/stripe-go/v76"
	"github.com/stripe/stripe-go/v76/customer"
)

func (s *Stripe) getCustomerList() interface{} {
	params := &stripe.CustomerListParams{}
	list := customer.List(params).CustomerList().Data
	return list
}

func (s *Stripe) getCustomer() interface{} {
	params := &stripe.CustomerParams{}
	data, _ := customer.Get("cus_P18gmbwLHAbLp3", params)
	return data
}

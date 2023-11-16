package stripe_int

import (
	"github.com/stripe/stripe-go/v76"
	"github.com/stripe/stripe-go/v76/taxrate"
)

func (s *Stripe) getTaxRates() interface{} {
	params := &stripe.TaxRateListParams{}
	list := taxrate.List(params).TaxRateList()
	return list
}

func (s *Stripe) getTaxRate() interface{} {
	params := &stripe.TaxRateParams{}
	data, _ := taxrate.Get("txr_1OD75f2eZvKYlo2ChcgBSsUD", params)
	return data
}

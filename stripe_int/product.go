package stripe_int

import (
	"github.com/stripe/stripe-go/v76"
	"github.com/stripe/stripe-go/v76/product"
)

func (s *Stripe) getProducts() interface{} {
	params := &stripe.ProductListParams{}
	return product.List(params).ProductList()
}

func (s *Stripe) getProduct() interface{} {
	params := &stripe.ProductParams{}
	data, _ := product.Get("prod_P18lma20MZNFCK", params)
	return data
}

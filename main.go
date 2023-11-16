package main

import "github.com/solo-finance/stripe-integration-poc/stripe_int"

func main() {
	stripeInt := stripe_int.New()
	stripeInt.GetAll()
}

package stripe_int

import (
	"github.com/stripe/stripe-go/v76"
	"github.com/stripe/stripe-go/v76/balancetransaction"
)

func (s *Stripe) getBalanceTransactions() interface{} {
	params := &stripe.BalanceTransactionListParams{}
	list := balancetransaction.List(params).BalanceTransactionList().Data
	return list
}

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

func (s *Stripe) getBalanceTransaction() interface{} {
	params := &stripe.BalanceTransactionParams{}
	data, _ := balancetransaction.Get("txn_3OD6ZX2eZvKYlo2C1uII11Gl", params)
	return data
}

package stripe_int

import (
	"encoding/json"
	"github.com/stripe/stripe-go/v76"
	"os"
)

type Stripe struct {
	Key string
}

func New() *Stripe {
	stripe.Key = "sk_test_4eC39HqLyjWDarjtT1zdp7dc"
	return &Stripe{Key: "sk_test_4eC39HqLyjWDarjtT1zdp7dc"}
}

func (s *Stripe) GetAll() {
	var data interface{}
	//data = s.getCustomerList()
	//constructJson("customers", data)
	//
	//data = s.getCustomer()
	//constructJson("customer", data)
	//
	//data = s.getProducts()
	//constructJson("products", data)
	//
	//data = s.getProduct()
	//constructJson("product", data)
	//
	//data = s.getBalanceTransactions()
	//constructJson("transactions", data)
	//
	//data = s.getBalanceTransaction()
	//constructJson("transaction", data)
	//
	//data = s.getBalance()
	//constructJson("balance", data)
	//
	//data = s.getSubscriptions()
	//constructJson("subscriptions", data)

	data = s.getSubscription()
	constructJson("subscription", data)
}

func constructJson(name string, obj interface{}) {
	file, _ := json.MarshalIndent(obj, "", " ")
	_ = os.WriteFile("poc_output/"+name+".json", file, 0644)
}

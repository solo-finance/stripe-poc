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
	data := s.getStore()
	constructJson("store", data)
}

func constructJson(name string, obj interface{}) {
	file, _ := json.MarshalIndent(obj, "", " ")
	_ = os.WriteFile("poc_output/"+name+".json", file, 0644)
}

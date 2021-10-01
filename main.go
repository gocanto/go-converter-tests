package main

import (
	"fmt"
	"github.com/voyago/converter/environment"
	"github.com/voyago/converter/pkg/store"
)

func main() {
	env, _ := environment.Make("converter-test")
	fmt.Println(fmt.Sprintf("App name: %s", env.Get("APP_NAME")))
	fmt.Println(fmt.Sprintf("App env: %s [live?: %v]\n", env.Get(environment.EnvKey), env.IsLive()))

	manager, _ := store.Make(store.NewCurrencyLayerRequest(env, "SGD"))
	rates, _ := manager.GetExchangeRates()
	sgd, _ := rates.Find("SGD")

	fmt.Println(rates.Count(), sgd)
}

package main

import (
	"fmt"
	"github.com/voyago/converter/environment"
	"github.com/voyago/converter/pkg/conversion"
	"github.com/voyago/converter/pkg/model"
	"github.com/voyago/converter/pkg/store"
	"github.com/voyago/converter/pkg/store/handler/currencyLayer"
)

func main() {
	divider("|          CURRENCY LAYER API DATA                    |")
	fetchCurrencyLayerData()

	divider("|               TESTING DATA                          |")
	fetchTestingData()
}

func fetchCurrencyLayerData() {
	env, _ := environment.Make("converter-test")

	if env.Get(currencyLayer.ApiKeyName) == "" {
		fmt.Println("Skipping ...")
		fmt.Println("Please, add a valid key for your currency layer setup. Ref: ", currencyLayer.ApiKeyName)

		return
	}

	manager, _ := store.Make(store.NewCurrencyLayerRequest(env, "SGD"))

	currencies, _ := manager.GetExchangeRates()
	usd, _ := currencies.Find("USD")

	fmt.Println("--- Environment")
	fmt.Println(fmt.Sprintf("App name: %s", env.Get("APP_NAME")))
	fmt.Println(fmt.Sprintf("App env: %s [live?: %v]\n", env.Get(environment.EnvKey), env.IsLive()))

	fmt.Println("--- Currencies exchange rate")
	fmt.Println("Available currencies rate:", currencies.Count())
	fmt.Println("USD rate:", usd.Rate)

	fmt.Println("\n--- Conversions")
	converter := conversion.Make(*manager)
	price, _ := model.MakePrice(usd, 1)
	result, _ := converter.Convert(price)
	fmt.Println(fmt.Sprintf("Result: [%s], Amount: [%.2f]", result.ToString(), result.ToFloat()))
}

func fetchTestingData() {
	env, _ := environment.MakeWith("converter-test", ".test.env.example")

	manager, _ := store.Make(store.NewCurrencyLayerRequest(env, "SGD"))

	rates, _ := manager.GetExchangeRates()
	sgd, _ := rates.Find("SGD")

	fmt.Println("--- Environment")
	fmt.Println(fmt.Sprintf("App name: %s", env.Get("APP_NAME")))
	fmt.Println(fmt.Sprintf("App env: %s [live?: %v]\n", env.Get(environment.EnvKey), env.IsLive()))

	fmt.Println("--- Currencies exchange rate")
	fmt.Println("Available currencies rate:", rates.Count())
	fmt.Println("SGD rate:", sgd.Rate)

	fmt.Println("\n--- Conversions")
	converter := conversion.Make(*manager)
	price, _ := model.MakePrice(sgd, 1)
	result, _ := converter.Convert(price)
	fmt.Println(fmt.Sprintf("Result: [%s], Amount: [%.2f]", result.ToString(), result.ToFloat()))
}

func divider(message string)  {
	fmt.Println("")
	fmt.Println("+ --------------------------------------------------- +")
	fmt.Println(message)
	fmt.Println("+ --------------------------------------------------- +")
	fmt.Println("")
}

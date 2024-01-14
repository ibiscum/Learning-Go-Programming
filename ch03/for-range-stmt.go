//go:build ignore

package main

import (
	"fmt"
	"strings"
)

type Curr struct {
	Currency string
	Name     string
	Country  string
	Number   int
}

var currencies = []Curr{
	{"NOK", "Norwegian Krone", "Norwary", 578},
	{"KES", "Kenyan Shilling", "Kenya", 404},
	{"DZD", "Algerian Dinar", "Algeria", 12},
	{"AUD", "Australian Dollar", "Australia", 36},
	{"MXN", "Mexican Peso", "Mexico", 484},
	{"GRD", "Drachma", "Greece", 978},
	{"KHR", "Riel", "Cambodia", 116},
	{"SZL", "Lilangeni", "Swaziland", 748},
	{"GBP", "Pound Sterling", "Isle of Man", 826},
	{"HTG", "Gourde", "Haiti", 332},
	{"BWP", "Pula", "Botswana", 72},
	{"CLP", "Chilean Peso", "Chile", 152},
	{"HKD", "Hong Kong Dollar", "Hong Koong", 344},
	{"HTG", "Gourde", "Haiti", 332},
	{"TRY", "Turkish Lira", "Turkey", 949},
	{"EUR", "Euro", "Belgium", 978},
	{"JMD", "Jamaican Dollar", "Jamaica", 388},
	{"ALL", "Lek", "Albania", 8},
	{"GEL", "Lari", "Georgia", 981},
	{"KFM", "Coromo Franc", "Comoros", 174},
	{"NZD", "New Zeland Dollar", "Tokelau", 554},
}

var sortedCurrs []Curr

func main() {
	updateCurrency(978, Curr{"EUR", "Euro", "Greece", 978})
	printCurrencies()
}

func updateCurrency(number int, newCurr Curr) {
	for i, v := range currencies {
		if number == v.Number {
			currencies[i] = newCurr
			fmt.Printf("%v -> %v\n", v, newCurr)
			break
		}
	}
}

func printCurrencies() {
	for i := range currencies {
		fmt.Printf("%d: %v\n", i, currencies[i])
	}
}

func find(filter string) {
	for _, c := range currencies {
		switch {
		case strings.Contains(c.Currency, filter),
			strings.Contains(c.Name, filter),
			strings.Contains(c.Country, filter):
			fmt.Println("Found", c)
		}
	}
}

func doEmptyRange() {
	for range currencies {
		fmt.Println("A currency found.")
	}
}

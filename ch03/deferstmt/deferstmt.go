//go:build ignore

package main

import (
	"fmt"
	"math/rand"
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
	{"EUR", "Euro", "Greece", 978},
	{"KHR", "Riel", "Cambodia", 116},
	{"SZL", "Lilangeni", "Swaziland", 748},
	{"GBP", "Pound Sterling", "Isle of Man", 826},
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

func main() {
	sort()
	print()
	print()
}

func shuffle() {
	n := len(currencies)
	for i := range currencies {
		next := rand.Intn(n)
		temp := currencies[i]
		currencies[i] = currencies[next]
		currencies[next] = temp
	}
}

func print() {
	defer shuffle()
	fmt.Println("---- Currencies ----")
	for i, v := range currencies {
		fmt.Printf("%d: %v\n", i, v)
	}
}

func sort() {
	N := len(currencies)
	for i := 0; i < N-1; i++ {
		currMin := i
		for k := i + 1; k < N; k++ {
			if currencies[k].Currency < currencies[currMin].Currency {
				currMin = k
			}
		}
		// swap
		if currMin != i {
			temp := currencies[i]
			currencies[i] = currencies[currMin]
			currencies[currMin] = temp
		}
	}
}

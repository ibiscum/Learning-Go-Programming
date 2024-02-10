package main

import "github.com/ibiscum/Learning-Go-Programming/ch02lib"

func main() {
	currency := "US Dollar"
	country := "United States"
	code := "USD"
	//fmt.Println("The currency of", country, "is the", currency, "(", code, ")")
	ch02lib.PrintCurrencyInfo(country, currency, code)

}

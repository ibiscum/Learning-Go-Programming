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
	{"DZD", "Algerian Dinar", "Algeria", 12},
	{"AUD", "Australian Dollar", "Australia", 36},
	{"EUR", "Euro", "Belgium", 978},
	{"CLP", "Chilean Peso", "Chile", 152},
	{"EUR", "Euro", "Greece", 978},
	{"HTG", "Gourde", "Haiti", 332},
	{"HKD", "Hong Kong Dollar", "Hong Koong", 344},
	{"KES", "Kenyan Shilling", "Kenya", 404},
	{"MXN", "Mexican Peso", "Mexico", 484},
	{"USD", "US Dollar", "United States", 840},
}

func find(name string) {
	for i := 0; i < 10; i++ {
		c := currencies[i]
		switch {
		case strings.Contains(c.Currency, name),
			strings.Contains(c.Name, name),
			strings.Contains(c.Country, name):
			fmt.Println("Found", c)
		}
	}
}

func findNumber(num int) {
	for _, curr := range currencies {
		if curr.Number == num {
			fmt.Println("Found", curr)
		}
	}
}

func findAny(val interface{}) {
	switch i := val.(type) {
	case int:
		findNumber(i)
	case string:
		find(i)
	default:
		fmt.Printf("Unable to search with type %T\n", val)
	}
}

func assertEuro(c Curr) bool {
	switch name, curr := "Euro", "EUR"; {
	case c.Name == name:
		return true
	case c.Currency == curr:
		return true
	}
	return false
}

func main() {
	find("Dinar")
	find("HTG")
	find("Hong Kong")

	findAny("Peso")
	findAny(404)
	findAny(978)
	findAny(false)

	curr1 := Curr{"EUR", "Euro", "Italy", 978}
	if assertEuro(curr1) {
		fmt.Println(curr1, "is Euro")
	}
}

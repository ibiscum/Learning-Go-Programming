package main

import "fmt"

// Using US volume units

type ounce float64

func (o ounce) cup() cup {
	return cup(o * 0.1250)
}

type cup float64

func (c cup) quart() quart {
	return quart(c * 0.25)
}

func (c cup) ounce() ounce {
	return ounce(c * 8.0)
}

type quart float64

func (q quart) gallon() gallon {
	return gallon(q * 0.25)
}

func (q quart) cup() cup {
	return cup(q * 4.0)
}

type gallon float64

func (g gallon) quart() quart {
	return quart(g * 4)
}

func main() {
	gal := gallon(5)
	fmt.Printf("%.2f gallons = %.2f quarts\n", gal, gal.quart())

	ozs := gal.quart().cup().ounce()
	fmt.Printf("%.2f gallons = %.2f ounces\n", gal, ozs)
	fmt.Printf("%.2f ounces = %.2f cups\n", ozs, ozs.cup())

	cps := ozs.cup()
	fmt.Printf("%.2f cups = %.2f quarts\n", cps, cps.quart())

	qrt := quart(5)
	fmt.Printf("%.2f quarts = %.2f cups\n", qrt, qrt.cup())
	fmt.Printf("%.2f quarts = %.2f gallons\n", qrt, qrt.gallon())
}

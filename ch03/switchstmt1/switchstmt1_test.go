package main

import (
	"testing"

	"gotest.tools/v3/assert"
)

func TestIsDollar(t *testing.T) {
	eur := Curr{"EUR", "Euro", "Italy", 978}
	got := isDollar(eur)
	assert.Assert(t, !got)

	dol := Curr{"AUD", "Australian Dollar", "Australia", 36}
	got = isDollar(dol)
	assert.Assert(t, got)
}

func TestIsDollar2(t *testing.T) {
	dol := Curr{"HKD", "Hong Kong Dollar", "Hong Koong", 344}
	got := isDollar2(dol)
	assert.Assert(t, got)

	eur := Curr{"EUR", "Euro", "Italy", 978}
	got = isDollar2(eur)
	assert.Assert(t, !got)
}

func TestIsEuro(t *testing.T) {
	dol := Curr{"HKD", "Hong Kong Dollar", "Hong Koong", 344}
	got := isEuro(dol)
	assert.Assert(t, !got)

	eur := Curr{"EUR", "Euro", "Italy", 978}
	got = isEuro(eur)
	assert.Assert(t, got)
}

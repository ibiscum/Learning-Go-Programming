package main

import (
	"testing"

	"gotest.tools/v3/assert"
)

func TestAssertEuro(t *testing.T) {
	eur := Curr{"EUR", "Euro", "Italy", 978}
	got := assertEuro(eur)
	assert.Assert(t, got)

	dol := Curr{"AUD", "Australian Dollar", "Australia", 36}
	got = assertEuro(dol)
	assert.Assert(t, !got)
}

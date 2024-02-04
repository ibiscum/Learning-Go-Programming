package main

import (
	"fmt"
	"strconv"
	"testing"

	"gotest.tools/v3/assert"
)

func TestAvg(t *testing.T) {
	a := fmt.Sprintf("%.2f", avg(1, 2.5, 3.75))
	b, _ := strconv.ParseFloat(a, 64)
	assert.Equal(t, b, 2.42)
}

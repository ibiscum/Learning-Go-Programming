package foo

import (
	"fmt"

	"github.com/ibiscum/Learning-Go-Programming/internal/foo/bazz"
)

func FooIt() {
	fmt.Println("Foo!")
	bazz.Qux()
}

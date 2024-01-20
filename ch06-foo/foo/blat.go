package foo

import (
	"fmt"

	"github.com/ibiscum/Learning-Go-Programming/ch06-foo/foo/bazz"
)

func FooIt() {
	fmt.Println("Foo!")
	bazz.Qux()
}

package demopkg

import (
	"fmt"
)

func init() {
	fmt.Print("Package init")
}
func Foo() {
	fmt.Print("Foo, say hello!!!")
}
func bar() {
	fmt.Print("bar")
}

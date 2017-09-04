package trees_test

import (
	"fmt"

	"github.com/baljanak/gogogo/trees"
)

func ExampleEquivalentTrees() {
	if trees.Same(trees.New(10, 1), trees.New(10, 2)) {
		fmt.Println("Trees are equivalent")
	} else {
		fmt.Println("Trees are different")
	}
	// Output:
	// Trees are different
}

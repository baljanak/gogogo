package trees_test

import (
	"fmt"

	"github.com/baljanak/gogogo/trees"
	"golang.org/x/tour/tree"
)

func ExampleEquivalentTrees() {
	if trees.Same(tree.New(1), tree.New(2)) {
		fmt.Println("Trees are equivalent")
	} else {
		fmt.Println("Trees are different")
	}
	// Output:
	// Trees are different
}

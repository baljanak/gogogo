package trees

import "math/rand"

// A Tree is a binary tree with integer values.
type Tree struct {
	Left  *Tree
	Value int
	Right *Tree
}

func insert(t *Tree, v int) *Tree {
	if t == nil {
		return &Tree{nil, v, nil}
	}
	if v < t.Value {
		t.Left = insert(t.Left, v)
		return t
	}
	t.Right = insert(t.Right, v)
	return t
}

// New returns a new, random binary tree
// holding the values 1k, 2k, ..., nk.
func New(n, k int) *Tree {
	var t *Tree
	for _, v := range rand.Perm(n) {
		t = insert(t, (1+v)*k)
	}
	return t
}

func walkFn(t *Tree, ch, quit chan int) {
	if t == nil {
		return
	}
	walkFn(t.Left, ch, quit)
	select {
	case ch <- t.Value:
	case <-quit:
		return
	}
	walkFn(t.Right, ch, quit)
}

// Walk walks the tree t sending all values
// from the tree to the channel ch.
func Walk(t *Tree, ch, quit chan int) {
	walkFn(t, ch, quit)
	close(ch)
}

// Same determines whether the trees
// t1 and t2 contain the same values.
func Same(t1, t2 *Tree) bool {
	ch1, ch2 := make(chan int), make(chan int)
	quit := make(chan int)

	defer close(quit)

	go Walk(t1, ch1, quit)
	go Walk(t2, ch2, quit)

	for {
		v1, ok1 := <-ch1
		v2, ok2 := <-ch2

		if !ok1 || !ok2 {
			return ok1 == ok2
		}

		if v1 != v2 {
			return false
		}
	}
}

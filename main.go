package main

import (
	"fmt"

	"golang.org/x/tour/tree"
)

func walkRecur(t *tree.Tree, c chan int) {
	if t.Left != nil {
		walkRecur(t.Left, c)
	}
	c <- t.Value
	if t.Right != nil {
		walkRecur(t.Right, c)
	}
}

func Walk(t *tree.Tree) chan int {
	c := make(chan int)
	if t == nil {
		close(c)
	} else {
		go func() {
			walkRecur(t, c)
			close(c)
		}()
	}
	return c
}

func Same(t1, t2 *tree.Tree) bool {
	if t1 == t2 {
		return true
	}
	c1 := Walk(t1)
	c2 := Walk(t2)
	for {
		v1, ok1 := <-c1
		v2, ok2 := <-c2
		if ok1 != ok2 || v1 != v2 {
			return false
		} else if ok1 && v1 == v2 {
			continue
		} else {
			break
		}
	}
	return true
}

func main() {
	t1, t2, t3 := tree.New(1), tree.New(1), tree.New(2)
	fmt.Println(Same(t1, t1))
	fmt.Println(Same(t1, t2))
	fmt.Println(Same(t1, t3))
}

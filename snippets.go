package main

import (
	"fmt"
	"golang.org/x/tour/tree"
) // Walk walks the tree t sending all values
// from the tree to the channel ch.
func Walk(t *tree.Tree, ch chan int) {
	if t == nil {
		return
	}
	Walk(t.Left, ch)
	fmt.Println(t.Value)
	ch <- t.Value
	Walk(t.Right, ch)
}

func Walker(t *tree.Tree, ch chan int) {
	Walk(t, ch)
	close(ch)
}

// Same determines whether the trees
// t1 and t2 contain the same values.
func Same(t1, t2 *tree.Tree) bool {
	ch1 := make(chan int)
	ch2 := make(chan int)

	go Walker(t1, ch1)
	go Walker(t2, ch2)

	for {
		v1, ok := <-ch1
		v2, ok2 := <-ch2
		if ok != ok2 || v1 != v2 {
			return false
		}
		if !ok {
			break
		}
	}
	return true
}

func Map[T any, U any](slice []T, f func(T) U) []U {
	result := make([]U, len(slice))
	for index, current := range slice {
		result[index] = f(current)
	}
	return result
}

//Pairs DNA chains
func PairElement(str string) []string {
	dnaMap := map[string]string{
		"a": "at",
		"t": "ta",
		"g": "gc",
		"c": "cg",
	}
	pairedElements := []string{}
	for _, char := range strings.Split(str, "") {
		if val, ok := dnaMap[strings.ToLower(char)]; ok {
			pairedElements = append(pairedElements, val)
		}
	}
	return pairedElements
}

func main() {
	fmt.Println(Same(tree.New(3), tree.New(3)))
}

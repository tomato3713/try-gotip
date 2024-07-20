package main

import (
	"fmt"
	"iter"
)

func generator(n int) iter.Seq[int] {
	return func(yield func(n int) bool) {
		for v := range n {
			yield(v + 1)
		}
	}
}

func main() {
	for v := range generator(10) {
		fmt.Println(v)
	}
}

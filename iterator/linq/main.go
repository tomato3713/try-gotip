package main

import (
	"fmt"
	"iter"
)

type linq[T any] iter.Seq[T]

func From[T any](list []T) linq[T] {
	return func(yield func(T) bool) {
		for _, item := range list {
			if !yield(item) {
				return
			}
		}
	}
}

func (f linq[T]) Where(eq func(x T) bool) linq[T] {
	return func(yield func(x T) bool) {
		for v := range f {
			if !eq(v) {
				continue
			}

			if !yield(v) {
				return
			}
		}
	}
}

func main() {
	list := []int{5, 3, 10, 1, 33, 4, 12, 11, 15}

	cond1 := func(v int) bool { return v > 5 }
	cond2 := func(v int) bool { return v < 30 }

	for v := range From(list).Where(cond1).Where(cond2) {
		fmt.Println(v)
	}
}

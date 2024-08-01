package main

import (
	"fmt"
	"iter"
)

func From[T any](list []T) iter.Seq[T] {
	return func(yield func(v T) bool) {
		for _, item := range list {
			if !yield(item) {
				return
			}
		}
	}
}

func Where[T comparable](f iter.Seq[T], eq func(x T) bool) iter.Seq[T] {
	return func(yield func(T) bool) {
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

	for v := range Where(Where(From(list), cond1), cond2) {
		fmt.Println(v)
	}
}

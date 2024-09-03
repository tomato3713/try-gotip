package main

import (
	"fmt"
)

// 1度繰り返すだけのイテレータ
func seq1(yield func() bool) {
	yield()
}

// list の値を順に返すイテレータ
func seq2(yield func(v string) bool) {
	for _, v := range []string{"apple", "orange", "lemon"} {
		yield(v)
	}
}

// list のインデックスと値を順に返すイテレータ
func seq3(yield func(k int, v string) bool) {
	list := []string{"apple", "orange", "lemon"}
	for k, v := range list {
		yield(k, v)
	}
}

// list のインデックスと値を順に返すイテレータ
func seq4(yield func(k int, v string) bool) {
	list := []string{"apple", "orange", "lemon"}
	for k, v := range list {
		yield(k, v)
	}
}

// list のインデックスと値を順に返すイテレータ
func seq5(yield func(k int, v string) bool) {
	list := []string{"apple", "orange", "lemon"}
	for k, v := range list {
		if !yield(k, v) {
			return
		}
	}
}

func main() {
	for range seq1 {
		fmt.Println("hello")
	}

	for v := range seq2 {
		fmt.Println(v)
	}

	for k, v := range seq3 {
		fmt.Println(k, v)
	}

	fmt.Println("break")

	for k, v := range seq5 {
		fmt.Println(k, v)
		if k > 0 {
			break
		}
	}
}

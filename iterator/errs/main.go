/*
失敗したものだけ返すイテレーターがあればリトライが便利に書けないかお試し
*/
package main

import (
	"errors"
	"fmt"
	"iter"
)

type Item string
type List []Item

func (i Item) String() string {
	return string(i)
}

func something(item Item) error {
	if len(item) <= 3 {
		return errors.New("failed")
	}
	return nil
}

func procs(arr List) iter.Seq2[Item, error] {
	return func(yield func(_ Item, v error) bool) {
		for _, item := range arr {
			err := something(item)
			if !yield(item, err) {
				return
			}
		}
	}
}

func All(arr List) List {
	var failedList List
	for key, err := range procs(arr) {
		if err != nil {
			fmt.Printf("err = %v, item = %v\n", err, key)
			failedList = append(failedList, key)
		}
	}

	return failedList
}

func main() {
	arr := List{
		"aaaaa",
		"bbbbb",
		"ccccc",
		"dddd",
		"eee",
		"ff",
		"d",
	}

	failedList := All(arr)
	if len(failedList) != 0 {
		// retry process...
		fmt.Println("finished to retry")
	}

	fmt.Printf("finally failed list is ... [")
	for i, v := range failedList {
		if i == len(failedList)-1 {
			fmt.Printf("%v]\n", v)
		} else {
			fmt.Printf("%v, ", v)
		}
	}
}

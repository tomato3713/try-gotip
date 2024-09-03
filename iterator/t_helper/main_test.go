package main_test

import (
	"iter"
	"slices"
	"testing"
)

func testSlice(t *testing.T, it []string, data, wants string) {
	t.Helper()

	for _, s := range it {
		if s != wants {
			t.Error("somthing wrong!")
		}
	}
}

func testIter(t *testing.T, it iter.Seq[string], data, wants string) {
	t.Helper()

	for s := range it {
		if s != wants {
			// 本来なら slice を使った testSlice と同様に呼び出し元の L35 でエラーが報告されてほしいが L25 でのエラーとして報告される
			t.Error("somthing wrong!")
		}
	}
}

func TestSomething(t *testing.T) {
	list := []string{"orange", "apple", "melon"}
	var it iter.Seq[string] = slices.Values(list)

	t.Log("use iterator")
	testIter(t, it, "data 1 ...", "wants 1 ...")

	t.Log("use slice")
	testSlice(t, list, "data 1 ...", "wants 1 ...")

}

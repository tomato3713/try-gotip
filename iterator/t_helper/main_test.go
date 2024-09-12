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

func check(t *testing.T, got, want string) {
	if got != want {
		t.Error("somthing wrong!")
	}
}

func testSlice2(t *testing.T, it []string, data, wants string) {
	t.Helper()

	for _, s := range it {
		check(t, s, wants)
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

func testIter2(t *testing.T, it iter.Seq[string], data, wants string) {
	t.Helper()

	for s := range it {
		// うまくいかない
		// iter.go:40 になった
		t.Helper()
		if s != wants {
			t.Error("somthing wrong!")
		}
	}
}

func testIterToSlice(t *testing.T, it iter.Seq[string], data, wants string) {
	t.Helper()

	for _, s := range slices.Collect(it) {
		if s != wants {
			t.Error("somthing wrong!")
		}
	}
}

func TestSomething(t *testing.T) {
	list := []string{"orange", "apple", "melon"}
	var it iter.Seq[string] = slices.Values(list)

	t.Log("use iterator")
	// range-loop のイテレータは range over func を使わない形に展開されるのでスタックが積まれてループ内が失敗場所として報告される
	testIter(t, it, "data 1 ...", "wants 1 ...")

	t.Log("use iterator2")
	testIter2(t, it, "data 1 ...", "wants 1 ...")

	t.Log("use iterator to slice")
	// スライスに変換してループをまわしているので t.Helper が効く
	testIterToSlice(t, it, "data 1 ...", "wants 1 ...")

	t.Log("use slice")
	// スライスなら t.Helper が効く
	testSlice(t, list, "data 1 ...", "wants 1 ...")

	t.Log("use slice2")
	// スライスでループしているが、追加で呼びだした関数内で失敗するので失敗位置の報告が呼び出し先になる
	// terIter() と同じような状況
	testSlice2(t, list, "data 1 ...", "wants 1 ...")
}

func checkHelper(t *testing.T, got iter.Seq2[int, int], wants []int) {
	t.Helper()

	for i, s := range got {
		if s != wants[i] {
			// 本来なら slice を使った testSlice と同様に呼び出し元の L35 でエラーが報告されてほしいが L25 でのエラーとして報告される
			t.Error("somthing wrong!")
		}
	}
}

func TestA(t *testing.T) {
	it := slices.All([]int{1, 2, 3})

	// range-loop のイテレータは range over func を使わない形に展開されるのでスタックが積まれてループ内が失敗場所として報告される
	checkHelper(t, it, []int{1, 2, 5})
}

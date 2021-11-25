package lister

import (
	"math/rand"
	"testing"
)

var my_list = New(2_000_000)
var skip_list = NewSkipList(2_000_000)

func BenchmarkList(b *testing.B) {
	for i := 0; i < b.N; i++ {
		my_list.Search(r())
	}
}

func BenchmarkBinaryList(b *testing.B) {
	for i := 0; i < b.N; i++ {
		my_list.BinarySearch(r())
	}
}

func BenchmarkSkipList(b *testing.B) {
	for i := 0; i < b.N; i++ {
		skip_list.Find(Element(r()))
	}
}

func r() int {
	min := 10000
	max := 1000000
	return rand.Intn(max-min) + min
}

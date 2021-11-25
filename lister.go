package lister

import (
	"fmt"
	"github.com/MauriceGit/skiplist"
)

type lister struct {
	data []int
}

type Element int

func (e Element) ExtractKey() float64 {
	return float64(e)
}

func (e Element) String() string {
	return fmt.Sprintf("%03d", e)
}

func New(size int) *lister {
	var d []int
	for i := 0; i < size; i++ {
		d = append(d, i)
	}
	return &lister{data: d}
}

func NewSkipList(size int) skiplist.SkipList {
	s := skiplist.New()
	for i := 0; i < size; i++ {
		s.Insert(Element(i))
	}
	return s
}

func (l lister) Search(k int) int {
	for i := range l.data {
		if i == k {
			return k
		}
	}
	return -1
}

func (l lister) BinarySearch(k int) int {
	low := 0
	high := len(l.data) - 1

	for low <= high {
		median := (low + high) / 2

		if l.data[median] < k {
			low = median + 1
		} else {
			high = median - 1
		}
	}

	if low == len(l.data) || l.data[low] != k {
		return -1
	}

	return k
}

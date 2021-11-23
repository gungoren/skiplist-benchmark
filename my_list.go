package mylist

import (
	"fmt"
	"github.com/MauriceGit/skiplist"
)

type mylist struct {
	data []int
}

type Element int

func (e Element) ExtractKey() float64 {
	return float64(e)
}
func (e Element) String() string {
	return fmt.Sprintf("%03d", e)
}

func New(size int) *mylist {
	var d []int

	for i := 0; i < size; i++ {
		d = append(d, i)
	}

	return &mylist{data: d}
}

func NewSkipList(size int) skiplist.SkipList {
	s := skiplist.New()
	for i := 0; i < size; i++ {
		s.Insert(Element(i))
	}
	return s
}

func (m mylist) Search(k int) int {
	for i := range m.data {
		if i == k {
			return k
		}
	}
	return -1
}

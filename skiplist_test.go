package mylist

import (
	"fmt"
	"testing"
	"time"
)

func TestList(t *testing.T) {

	mylist := New(20000000)
	now := time.Now()
	mylist.Search(100000)
	fmt.Println(time.Since(now))

}

func TestSkipList(t *testing.T) {
	s := NewSkipList(20000000)
	now := time.Now()
	s.Find(Element(100000))
	fmt.Println(time.Since(now))

}

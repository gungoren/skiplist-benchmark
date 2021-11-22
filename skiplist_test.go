package skiplist

import (
	"bytes"
	"testing"
)

func TestSkipList_Search(t *testing.T) {
	s := New()
	s.Insert([]byte("john"), []byte("doe"))
	s.Insert([]byte("jack"), []byte("john"))
	s.Insert([]byte("rocky"), []byte("mathew"))
	s.Insert([]byte("walley"), []byte("richard"))
	s.Insert([]byte("wolfram"), []byte("matthew"))

	val := s.Search([]byte("wolfram"))
	if bytes.Compare(val.value, []byte("matthew")) != 0 {
		t.Errorf("failed")
	}
}

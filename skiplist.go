package skiplist

import (
	"bytes"
	"math/rand"
	"time"
)

const maxLevel int = 20

type (

	Element struct {
		right *Element
		down *Element
		key   []byte
		value []byte
	}

	SkipList struct {
		head *Element
		maxLevel int
		rand     *rand.Rand
	}
)

func (e *Element) equals(element Element) bool {
	return bytes.Compare(e.key, element.key) == 0
}

func (e *Element) byteEquals(key []byte) bool {
	return bytes.Compare(e.key, key) == 0
}

func (e *Element) byteGreaterThan(key []byte) bool {
	return bytes.Compare(e.key, key) > 0
}

func New() *SkipList {
	return &SkipList{
		maxLevel:  maxLevel,
		rand:      rand.New(rand.NewSource(time.Now().UnixNano())),
	}
}

func (s *SkipList) Search(key []byte) *Element {
	element := s.head
	for ; element != nil; {
		if element.byteEquals(key) {
			return element
		} else if element.right == nil || element.right.byteGreaterThan(key) {
			element = element.down
		} else {
			element = element.right
		}
	}
	return nil
}

func (s *SkipList) Insert(key, value []byte) {
	element := s.Search(key)
	if element != nil{
		element.value = value
		return
	}
	if s.head == nil {
		s.head = &Element{
			key:   key,
			value: value,
		}
		return
	}
	var stack []*Element
	current := s.head
	for ; current!= nil; {
		if current.right == nil || current.right.byteGreaterThan(key) {
			stack = append(stack, current)
			current = current.down
		} else {
			stack = nil
			stack = append(stack, current)
			current = current.right
		}
	}
	current = stack[len(stack)-1]
	level := s.randomLevel()
	var downElement *Element
	for i := 0; i < len(stack); i++  {
		if i >= level {
			break
		}
		rightElement := stack[len(stack) - 1 -i].right
		currentElement := &Element{
			right: rightElement,
			down:  downElement,
			key:   key,
			value: value,
		}
		s.head.right = currentElement
		downElement = currentElement
	}
	s.head = s.head.right
	return
}

func (s *SkipList) randomLevel() int {
	const probability int32 = 1 << 30

	i := 1

	for ; i < s.maxLevel; i++ {
		if s.rand.Int31() < probability {
			break
		}
	}
	return i

}
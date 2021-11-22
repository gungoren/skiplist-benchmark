package skiplist

import (
	"bytes"
	"math/rand"
	"time"
)

const maxLevel int = 20

type (
	Element struct {
		next  []*Element
		prev  *Element
		key   []byte
		value []byte
		level int
	}

	SkipList struct {
		headLevels []*Element
		tailLevels []*Element
		maxLevel   int
		size       int
		rand       *rand.Rand
	}
)

func New() *SkipList {
	return &SkipList{
		headLevels: make([]*Element, maxLevel),
		tailLevels: make([]*Element, maxLevel),
		maxLevel:   maxLevel,
		rand:       rand.New(rand.NewSource(time.Now().UnixNano())),
	}
}

func (s *SkipList) Search(key []byte) *Element {
	currentIndex := s.findEntryPointIndex(key, 0)
	current := s.headLevels[currentIndex]
	if current != nil && current.equals(key) {
		return current
	}

	for {
		next := current.next[currentIndex]

		if next != nil && next.lessThan(key) {
			current = next
		} else {
			if currentIndex > 0 {
				if current.next[0] != nil && current.next[0].equals(key) {
					return current.next[0]
				}
				currentIndex--
			} else {
				return next
			}
		}
	}
}

func (s *SkipList) Insert(key, value []byte) {
	element := s.Search(key)
	if element != nil {
		element.value = value
		return
	}
	level := s.randomLevel()

	newElement := &Element{
		next:  make([]*Element, s.maxLevel),
		key:   key,
		value: value,
		level: level,
	}
	index := s.findEntryPointIndex(key, level)
	var current *Element
	var next *Element
	for {
		if current == nil {
			next = s.headLevels[index]
		} else {
			next = current.next[index]
		}

		if index <= level && (next == nil || next.greaterThan(key)) {
			newElement.next[index] = next
			if current != nil {
				current.next[index] = newElement
			}

			if index == 0 {
				newElement.prev = current
				if next != nil {
					next.prev = newElement
				}
			}
		}

		if next != nil && next.lessThan(key) {
			current = next
		} else {
			index--
			if index < 0 {
				break
			}
		}
	}

}

func (s SkipList) findEntryPointIndex(key []byte, level int) int {
	for i := s.maxLevel - 1; i >= 0; i-- {
		if s.headLevels[i] != nil && !s.headLevels[i].greaterThan(key) || i < level {
			return i
		}
	}
	return 0
}

func (e *Element) equals(key []byte) bool {
	return bytes.Compare(e.key, key) == 0
}

func (e *Element) greaterThan(key []byte) bool {
	return bytes.Compare(e.key, key) > 0
}

func (e *Element) lessThan(key []byte) bool {
	return bytes.Compare(e.key, key) < 0
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

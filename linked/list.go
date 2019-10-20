package linked

import (
	"errors"
	"fmt"
)

// List is the linked list structure
type List struct {
	start *node
	end   *node
	size  uint
}

// NewList creates a new list
func NewList() *List {
	return &List{}
}

// Retrieve will return the data at a specific idx
func (l *List) Retrieve(idx uint) (interface{}, error) {
	if idx >= l.size {
		return nil, fmt.Errorf("linked list: index [%d] must be less than size [%d]", idx, l.size)
	}
	switch {
	case idx == 0:
		return l.start.data, nil
	case idx == (l.size - 1):
		return l.end.data, nil
	default:
		curr := l.start
		for i := uint(0); i < idx; i++ {
			curr = curr.next
		}
		return curr.data, nil
	}
}

// Insert data to in a specific location
func (l *List) Insert(data interface{}, idx uint) error {
	if data == nil {
		return errors.New("linked list: unable to append nil data")
	}
	n := &node{
		data: data,
	}
	switch {
	case idx > l.size:
		return errors.New("linked list: idx out of bounds")
	case idx == l.size:
		return l.Append(data)
	case idx == 0:
		n.next = l.start
		l.start = n
	default:
		curr := l.start
		prev := l.start
		for i := uint(0); i < idx; i++ {
			prev = curr
			curr = curr.next
		}
		prev.next = n
		n.next = curr
	}
	l.size++
	return nil
}

// Append is place a linked node at the end of the list
func (l *List) Append(data interface{}) error {
	if data == nil {
		return errors.New("linked list: unable to append nil data")
	}
	n := &node{
		data: data,
	}
	switch {
	case l.start == nil:
		l.start = n
		l.end = n
	default:
		l.end.next = n
		l.end = n
	}
	l.size++
	return nil
}

// Delete will remove a node from the list
func (l *List) Delete(idx uint) error {
	if idx >= l.size {
		return fmt.Errorf("linked list: index [%d] must be less than size [%d]", idx, l.size)
	}
	switch {
	case idx == 0:
		l.start = l.start.next
		if l.start == nil {
			l.end = nil
		}
	default:
		curr := l.start
		prev := l.start
		for idx > 0 {
			prev = curr
			curr = curr.next
			idx--
		}
		prev.next = curr.next
		if curr == l.end {
			l.end = prev
		}
	}
	l.size--
	return nil
}

// Size is the total number of nodes
func (l *List) Size() uint {
	return l.size
}

// ForEach will iterate through the nodes
func (l *List) ForEach(f func(interface{}, uint) error) error {
	curr := l.start
	idx := uint(0)
	for curr != nil {
		if err := f(curr.data, idx); err != nil {
			return err
		}
		idx++
		curr = curr.next
	}
	return nil
}

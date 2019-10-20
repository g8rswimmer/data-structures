package linked

import (
	"errors"
	"fmt"
)

type node struct {
	data interface{}
	next *node
}

// List is the linked list structure
type List struct {
	start *node
	size  int
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
	default:
		curr := l.start
		for curr.next != nil {
			curr = curr.next
		}
		curr.next = n
	}
	l.size++
	return nil
}

// Delete will remove a node from the list
func (l *List) Delete(idx int) error {
	if idx >= l.size {
		return fmt.Errorf("linked list: index [%d] must be less than size [%d]", idx, l.size)
	}
	switch {
	case idx == 0:
		l.start = l.start.next
	default:
		curr := l.start
		prev := l.start
		for idx > 0 {
			prev = curr
			curr = curr.next
			idx--
		}
		prev.next = curr.next
	}
	l.size--
	return nil
}

// Size is the total number of nodes
func (l *List) Size() int {
	return l.size
}

// ForEach will iterate through the nodes
func (l *List) ForEach(f func(interface{}, int) error) error {
	curr := l.start
	idx := 0
	for curr != nil {
		if err := f(curr.data, idx); err != nil {
			return err
		}
		idx++
		curr = curr.next
	}
	return nil
}

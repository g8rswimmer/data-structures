package linked

import "errors"

// Stack is a linked list stack implementation
type Stack struct {
	list *List
}

// Push will place data at the top of the stack
func (s *Stack) Push(data interface{}) error {
	if data == nil {
		return errors.New("linked stack: the data can not be nil")
	}
	n := &node{
		data: data,
		next: s.list.start,
	}
	s.list.start = n
	s.list.size++
	return nil
}

// Pop will remove data from the top of the stack
func (s *Stack) Pop() (interface{}, error) {

	data, err := s.Peek()
	if err != nil {
		return nil, err
	}
	return data, s.list.Delete(0)
}

// Peek retuens the data from the top of the stack, does not remove it.
func (s *Stack) Peek() (interface{}, error) {
	if s.list.size == 0 {
		return nil, errors.New("linked stack: the stack is empty")
	}
	return s.list.start.data, nil
}

// Size returns the size of the stack
func (s *Stack) Size() int {
	return s.list.size
}

package linked

// Stack is a linked list stack implementation
type Stack struct {
	list *List
}

// NewStack create a new linked stack.
func NewStack() *Stack {
	return &Stack{
		list: NewList(),
	}
}

// Push will place data at the top of the stack
func (s *Stack) Push(data interface{}) error {
	return s.list.Insert(data, 0)
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
	return s.list.Retrieve(0)
}

// Size returns the size of the stack
func (s *Stack) Size() uint {
	return s.list.size
}

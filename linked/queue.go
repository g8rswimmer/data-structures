package linked

// Queue is a linked list implementation
type Queue struct {
	list *List
}

// NewQueue create a new Queue
func NewQueue() *Queue {
	return &Queue{
		list: NewList(),
	}
}

// Enqueue adds the data at the end of the queue
func (q *Queue) Enqueue(data interface{}) error {
	return q.list.Append(data)
}

// Dequeue removes the data from the start of the queue
func (q *Queue) Dequeue() (interface{}, error) {
	data, err := q.list.Retrieve(0)
	if err != nil {
		return nil, err
	}
	if err := q.list.Delete(0); err != nil {
		return nil, err
	}
	return data, nil
}

// Size is the length of the queue
func (q *Queue) Size() uint {
	return q.list.size
}

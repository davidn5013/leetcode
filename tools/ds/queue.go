package ds

// Queue storing int
type Queue struct {
	items []int
}

// NewMyQueue Will create a new MyQueue
func NewMyQueue() *Queue {
	return &Queue{
		items: []int{},
	}
}

// Push Pushes element x to the back of the queue.
func (q *Queue) Push(i int) {
	q.items = append(q.items, i)
}

// Pop Removes the element from the front of the queue and returns it.
func (q *Queue) Pop() (res int) {
	res = q.items[0]
	q.items = q.items[1:]
	return res
}

// Peek Returns the element at the front of the queue.
func (q *Queue) Peek() int {
	if len(q.items) > 0 {
		return q.items[0]
	}
	return 0
}

// Empty Returns true if the queue is empty, false otherwise.
func (q *Queue) Empty() bool {
	return len(q.items) == 0
}

/**
 * Your MyQueue object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Push(x);
 * param_2 := obj.Pop();
 * param_3 := obj.Peek();
 * param_4 := obj.Empty();
 */

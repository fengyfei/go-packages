package queue

// Queue represents a circular queue.
type Queue struct {
	data   []interface{}
	head   int
	tail   int
	rooms  int
	growth int
}

// NewQueue creates a queue with a initial size.
func NewQueue(size int) *Queue {
	if size <= 0 {
		size = 32
	}

	q := &Queue{
		growth: size,
	}

	q.grow()

	return q
}

func (q *Queue) grow() {
	growth := q.growth
	data := make([]interface{}, len(q.data)+growth)

	// the only condition is queue is out of rooms
	pos := copy(data, q.data[q.tail:])
	copy(data[pos:], q.data[:q.head])

	q.head = len(q.data)
	q.tail = 0
	q.rooms = growth
	q.data = data
}

// Put a element on the front of the queue.
func (q *Queue) Put(v interface{}) {
	if v == nil {
		return
	}

	if q.rooms == 0 {
		q.grow()
	}

	q.data[q.head] = v
	q.head = (q.head + 1) % len(q.data)
	q.rooms--
}

// Get a element on the end of the queue.
func (q *Queue) Get() interface{} {
	if q.head == q.tail && q.rooms > 0 {
		return nil
	}

	v := q.data[q.tail]
	q.tail = (q.tail + 1) % len(q.data)
	q.rooms++

	return v
}

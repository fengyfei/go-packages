package bytes

func round8(v int) int {
	const r = 7

	return (v + r) & (^r)
}

// RingBuffer is a general purpose byte buffer.
// RingBuffer is not thread-safe, it's your responsibility for protect the ring buffer instance.
type RingBuffer struct {
	readOff  int
	writeOff int
	count    int
	cap      int
	data     []byte
}

// NewRingBuffer creates a ring-buffer instance with at least size capacity.
func NewRingBuffer(cap int) *RingBuffer {
	if cap == 0 {
		cap = 7
	}
	cap = round8(cap)

	return &RingBuffer{
		cap:  cap,
		data: make([]byte, cap),
	}
}

// Length returns the number of byte stored in the buffer.
func (rb *RingBuffer) Length() int {
	return rb.count
}

// Cap returns the capacity of the underlying byte slice.
func (rb *RingBuffer) Cap() int {
	return rb.cap
}

func (rb *RingBuffer) grow(n int) error {
	capacity := rb.cap + round8(n)

	buf, err := safeMakeByteSlice(capacity)
	if err != nil {
		return err
	}

	// copy
	if rb.writeOff < rb.readOff {
		pos := copy(buf, rb.data[rb.readOff:])
		copy(buf[pos:], rb.data[:rb.writeOff])
	} else {
		copy(buf, rb.data[rb.readOff:rb.writeOff])
	}

	rb.readOff = 0
	rb.writeOff = rb.count
	rb.cap = capacity
	rb.data = buf

	return nil
}

// Write implements the io.Writer interface.
func (rb *RingBuffer) Write(p []byte) (int, error) {
	len := len(p)
	expected := len

	if len > rb.cap-rb.count {
		if err := rb.grow(len + rb.count - rb.cap); err != nil {
			expected = rb.cap - rb.count
		}
	}

	n := copy(rb.data[rb.writeOff:], p[:expected])
	if n < expected {
		n += copy(rb.data[:rb.readOff], p[n:expected])
	}

	rb.count += n
	rb.writeOff = (rb.writeOff + n) % rb.cap

	return n, nil
}

// Read implements the io.Reader interface
func (rb *RingBuffer) Read(p []byte) (int, error) {
	expected := len(p)
	n := 0

	if rb.writeOff > rb.readOff {
		n += copy(p, rb.data[rb.readOff:rb.writeOff])
	} else {
		pos := copy(p, rb.data[rb.readOff:])
		n += pos

		if n < expected {
			n += copy(p[pos:], rb.data[:rb.writeOff])
		}
	}

	rb.count -= n
	rb.readOff += (rb.readOff + n) % rb.cap

	return n, nil
}

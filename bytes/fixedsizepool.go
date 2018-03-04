package bytes

import (
	"sync"
)

// FixedSizeBuffer is a list in which each element contains a fixed length byte slice.
type FixedSizeBuffer struct {
	buf  []byte
	next *FixedSizeBuffer
}

// Bytes retures the underlying byte slice.
func (fb *FixedSizeBuffer) Bytes() []byte {
	return fb.buf
}

// Length returns the length of the buffer.
func (fb *FixedSizeBuffer) Length() int {
	return len(fb.buf)
}

// Cap returns the capacity of the buffer, not same with cap(buf).
func (fb *FixedSizeBuffer) Cap() int {
	return len(fb.buf)
}

// FixedSizeBufferPool is a thread safe pool contains a FixedSizeBuffer list.
type FixedSizeBufferPool struct {
	mu   sync.Mutex
	free *FixedSizeBuffer

	counts int // initial counts of FixedSizeBuffer, read-only
	size   int // size of each FixedSizeBuffer, read-only
	avail  int // counts of FixedSizeBuffer available to Get
}

// NewFixedSizeBufferPool creates a instance contains counts elements with
// each have a byte slice length of size.
func NewFixedSizeBufferPool(counts, size int) *FixedSizeBufferPool {
	if counts <= 0 {
		counts = 16
	}

	if size <= 0 {
		size = 128
	}

	p := &FixedSizeBufferPool{
		counts: counts,
		size:   size,
	}
	p.grow(counts)

	return p
}

func (fbp *FixedSizeBufferPool) grow(counts int) (err error) {
	size := fbp.size

	defer func() {
		if recover() != nil {
			err = errOutOfMemory
		}
	}()

	p := make([]byte, counts*size)
	fsb := make([]FixedSizeBuffer, counts)

	// make a flat memory map
	for i := counts; i > 0; i-- {
		fsb[i-1] = FixedSizeBuffer{
			buf: p[(i-1)*size : i*size],
		}

		fsb[i-1].next = fbp.free
		fbp.free = &fsb[i-1]
	}

	fbp.avail += counts
	return
}

// Get a buffer from the pool, if no buffer left, grow the pool with half of the counts.
func (fbp *FixedSizeBufferPool) Get() *FixedSizeBuffer {
	fbp.mu.Lock()

	if fbp.avail == 0 {
		if err := fbp.grow(fbp.counts >> 1); err != nil {
			fbp.mu.Unlock()
			return nil
		}
	}

	b := fbp.free
	fbp.free = b.next
	fbp.avail--

	fbp.mu.Unlock()

	return b
}

// Put back a buffer.
func (fbp *FixedSizeBufferPool) Put(b *FixedSizeBuffer) {
	fbp.mu.Lock()

	b.next = fbp.free
	fbp.free = b

	fbp.avail++

	fbp.mu.Unlock()
}

// Available returns the counts of buffer left in the pool.
func (fbp *FixedSizeBufferPool) Available() int {
	fbp.mu.Lock()
	avail := fbp.avail
	fbp.mu.Unlock()

	return avail
}

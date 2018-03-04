package bytes

import (
	"testing"
)

func TestFixedSizeBufferPool(t *testing.T) {
	counts := []int{4, 6}
	size := 128
	pool := NewFixedSizeBufferPool(counts[0], size)

	for loop := 0; loop < 2; loop++ {
		fbs := make([]*FixedSizeBuffer, counts[loop])

		for i := 0; i < counts[loop]; i++ {
			fbs[i] = pool.Get()

			if fbs[i].Length() != size {
				t.Fatalf("fixedSizeBufferPool %d Get failed", i)
			}

			if fbs[i].Cap() != size {
				t.Fatalf("fixedSizeBufferPool %d Get failed, cap: %d", i, fbs[i].Cap())
			}
		}

		if pool.Available() != 0 {
			t.Fatalf("fixedSizeBufferPool available expected 0, but %d", pool.Available())
		}

		another := pool.Get()
		if another.Bytes() == nil || len(another.Bytes()) == 0 {
			t.Fatalf("fixedSizeBufferPool grow failed")
		}

		for i := 0; i < counts[loop]; i++ {
			pool.Put(fbs[i])
		}
		pool.Put(another)

		if pool.Available() != counts[0]*1+counts[0]*(loop+1)/2 {
			t.Fatalf("fixedSizeBufferPool loop %d available is %d", loop+1, pool.Available())
		}
	}
}

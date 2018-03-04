package bytes

import (
	"fmt"
	"testing"
)

/*func TestRound8(t *testing.T) {
	data := []struct {
		v int
		e int
	}{
		{v: 0, e: 0},
		{v: 1, e: 8}, {v: 2, e: 8}, {v: 7, e: 8},
		{v: 9, e: 16}, {v: 15, e: 16},
	}

	for _, d := range data {
		v := round8(d.v)

		if v != d.e {
			t.Fatalf("round8 failed for %v, expected: %v, actual: %v", d.v, d.e, v)
		}
	}
}*/

func TestWrite(t *testing.T) {
	rb := NewRingBuffer(7)

	if rb.Cap() != 8 {
		t.Fatal("write: cap(rb) is expected 8")
	}

	n, err := rb.Write([]byte("abcdefg"))
	if err != nil {
		t.Fatal("write: write failed:", err)
	}
	if n != 7 {
		t.Fatalf("write: invalid size %d, 7 expected", n)
	}

	n, err = rb.Write([]byte("efghi"))
	if err != nil {
		t.Fatal("write: write failed:", err)
	}
	if n != 5 {
		t.Fatalf("write: invalid size %d, 5 expected, buf: %v", n, rb.data)
	}

	fmt.Println(rb.data)
}

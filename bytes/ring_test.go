package bytes

import (
	"bytes"
	"testing"
)

func TestRound8(t *testing.T) {
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
}

func TestRingBuffer(t *testing.T) {
	rb := NewRingBuffer(7)
	buf := make([]byte, 3)

	if rb.Cap() != 8 {
		t.Fatal("write: cap(rb) is expected 8")
	}

	// write 8
	n, err := rb.Write([]byte("01234567"))
	if err != nil {
		t.Fatalf("write: write failed: %v, write: %d", err, n)
	}

	// read 3, left 5
	if n, err = rb.Read(buf); err != nil {
		t.Fatalf("read %d bytes error: %v", n, err)
	}
	if !bytes.Equal(buf, []byte("012")) {
		t.Fatalf("read error: %v", buf)
	}

	// read 3, left 2
	if n, err = rb.Read(buf); err != nil {
		t.Fatalf("read %d bytes error: %v", n, err)
	}
	if !bytes.Equal(buf, []byte("345")) {
		t.Fatalf("read error: %v", buf)
	}

	// write 5, left 7
	if n, err = rb.Write([]byte("89012")); err != nil {
		t.Fatal("write: write failed:", err)
	}

	// write 2, realloc, left 9
	if n, err = rb.Write([]byte("34")); err != nil {
		t.Fatal("write: write failed:", err)
	}

	// read 3, left 6
	if n, err = rb.Read(buf); err != nil {
		t.Fatalf("read %d bytes error: %v", n, err)
	}
	if !bytes.Equal(buf, []byte("678")) {
		t.Fatalf("read error: %v", buf)
	}

	// read 3, left 3
	if n, err = rb.Read(buf); err != nil {
		t.Fatalf("read %d bytes error: %v", n, err)
	}
	if !bytes.Equal(buf, []byte("901")) {
		t.Fatalf("read error: %v", buf)
	}

	// read 3, left 0
	if n, err = rb.Read(buf); err != nil {
		t.Fatalf("read %d bytes error: %v", n, err)
	}
	if !bytes.Equal(buf, []byte("234")) {
		t.Fatalf("read error: %v", buf)
	}

	// read 3, should zero
	if n, err = rb.Read(buf); err != nil {
		t.Fatalf("read %d bytes error: %v", n, err)
	}
	if n != 0 {
		t.Fatalf("read error: %d", n)
	}
}

package bytes

import (
	"testing"
)

func TestSafeMakeByteSlice(t *testing.T) {
	s, err := safeMakeByteSlice(-1)

	if err != errOutOfMemory {
		t.Fatal("slice error not captured, slice capacity:", cap(s))
	}

	s, err = safeMakeByteSlice(1024)
	if err != nil {
		t.Fatal("slice error: ", err)
	}

	if cap(s) != 1024 {
		t.Fatal("slice capacity is expected 1024, but is:", cap(s))
	}
}

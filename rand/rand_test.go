package rand

import (
	"fmt"
	"testing"
)

func TestRandom32(t *testing.T) {
	base := 1000
	counts := make([]uint32, base)

	for i := 0; i < base*10; i++ {
		v, err := RandomUInt32(uint32(base))
		if err != nil {
			t.Fatal(err)
		}

		counts[v]++
	}

	for k, v := range counts {
		fmt.Printf("key %d shows %d times\n", k, v)
	}
}

func BenchmarkRandom32(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_, err := RandomUInt32(1024)
		if err != nil {
			fmt.Println(err)
		}
	}
}

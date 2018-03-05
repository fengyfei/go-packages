package runner

import (
	"testing"
)

func TestGoWithRecover(t *testing.T) {
	GoWithRecover(func() {
		panic("123")
	}, func(i interface{}) {
		if i != "123" {
			t.Error("go with custom recover handler get error")
		}
	})
}

package bytes

import "errors"

var (
	errOutOfMemory = errors.New("byte slice: not enough memory")
)

func safeMakeByteSlice(n int) (s []byte, err error) {
	defer func() {
		if recover() != nil {
			s = nil
			err = errOutOfMemory
		}
	}()

	s = make([]byte, n)
	return
}

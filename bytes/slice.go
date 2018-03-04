package bytes

import "errors"

var (
	errSliceTooLarge = errors.New("slice: too large size")
)

func safeMakeByteSlice(n int) (s []byte, err error) {
	defer func() {
		if recover() != nil {
			s = nil
			err = errSliceTooLarge
		}
	}()

	s = make([]byte, n)
	return
}

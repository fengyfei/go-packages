package rand

import (
	crand "crypto/rand"
	"fmt"
)

func securityRandom(size int) ([]byte, error) {
	if size < 0 {
		return nil, fmt.Errorf("secure rand: negative size %d", size)
	}

	buf := make([]byte, size)
	n, err := crand.Read(buf)

	if err != nil {
		return nil, err
	}

	if n != size {
		return nil, fmt.Errorf("secure rand: %d rand read but %d expected", n, size)
	}

	return buf, nil
}

// String generates a random string.
func String(size int) ([]byte, error) {
	return securityRandom(size)
}

// RandomUInt16 generates a random uint16 value in [0, max).
// If max is zero, returns the original random value.
func RandomUInt16(max uint16) (uint16, error) {
	buf, err := securityRandom(2)
	if err != nil {
		return 0, err
	}

	if max == 0 {
		return endian.Uint16(buf), nil
	}

	return endian.Uint16(buf) % max, nil
}

// RandomUInt32 generates a random uint32 value in [0, max).
// If max is zero, returns the original random value.
func RandomUInt32(max uint32) (uint32, error) {
	buf, err := securityRandom(4)
	if err != nil {
		return 0, err
	}

	if max == 0 {
		return endian.Uint32(buf), nil
	}

	return endian.Uint32(buf) % max, nil
}

// RandomUInt64 generates a random uint64 value in [0, max).
// If max is zero, returns the original random value.
func RandomUInt64(max uint64) (uint64, error) {
	buf, err := securityRandom(8)
	if err != nil {
		return 0, err
	}

	if max == 0 {
		return endian.Uint64(buf), nil
	}

	return endian.Uint64(buf) % max, nil
}

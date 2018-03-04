package rand

import (
	"encoding/binary"
	"unsafe"
)

var (
	endian binary.ByteOrder = binary.BigEndian
)

func init() {
	var i int32 = 0x01020304
	u := unsafe.Pointer(&i)
	pb := (*byte)(u)
	b := *pb

	if b == 0x04 {
		endian = binary.LittleEndian
	}

}

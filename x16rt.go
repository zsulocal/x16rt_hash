package x16rt_hash

// #cgo CFLAGS: -I.
// #cgo LDFLAGS: -L. -lx16rt_hash
// #include "x16rt.h"
import "C"

import (
	"unsafe"
)

func GetPowHash(hash []byte) []byte {
	powhash := make([]byte, 32)
	C.x16rt_hash((*C.char)(unsafe.Pointer(&hash[0])), (*C.char)(unsafe.Pointer(&powhash[0])))
	return powhash
}

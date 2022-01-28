package unsafe

import (
	"unsafe"
)

// StringHeader is the runtime representation of a string.
// It cannot be used safely or portably and its representation may
// change in a later release.
// Moreover, the Data field is not sufficient to guarantee the data
// it references will not be garbage collected, so programs must keep
// a separate, correctly typed pointer to the underlying data.
type StringHeader struct {
	Data uintptr
	Len  int
}

// SliceHeader is the runtime representation of a slice.
// It cannot be used safely or portably and its representation may
// change in a later release.
// Moreover, the Data field is not sufficient to guarantee the data
// it references will not be garbage collected, so programs must keep
// a separate, correctly typed pointer to the underlying data.
type SliceHeader struct {
	Data uintptr
	Len  int
	Cap  int
}

// BytesToString converts a byte slice to a string.
// It's fast, but not safe. Use it only if you know what you're doing.
// The byte slice passed to this function is not to be used after this call as it's unsafe; you have been warned.
func BytesToString(b []byte) string {
	bytesHeader := (*SliceHeader)(unsafe.Pointer(&b))
	strHeader := StringHeader{
		Data: bytesHeader.Data,
		Len:  bytesHeader.Len,
	}
	return *(*string)(unsafe.Pointer(&strHeader))
}

// StringToBytes converts a string to a byte slice.
// It's fast, but not safe. Use it only if you know what you're doing.
// The string passed to this functions is not to be used again after this call as it's unsafe; you have been warned.
func StringToBytes(s string) []byte {
	strHeader := (*StringHeader)(unsafe.Pointer(&s))
	bytesHeader := SliceHeader{
		Data: strHeader.Data,
		Len:  strHeader.Len,
		Cap:  strHeader.Len,
	}
	return *(*[]byte)(unsafe.Pointer(&bytesHeader))
}

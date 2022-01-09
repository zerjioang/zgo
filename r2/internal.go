package r2

import "unsafe"

//go:linkname typedmemmove reflect.typedmemmove
func typedmemmove(rtype unsafe.Pointer, dst, src unsafe.Pointer)
